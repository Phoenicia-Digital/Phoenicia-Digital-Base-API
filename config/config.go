// File: `Config Initialization File` config/config.go
package PhoeniciaDigitalConfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type _PhoeniciaDigitalConfig struct {
	Port          string
	MongodbServer string
}

func loadConfig() (*_PhoeniciaDigitalConfig, error) {
	// Load environment variables from the .env file
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Create a new _BEUConfig struct and populate it with values from environment variables
	config := &_PhoeniciaDigitalConfig{
		Port:          os.Getenv("PORT"),
		MongodbServer: os.Getenv("MONGODB"),
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
