// File: `Config Initialization File` config/config.go
package PhoeniciaDigitalConfig

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type _PhoeniciaDigitalConfig struct {
	Project_Name string
	Port         string
	CORS         string
	Postgres     postgres
	Mongo        mongo
	Redis        redis
}

type postgres struct {
	Postgres_host     string
	Postgres_port     string
	Postgres_user     string
	Postgres_password string
	Postgres_db       string
	Postgres_ssl      string
}

type mongo struct {
	Mongo_host     string
	Mongo_port     string
	Mongo_db       string
	Mongo_user     string
	Mongo_password string
	Mongo_ssl      string
}

type redis struct {
	Redis_host     string
	Redis_port     string
	Redis_password string
}

func loadConfig() (*_PhoeniciaDigitalConfig, error) {
	// Load environment variables from the .env file
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Create a new _BEUConfig struct and populate it with values from environment variables
	config := &_PhoeniciaDigitalConfig{
		Project_Name: os.Getenv("PROJECT_NAME"),
		Port:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		CORS:         os.Getenv("CORS_ORIGIN"),
		Postgres: postgres{
			Postgres_host:     os.Getenv("POSTGRES_HOST"),
			Postgres_port:     os.Getenv("POSTGRES_PORT"),
			Postgres_user:     os.Getenv("POSTGRES_USER"),
			Postgres_password: os.Getenv("POSTGRES_PASSWORD"),
			Postgres_db:       os.Getenv("POSTGRES_DB"),
			Postgres_ssl:      os.Getenv("POSTGRES_SSL"),
		},
		Mongo: mongo{
			Mongo_host:     os.Getenv("MONGODB_HOST"),
			Mongo_port:     os.Getenv("MONGODB_PORT"),
			Mongo_db:       os.Getenv("MONGODB_DATABASE"),
			Mongo_user:     os.Getenv("MONGODB_USER"),
			Mongo_password: os.Getenv("MONGODB_PASSWORD"),
			Mongo_ssl:      os.Getenv("MONGODB_SSL"),
		},
		Redis: redis{
			Redis_host:     os.Getenv("REDIS_HOST"),
			Redis_port:     os.Getenv("REDIS_PORT"),
			Redis_password: os.Getenv("Redis_PASSWORD"),
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
