package service

import (
	"geata/internal/app/handler"
	"geata/internal/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Station struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Host           string `json:"host"`
	Port           int64  `json:"port"`
	ModelHash      string `json:"modelHash"`
	IsOnline       bool   `json:"isOnline"`
	LastOnlineTime string `json:"lastOnlineTime"`
	Handlers       map[handler.HandlerType]handler.Handler
	Configs        map[handler.HandlerType]handler.HandlerConfig
}

func StationFromDB(stationFromDB *model.Station) *Station {
	stationConfigs, err := model.GetStationConfigByStationID(Engine, stationFromDB.ID)
	if err != nil {
		return nil
	}
	ic := handler.NewIEC61850HandlerConfig(stationConfigs.IEC61850Host, strconv.Itoa(int(stationConfigs.IEC61850Port)))
	mc := handler.NewModbusHandlerConfig(stationConfigs.ModbusURL)
	mqc := handler.NewMQTTHandlerConfig(stationConfigs.MQTTBroker, stationConfigs.MQTTClientID, stationConfigs.MQTTUsername, stationConfigs.MQTTPassword, stationConfigs.MQTTTopic)
	ih := handler.NewIEC61850Handler(ic)
	mh := handler.NewModbusHandler(mc)
	mqh := handler.NewMQTTHandler(mqc)

	configs := make(map[handler.HandlerType]handler.HandlerConfig)
	handlers := make(map[handler.HandlerType]handler.Handler)
	for _, h := range []handler.Handler{ih, mh, mqh} {
		handlers[h.Type()] = h
	}
	for _, hc := range []handler.HandlerConfig{ic, mc, mqc} {
		configs[hc.Type()] = hc
	}

	return &Station{
		ID:             stationFromDB.ID,
		Name:           stationFromDB.Name,
		Host:           stationFromDB.Host,
		Port:           stationFromDB.Port,
		ModelHash:      stationFromDB.ModelHash,
		IsOnline:       stationFromDB.IsOnline,
		LastOnlineTime: stationFromDB.LastOnlineTime.Format("2006-01-02 15:04:05"),
		Handlers:       make(map[handler.HandlerType]handler.Handler),
		Configs:        make(map[handler.HandlerType]handler.HandlerConfig),
	}
}

// @Summary List all stations
// @Description List all stations
// @Produce json
// @Success 200 {object} []Station
// @Router /api/v1/stations [get]
func ListStations(c *gin.Context) {
	stations, err := model.GetAllStations(Engine)
	resp := make([]Station, 0)
	for _, station := range stations {
		resp = append(resp, *StationFromDB(station))
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
// @Router /api/v1/stations/{station_id} [get]
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
	c.JSON(http.StatusOK, *resp)
}

// @Summary Create station
// @Description Create station
// @Produce json
// @Param station body Station true "Station"
// @Success 200 {object} Station
// @Router /api/v1/stations [post]
func CreateStation(c *gin.Context) {
	station := &Station{}
	err := c.BindJSON(station)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create station and station config
	session := Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = model.CreateStation(Engine, &model.Station{
		Name:      station.Name,
		Host:      station.Host,
		Port:      station.Port,
		ModelHash: station.ModelHash,
		IsOnline:  false,
	})
	if err != nil {
		session.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	stationId := model.GetStationIDByModelHash(Engine, station.ModelHash)
	stationConfig := CreateDefaultStationConfig(stationId)
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
	c.JSON(http.StatusOK, station)
}

// @Summary Update station by ID
// @Description Update station by ID
// @Produce json
// @Param station_id path int true "Station ID"
// @Param station body Station true "Station"
// @Success 200 {object} Station
// @Router /api/v1/stations/{station_id} [put]
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
// @Router /api/v1/stations/{station_id} [delete]
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
	err = model.DeleteStation(Engine, station)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, station)
}
