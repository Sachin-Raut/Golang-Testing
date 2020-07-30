package controllers

import (
	"net/http"

	"github.com/Sachin-Raut/Golang-Testing/src/api/services"
	"github.com/gin-gonic/gin"
)

//GetCountry is
func GetCountry(c *gin.Context) {
	country, err := services.LocationService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
