package PhoeniciaDigitalServer

import (
	PhoeniciaDigitalConfig "Phoenicia-Digital-Base-API/config"
	PhoeniciaDigitalUtils "Phoenicia-Digital-Base-API/source/utils"
	"fmt"
	"log"
	"net/http"
)

// Initialize Server Ecosystem Variables
var _PhoeniciaDigitalMultiplexer *http.ServeMux = http.NewServeMux()

var PhoeniciaDigitalServer *http.Server = &http.Server{
	Addr:    PhoeniciaDigitalConfig.Config.Port,
	Handler: _PhoeniciaDigitalMultiplexer,
}

func StartServer() {
	log.Printf("Server Running on http://localhost%s", PhoeniciaDigitalServer.Addr)
	PhoeniciaDigitalUtils.Log(fmt.Sprintf("Server Started on PORT --> %s", PhoeniciaDigitalServer.Addr))
	log.Fatal(PhoeniciaDigitalServer.ListenAndServe())
}

// Initialize Server Logic
// func init() {}
