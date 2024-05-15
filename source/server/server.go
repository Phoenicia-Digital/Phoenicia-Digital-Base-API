// File: `Server Implementation File` source/server/server.go
package PhoeniciaDigitalServer

import (
	PhoeniciaDigitalConfig "Phoenicia-Digital-Base-API/config"
	PhoeniciaDigitalUtils "Phoenicia-Digital-Base-API/source/utils"
	"fmt"
	"log"
	"net/http"
)

// Initialize Server Ecosystem Variables
var multiplexer *http.ServeMux = http.NewServeMux()

var PhoeniciaDigitalServer *http.Server = &http.Server{
	Addr:    PhoeniciaDigitalConfig.Config.Port,
	Handler: multiplexer,
}

func StartServer() {
	log.Printf("Server Running on http://localhost%s", PhoeniciaDigitalServer.Addr)
	PhoeniciaDigitalUtils.Log(fmt.Sprintf("Server Started on PORT --> %s", PhoeniciaDigitalServer.Addr))
	log.Fatal(PhoeniciaDigitalServer.ListenAndServe())
}

// Initialize Server Logic
// func init() {}
