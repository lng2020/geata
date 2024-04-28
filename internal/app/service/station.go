package service

import (
	"geata/internal/app/handler"
	"geata/internal/app/model"
	"net/http"

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
