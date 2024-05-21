package main

import (
	_ "Phoenicia-Digital-Base-API/source/database"
	PhoeniciaDigitalServer "Phoenicia-Digital-Base-API/source/server"
)

func main() {
	PhoeniciaDigitalServer.StartServer()
}
