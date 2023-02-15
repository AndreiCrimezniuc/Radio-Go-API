package router_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	stationRepository "nokogiriwatir/radio-main/pkg/station_repository"
)

func (env *DIEnv) HandleStation(c *gin.Context) {
	station, er := stationRepository.GetStation(c.Param("slug"), env.Db)

	if er != nil {
		env.Logger.Error("Error in getStation" + er.Error())
		c.String(http.StatusNotFound, "there is no station or "+er.Error())
	} else {
		c.JSON(http.StatusOK, &station)
	}
}

func (env *DIEnv) HandleStations(c *gin.Context) {
	stations, er := stationRepository.GetStations(env.Db)

	if er != nil {
		env.Logger.Error("Error in getStations" + er.Error())
		c.String(http.StatusNotFound, "there is no stations or "+er.Error())
	} else {
		c.JSON(http.StatusOK, &stations)
	}
}
