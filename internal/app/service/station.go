package service

import (
	"geata/internal/app/handler"
	"geata/internal/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Station struct {
	ID       int64
	Name     string
	Host     string
	Port     int64
	Handlers map[handler.HandlerType]handler.Handler
	Configs  map[handler.HandlerType]handler.HandlerConfig
}

func (s *Station) InitFromDB(stationFromDB *model.Station) error {
	s.ID = stationFromDB.ID
	s.Name = stationFromDB.Name
	s.Host = stationFromDB.Host
	s.Port = stationFromDB.Port

	return nil
}

// @Summary List all stations
// @Description List all stations
// @Produce json
// @Success 200 {object} []Station
// @Router /api/v1/stations [get]
func ListStations(c *gin.Context) {
	stations, err := model.GetAllStations(Engine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stations)
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, station)
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
	err = model.CreateStation(Engine, &model.Station{
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
