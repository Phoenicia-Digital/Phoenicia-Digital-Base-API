package main

import (
	_ "Phoenicia-Digital-Base-API/base/database"
	PhoeniciaDigitalServer "Phoenicia-Digital-Base-API/base/server"
)

func main() {
	PhoeniciaDigitalServer.StartServer()
}
