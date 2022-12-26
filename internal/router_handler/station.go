package router_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nokogiriwatir/radio-main/internal/station_repository"
)

func (env *DIEnv) HandleStation(c *gin.Context) {
	station, er := station_repository.GetStation(c.Param("slug"), env.Db)

	if er != nil {
		env.Logger.Error("Error in getStations" + er.Error())
		c.String(http.StatusNotFound, "there is no station or error in code")
	} else {
		c.JSON(http.StatusOK, &station)
	}
}
