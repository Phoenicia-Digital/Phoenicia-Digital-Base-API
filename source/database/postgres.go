package PhoeniciaDigitalDatabase

import (
	PhoeniciaDigitalConfig "Phoenicia-Digital-Base-API/config"
	PhoeniciaDigitalUtils "Phoenicia-Digital-Base-API/source/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type postgres struct {
	DB *sql.DB
}

var Postgres *postgres = &postgres{
	DB: implementPostgres(),
}

func implementPostgres() *sql.DB {
	if PhoeniciaDigitalConfig.Config.Postgres.Postgres_user != "" && PhoeniciaDigitalConfig.Config.Postgres.Postgres_db != "" {
		if PhoeniciaDigitalConfig.Config.Postgres.Postgres_password != "" {
			if PhoeniciaDigitalConfig.Config.Postgres.Postgres_ssl != "" {
				var conStr string = fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", PhoeniciaDigitalConfig.Config.Postgres.Postgres_user, PhoeniciaDigitalConfig.Config.Postgres.Postgres_db, PhoeniciaDigitalConfig.Config.Postgres.Postgres_password, PhoeniciaDigitalConfig.Config.Postgres.Postgres_ssl)
				if db, err := sql.Open("postgres", conStr); err != nil && db != nil {
					log.Fatalf("Failed to implement Postgres Database | Error: %s", err.Error())
					PhoeniciaDigitalUtils.Log(fmt.Sprintf("Failed to implement Postgres Database | Error: %s", err.Error()))
					return nil
				} else {
					if err := db.Ping(); err != nil {
						log.Fatalf("Failed to connect to Postgres Database | Verify Postgres Database config values ./config/.env | Error: %s", err.Error())
						return nil
					} else {
						log.Printf("Implemented Postgres Database connection settings: user=%s dbname=%s password=*** sslmode=%s\n", PhoeniciaDigitalConfig.Config.Postgres.Postgres_user, PhoeniciaDigitalConfig.Config.Postgres.Postgres_db, PhoeniciaDigitalConfig.Config.Postgres.Postgres_ssl)
						return db
					}
				}
			} else {
				var conStr string = fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", PhoeniciaDigitalConfig.Config.Postgres.Postgres_user, PhoeniciaDigitalConfig.Config.Postgres.Postgres_db, PhoeniciaDigitalConfig.Config.Postgres.Postgres_password)
				if db, err := sql.Open("postgres", conStr); err != nil && db != nil {
					log.Fatalf("Failed to implement Postgres Database | Error: %s", err.Error())
					PhoeniciaDigitalUtils.Log(fmt.Sprintf("Failed to implement Postgres Database | Error: %s", err.Error()))
					return nil
				} else {
					if err := db.Ping(); err != nil {
						log.Fatalf("Failed to connect to Postgres Database | Verify Postgres Database config values ./config/.env | Error: %s", err.Error())
						return nil
					} else {
						log.Printf("Implemented Postgres Database connection settings: user=%s dbname=%s password=*** sslmode=disable\n", PhoeniciaDigitalConfig.Config.Postgres.Postgres_user, PhoeniciaDigitalConfig.Config.Postgres.Postgres_db)
						return db
					}
				}
			}
		} else {
			if PhoeniciaDigitalConfig.Config.Postgres.Postgres_ssl != "" {
				var conStr string = fmt.Sprintf("user=%s dbname=%s sslmode=%s", PhoeniciaDigitalConfig.Config.Postgres.Postgres_user, PhoeniciaDigitalConfig.Config.Postgres.Postgres_db, PhoeniciaDigitalConfig.Config.Postgres.Postgres_ssl)
				if db, err := sql.Open("postgres", conStr); err != nil && db != nil {
					log.Fatalf("Failed to implement Postgres Database | Error: %s", err.Error())
					PhoeniciaDigitalUtils.Log(fmt.Sprintf("Failed to implement Postgres Database | Error: %s", err.Error()))
					return nil
				} else {
					if err := db.Ping(); err != nil {
						log.Fatalf("Failed to connect to Postgres Database | Verify Postgres Database config values ./config/.env | Error: %s", err.Error())
						return nil
					} else {
						log.Printf("Implemented Postgres Database connection settings: user=%s dbname=%s sslmode=%s\n", PhoeniciaDigitalConfig.Config.Postgres.Postgres_user, PhoeniciaDigitalConfig.Config.Postgres.Postgres_db, PhoeniciaDigitalConfig.Config.Postgres.Postgres_ssl)
						return db
					}
				}
			} else {
				var conStr string = fmt.Sprintf("user=%s dbname=%s sslmode=disable", PhoeniciaDigitalConfig.Config.Postgres.Postgres_user, PhoeniciaDigitalConfig.Config.Postgres.Postgres_db)
				if db, err := sql.Open("postgres", conStr); err != nil && db != nil {
					log.Fatalf("Failed to implement Postgres Database | Error: %s", err.Error())
					PhoeniciaDigitalUtils.Log(fmt.Sprintf("Failed to implement Postgres Database | Error: %s", err.Error()))
					return nil
				} else {
					if err := db.Ping(); err != nil {
						log.Fatalf("Failed to connect to Postgres Database | Verify Postgres Database config values ./config/.env | Error: %s", err.Error())
						return nil
					} else {
						log.Printf("Implemented Postgres Database connection settings: user=%s dbname=%s sslmode=disable\n", PhoeniciaDigitalConfig.Config.Postgres.Postgres_user, PhoeniciaDigitalConfig.Config.Postgres.Postgres_db)
						return db
					}
				}
			}
		}
	} else {
		log.Printf("Continued with No Postgres Database implementation! | In case expected a db connection POSTGRES_USER & POSTGRES_DB fields REQUIRED ./config/.env\n")
		PhoeniciaDigitalUtils.Log("Continued with No Postgres Database implementation! | In case expected a db connection POSTGRES_USER & POSTGRES_DB fields REQUIRED ./config/.env")
		return nil
	}
}

func init() {
	if Postgres.DB != nil {
		defer Postgres.DB.Close()
	}
}
