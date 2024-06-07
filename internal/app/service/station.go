package service

import (
	"context"
	"geata/internal/app/handler"
	"geata/internal/app/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Station struct {
	ID             int64                                         `json:"id"`
	Name           string                                        `json:"name"`
	Host           string                                        `json:"host"`
	Port           int64                                         `json:"port"`
	ModelID        int64                                         `json:"modelId"`
	IsOnline       bool                                          `json:"isOnline"`
	LastOnlineTime string                                        `json:"lastOnlineTime"`
	Handlers       map[handler.HandlerType]handler.Handler       `json:"-"`
	Configs        map[handler.HandlerType]handler.HandlerConfig `json:"-"`
}

type StationData struct {
	ModelID int64
	Data    handler.Data
}

type IEC61750Config struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

type ModbusConfig struct {
	URL string `json:"url"`
}

type MQTTConfig struct {
	Broker   string `json:"broker"`
	ClientID string `json:"clientId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Topic    string `json:"topic"`
}

func (s *Station) Start(ctx context.Context, sd chan StationData) {
	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	handlerDataQueue := make(chan handler.Data)
	defer close(handlerDataQueue)
	for _, handler := range s.Handlers {
		go handler.Handle(newCtx, handlerDataQueue)
	}
	for {
		select {
		case <-ctx.Done():
			cancel()
			return
		case data := <-handlerDataQueue:
			sd <- StationData{
				ModelID: s.ModelID,
				Data:    data,
			}
		}
	}
}

// This function is used for initializing stations in whole system.
func StationInitFromDB(stationFromDB *model.Station) *Station {
	stationConfigs, err := model.GetStationConfigByStationID(Engine, stationFromDB.ID)
	if err != nil {
		return nil
	}
	configs := make(map[handler.HandlerType]handler.HandlerConfig)
	ic := handler.NewIEC61850HandlerConfig(stationConfigs.IEC61850Host, strconv.Itoa(int(stationConfigs.IEC61850Port)))
	configs[ic.Type()] = ic
	mc := handler.NewModbusHandlerConfig(stationConfigs.ModbusURL, stationFromDB.ModelID, Engine)
	configs[mc.Type()] = mc
	mqc := handler.NewMQTTHandlerConfig(stationConfigs.MQTTBroker, stationConfigs.MQTTClientID, stationConfigs.MQTTUsername, stationConfigs.MQTTPassword, stationConfigs.MQTTTopic)
	configs[mqc.Type()] = mqc

	handlers := make(map[handler.HandlerType]handler.Handler)
	for _, handlerType := range handler.HandlerTypes {
		config := configs[handlerType]
		handlers[handlerType] = config.NewHandler()
	}

	return &Station{
		ID:             stationFromDB.ID,
		Name:           stationFromDB.Name,
		Host:           stationFromDB.Host,
		Port:           stationFromDB.Port,
		ModelID:        stationFromDB.ModelID,
		IsOnline:       stationFromDB.IsOnline,
		LastOnlineTime: stationFromDB.LastOnlineTime.Format("2006-01-02 15:04:05"),
		Handlers:       handlers,
		Configs:        configs,
	}
}

// This function is used for converting station from DB to API response.
func StationFromDB(stationFromDB *model.Station) Station {
	return Station{
		ID:             stationFromDB.ID,
		Name:           stationFromDB.Name,
		Host:           stationFromDB.Host,
		Port:           stationFromDB.Port,
		ModelID:        stationFromDB.ModelID,
		IsOnline:       stationFromDB.IsOnline,
		LastOnlineTime: stationFromDB.LastOnlineTime.Format("2006-01-02 15:04:05"),
	}
}

// @Summary List all stations
// @Description List all stations
// @Produce json
// @Success 200 {object} []Station
// @Router /api/v1/station [get]
func ListStations(c *gin.Context) {
	stations, err := model.GetAllStations(Engine)
	resp := make([]Station, 0)
	for _, station := range stations {
		resp = append(resp, StationFromDB(station))
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get station by ID
// @Description Get station by ID
// @Produce json
// @Param station_id path int true "Station ID"
// @Success 200 {object} Station
// @Router /api/v1/station/{station_id} [get]
func GetStation(c *gin.Context) {
	ID := c.Param("station_id")
	statoinID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	station, err := model.GetStationByID(Engine, int64(statoinID))
	resp := StationFromDB(station)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Create station
// @Description Create station
// @Produce json
// @Param station body Station true "Station"
// @Success 200 {object} Station
// @Router /api/v1/station [post]
func CreateStation(c *gin.Context) {
	station := &Station{}
	err := c.BindJSON(station)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	stationInDB := &model.Station{
		Name:           station.Name,
		Host:           station.Host,
		Port:           station.Port,
		ModelID:        station.ModelID,
		IsOnline:       true,
		LastOnlineTime: time.Now(),
	}
	err = model.CreateStation(Engine, stationInDB)
	if err != nil {
		session.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	stationConfig := CreateStationConfigByStation(stationInDB)
	err = model.CreateStationConfig(Engine, stationConfig)
	if err != nil {
		session.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = session.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	station.ID = stationInDB.ID
	c.JSON(http.StatusOK, station)
}

// @Summary Update station by ID
// @Description Update station by ID
// @Produce json
// @Param station_id path int true "Station ID"
// @Param station body Station true "Station"
// @Success 200 {object} Station
// @Router /api/v1/station/{station_id} [put]
func UpdateStation(c *gin.Context) {
	station := &Station{}
	err := c.BindJSON(station)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = model.UpdateStation(Engine, &model.Station{
		ID:   station.ID,
		Name: station.Name,
		Host: station.Host,
		Port: station.Port,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, station)
}

// @Summary Delete station by ID
// @Description Delete station by ID
// @Produce json
// @Param station_id path int true "Station ID"
// @Success 200 {object} Station
// @Router /api/v1/station/{station_id} [delete]
func DeleteStation(c *gin.Context) {
	ID := c.Param("station_id")
	statoinID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	station, err := model.GetStationByID(Engine, int64(statoinID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sess := Engine.NewSession()
	defer sess.Close()
	err = sess.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = model.DeleteStationConfigByStationID(Engine, station.ID)
	if err != nil {
		sess.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = model.DeleteStation(Engine, station)
	if err != nil {
		sess.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = sess.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, station)
}

// @Summary Get Station IEC61850 Config by ID
// @Description Get Station IEC61850 Config by ID
// @Produce json
// @Param station_id path int true "Station ID"
// @Success 200 {object} IEC61850HandlerConfig
// @Router /api/v1/station/{station_id}/config/iec61850 [get]
func GetIEC61850ConfigForStation(c *gin.Context) {
	ID := c.Param("station_id")
	statoinID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stationConfig, err := model.GetStationConfigByStationID(Engine, int64(statoinID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := IEC61750Config{
		Host: stationConfig.IEC61850Host,
		Port: stationConfig.IEC61850Port,
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get Station Modbus Config by ID
// @Description Get Station Modbus Config by ID
// @Produce json
// @Param station_id path int true "Station ID"
// @Success 200 {object} ModbusHandlerConfig
// @Router /api/v1/station/{station_id}/config/modbus [get]
func GetModbusConfigForStation(c *gin.Context) {
	ID := c.Param("station_id")
	statoinID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stationConfig, err := model.GetStationConfigByStationID(Engine, int64(statoinID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := ModbusConfig{
		URL: stationConfig.ModbusURL,
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get Station MQTT Config by ID
// @Description Get Station MQTT Config by ID
// @Produce json
// @Param station_id path int true "Station ID"
// @Success 200 {object} MQTTHandlerConfig
// @Router /api/v1/station/{station_id}/config/mqtt [get]
func GetMqttConfigForStation(c *gin.Context) {
	ID := c.Param("station_id")
	statoinID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stationConfig, err := model.GetStationConfigByStationID(Engine, int64(statoinID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := MQTTConfig{
		Broker:   stationConfig.MQTTBroker,
		ClientID: stationConfig.MQTTClientID,
		Username: stationConfig.MQTTUsername,
		Password: stationConfig.MQTTPassword,
		Topic:    stationConfig.MQTTTopic,
	}
	c.JSON(http.StatusOK, resp)
}
