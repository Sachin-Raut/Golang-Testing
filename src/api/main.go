package main

import (
	"fmt"

	"github.com/Sachin-Raut/Golang-Testing/src/api/domain/locations/providerlocation"
)

func main() {
	country, err := providerlocation.GetCountry("AR")
	fmt.Println(err)
	fmt.Println(country)
}
