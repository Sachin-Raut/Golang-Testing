package app

import "github.com/Sachin-Raut/Golang-Testing/src/api/controllers"

func mapUrls() {
	router.GET("/countries/:country_id", controllers.GetCountry)
}
