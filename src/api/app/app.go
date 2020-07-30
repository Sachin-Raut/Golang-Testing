package app

import (
	"github.com/Sachin-Raut/Golang-Testing/src/api/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApp is
func StartApp() {
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func mapUrls() {
	router.GET("/countries/:country_id", controllers.GetCountry)
}
