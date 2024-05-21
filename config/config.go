// File: `Config Initialization File` config/config.go
package PhoeniciaDigitalConfig

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type postgres struct {
	Postgres_user     string
	Postgres_password string
	Postgres_db       string
	Postgres_ssl      string
}

type _PhoeniciaDigitalConfig struct {
	Port     string
	Mongodb  string
	Postgres postgres
}

func loadConfig() (*_PhoeniciaDigitalConfig, error) {
	// Load environment variables from the .env file
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Create a new _BEUConfig struct and populate it with values from environment variables
	config := &_PhoeniciaDigitalConfig{
		Port:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Mongodb: os.Getenv("MONGODB"),
		Postgres: postgres{
			Postgres_user:     os.Getenv("POSTGRES_USER"),
			Postgres_password: os.Getenv("POSTGRES_PASSWORD"),
			Postgres_db:       os.Getenv("POSTGRES_DB"),
			Postgres_ssl:      os.Getenv("POSTGRES_SSL"),
		},
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
