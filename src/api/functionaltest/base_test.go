package functionaltest

import (
	"fmt"
	"os"
	"testing"

	"github.com/Sachin-Raut/Golang-Testing/src/api/app"
	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("about to start app")
	go app.StartApp()
	fmt.Println("about to start test cases")
	os.Exit(m.Run())
}
