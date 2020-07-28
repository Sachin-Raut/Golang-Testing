package main

import (
	"fmt"

	"github.com/Sachin-Raut/Golang-Testing/src/api/domain/locations/providerlocation"
)

func main() {
	country, err := providerlocation.GetCountry("ARs")
	fmt.Println(err)
	fmt.Println(country)
}
