// File: `Config Initialization File` config/config.go
package PhoeniciaDigitalConfig

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type _PhoeniciaDigitalConfig struct {
	Port     string
	Mongodb  string
	Postgres string
}

func loadConfig() (*_PhoeniciaDigitalConfig, error) {
	// Load environment variables from the .env file
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Create a new _BEUConfig struct and populate it with values from environment variables
	config := &_PhoeniciaDigitalConfig{
		Port:     fmt.Sprintf(":%s", os.Getenv("PORT")),
		Mongodb:  os.Getenv("MONGODB"),
		Postgres: os.Getenv("POSTGRES"),
	}

	return config, nil
}

var Config *_PhoeniciaDigitalConfig

func init() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	Config = config
}
