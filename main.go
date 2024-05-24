package main

import (
	DB "Phoenicia-Digital-Base-API/base/database"
	PhoeniciaDigitalServer "Phoenicia-Digital-Base-API/base/server"
	"context"
)

func main() {
	if DB.Postgres.DB != nil {
		defer DB.Postgres.DB.Close()
	}

	if DB.Mongo != nil && DB.Mongo.Client != nil && DB.Mongo.DB != nil && DB.Mongo.Collection != nil {
		defer DB.Mongo.Client.Disconnect(context.Background())
	}

	PhoeniciaDigitalServer.StartServer()
}
