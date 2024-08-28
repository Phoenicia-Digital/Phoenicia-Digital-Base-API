package PhoeniciaDigitalDatabase

import (
	PhoeniciaDigitalUtils "Phoenicia-Digital-Base-API/base/utils"
	PhoeniciaDigitalConfig "Phoenicia-Digital-Base-API/config"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type postgres struct {
	DB *sql.DB
}

var Postgres *postgres = &postgres{
	DB: implementPostgres(),
}

// This Function Reads .sql Files With their queries or sql commands
// from the sql folder returning the query string inside the .sql file

func (p postgres) ReadSQL(fileName string) (string, error) {

	// Read the File Provided <Sould Only Be the Filename withought .sql || /sql/file.sql
	// as in myQuery NOT myQuery.sql || /sql/myQuery

	// Structures the ReadFile Path Automatically returns a log error message in case failed to read
	// the .sql file returning an empty string as a query and an error
	if query, err := os.ReadFile(fmt.Sprintf("./sql/%s.sql", fileName)); err != nil {
		log.Printf("Error reading query file: %s | Error: %s", query, err.Error())
		PhoeniciaDigitalUtils.Log(fmt.Sprintf("Error reading query file: %s | Error: %s", query, err.Error()))
		return "", err
	} else {
		// In case succesful returns a string query and nil for error
		return string(query), nil
	}
}

// This Function Prepares && Returns a *sql.Stmt that can be used for queries
// Works hand in hand with the ReadSQL function for postgres struct

func (p postgres) PrepareSQL(fileName string) (*sql.Stmt, error) {

	// Works On the Reading <For More info check the function above 'ReadSQL'>
	// Returns a nil *sql.stmt and an error if failed to read query
	if query, err := p.ReadSQL(fileName); err != nil {
		return nil, err
	} else {
		// Tries to prepare the query to be executed
		// PREPARING QUERIES IS THE SAFEST METHOD TO USE QUERIES SINCE THEY PREVENT SQL INJECTIONS
		// If the preparation failed returns a nil *sql.stmt and an error
		if stmt, err := p.DB.Prepare(query); err != nil {
			log.Printf("Error preparing query: %s | Error: %s", query, err.Error())
			PhoeniciaDigitalUtils.Log(fmt.Sprintf("Error preparing query: %s | Error: %s", query, err.Error()))
			return nil, err
		} else {
			// If all went well returns a *sql.stmt that was prepared and nil for error
			// DONT FORGET TO DEFER THE STMT ONCE DONE defer stmt.Close()
			return stmt, nil
		}
	}
}

// This Function Queries a SQL Row Returning a *sql.Row & an error
// Works hand in hand with the ReadSQL & PrepareSQL function for postgres struct
// THE BENEFIT OF THIS FUNCTION IS THAT IT RETURNS A *sql.Row that does not need to be defer Closed
// And the *sql.Stmt returned by PrepareSQL will be closed automatically making this more secure

func (p postgres) SecureQuerySQLRow(fileName string, args ...any) (*sql.Row, error) {
	// Uses the PrepareSQL method of postgres struct to return a stmt in case an error occured it will return
	// a nil stmt with the error
	if stmt, err := p.PrepareSQL(fileName); err != nil {
		return nil, err
	} else {
		// In case successful we will query stmt with the args provided and make sure that we defer stmt.Close()
		// for more secure queries
		defer stmt.Close()
		return stmt.QueryRow(args...), nil
	}
}

// This Function Executes a SQL Stmt Returning a *sql.Result & an error
// Works hand in hand with the ReadSQL & PrepareSQL function for postgres struct
// THE BENEFIT OF THIS FUNCTION IS THAT IT RETURNS A *sql.Result that does not need to be defer Closed
// And the *sql.Stmt returned by PrepareSQL will be closed automatically making this more secure

func (p postgres) SecureExecSQL(fileName string, args ...any) (*sql.Result, error) {
	// Uses the PrepareSQL method of postgres struct to return a stmt in case an error occured it will return
	// a nil stmt with the error
	if stmt, err := p.PrepareSQL(fileName); err != nil {
		return nil, err
	} else {
		// In case successful we will Execute stmt with the args provided and make sure that we defer stmt.Close()
		// for more secure queries
		defer stmt.Close()
		if res, err := stmt.Exec(args...); err != nil {
			log.Printf("Error Executing Query | Error: %s", err.Error())
			PhoeniciaDigitalUtils.Log(fmt.Sprintf("Error Executing Query | Error: %s", err.Error()))
			return nil, err
		} else {
			return &res, nil
		}
	}
}

// This Function Implements The Postgresql Database Connection Returning a *sql.DB
// to var Postgres.DB Which Can be used Globally In the project

func implementPostgres() *sql.DB {
	// Set a conStr Variable That will be used as the connection string to our Postgres Database
	var conStr string

	// Check That User & DB fields are not empty in .env file since logically they cant be Empty
	// If implementing a Postgres Database Connection
	// If fields are empty the API Will consider it intentional since the Specific API will not use
	// Postgres Database & Will return nil & Log The Warning that it will continue Running Withought
	// A Postgres Database Connection
	if PhoeniciaDigitalConfig.Config.Postgres.Postgres_user == "" || PhoeniciaDigitalConfig.Config.Postgres.Postgres_db == "" {
		log.Printf("Continued with No Postgres Database implementation! | In case expected a db connection POSTGRES_USER & POSTGRES_DB fields REQUIRED ./config/.env\n")
		PhoeniciaDigitalUtils.Log("Continued with No Postgres Database implementation! | In case expected a db connection POSTGRES_USER & POSTGRES_DB fields REQUIRED ./config/.env")
		return nil
	} else {
		// Other Wise Append to conStr the user & db properties by default
		conStr += fmt.Sprintf("user=%s dbname=%s", PhoeniciaDigitalConfig.Config.Postgres.Postgres_user, PhoeniciaDigitalConfig.Config.Postgres.Postgres_db)
	}

	// Check if a specific host is provided localhost || IP-Address to the Postgres Database || Service Host
	// Will be appended to conStr
	if PhoeniciaDigitalConfig.Config.Postgres.Postgres_host != "" {
		conStr += fmt.Sprintf(" host=%s", PhoeniciaDigitalConfig.Config.Postgres.Postgres_host)
	} else if PhoeniciaDigitalConfig.Config.Project_Name != "" {
		// By default if field is commented or not set We revert Back to Project Name
		// Due to how our backend Containers are set up if a project name exists we connect
		// via {Project Name}-Postgres then
		// this {Project Name}-Postgres be appended to conStr
		conStr += fmt.Sprintf(" host=%s-Postgres", PhoeniciaDigitalConfig.Config.Project_Name)
	} else {
		// 	By default if field is commented or not set in .env it will be set as Phoenicia-Digital-Postgres
		//  <POSTGRESQL DEFAULT> Will be appended to conStr considering this project will be dockerised and make use
		//	of docker-compose & Dockerfile
		conStr += " host=Phoenicia-Digital-Postgres"
	}

	// Check If Port field is filled out & Make sure it an intiger from Range 0 to 65535
	// In cases where port is not an int or out of range The API will consider it unintentional
	// Due to a typo and exist the process logging the issue
	if PhoeniciaDigitalConfig.Config.Postgres.Postgres_port != "" {
		if portNumber, err := strconv.Atoi(PhoeniciaDigitalConfig.Config.Postgres.Postgres_port); err != nil {
			PhoeniciaDigitalUtils.Log(fmt.Sprintf("POSTGRES PORT is Invalid: %s != int | Change in ./config/.env", PhoeniciaDigitalConfig.Config.Postgres.Postgres_port))
			log.Fatalf("POSTGRES PORT is Invalid: %s != int | Change in ./config/.env", PhoeniciaDigitalConfig.Config.Postgres.Postgres_port)
			return nil
		} else {
			if portNumber < 0 || portNumber > 65535 {
				PhoeniciaDigitalUtils.Log(fmt.Sprintf("POSTGRES PORT: %s is OUT OF RANGE 0 --> 65535 | Change in ./config/.env", PhoeniciaDigitalConfig.Config.Postgres.Postgres_port))
				log.Fatalf("POSTGRES PORT: %s is OUT OF RANGE 0 --> 65535 | Change in ./config/.env", PhoeniciaDigitalConfig.Config.Postgres.Postgres_port)
				return nil
			} else {
				// If all is good the port will become set to the specified PORT in .env
				// Will be appended to conStr
				conStr += fmt.Sprintf(" port=%s", PhoeniciaDigitalConfig.Config.Postgres.Postgres_port)
			}
		}
	} else {
		// By default if Port field is not filled out it will be set to 5432 <POSTGRESQL DEFAULT>
		// Will be appended to conStr
		conStr += " port=5432"
	}

	// Check if a password is given & append the password Field to the conStr else ignore this step
	// So NO PASSWORD WILL BE USED
	if PhoeniciaDigitalConfig.Config.Postgres.Postgres_password != "" {
		conStr += fmt.Sprintf(" password=%s", PhoeniciaDigitalConfig.Config.Postgres.Postgres_password)
	}

	// Check the Postgres SSL Mode verifying it a valid mode & Set it to that specific mode
	// Will be appended to conStr
	switch PhoeniciaDigitalConfig.Config.Postgres.Postgres_ssl {
	case "disable", "require", "verify-ca", "verify-full":
		conStr += fmt.Sprintf(" sslmode=%s", PhoeniciaDigitalConfig.Config.Postgres.Postgres_ssl)
	default:
		// In case SSL Mode Field is empty then the sslmode will be set to disable <POSTGRESQL DEFAULT>
		// Will be appended to conStr
		conStr += " sslmode=disable"
	}

	// Try and implement a Postgresql Database Connection with the provided conStr
	// If error is encountered Exit out of the process loging the issue & Error
	if db, err := sql.Open("postgres", conStr); err != nil {
		log.Fatalf("Failed to implement Postgres Database | Error: %s", err.Error())
		PhoeniciaDigitalUtils.Log(fmt.Sprintf("Failed to implement Postgres Database | Error: %s", err.Error()))
		return nil
	} else {
		// If the Connection was established Ping the Database to check if all is good
		// Otherwise Exit out of the process logging the issue & Error
		if err := db.Ping(); err != nil {
			log.Fatalf("Failed to connect to Postgres Database | Verify Postgres Database config values ./config/.env | Error: %s", err.Error())
			return nil
		} else {
			// Make sure the database name provided is correct by querying something <RETRIEVING SOME ROW>
			// if an error occured most likely due to typo in database name or non existance of the database
			// Therefore Exist the process logging the issue & Error
			if rows, err := db.Query("SELECT 1"); err != nil {
				log.Fatalf("Database Name: %s Does NOT EXIST | Change at ./config/.env | Error: %s", PhoeniciaDigitalConfig.Config.Postgres.Postgres_db, err.Error())
				return nil
			} else {
				// Once all is good close the queried row to check database existance and log that a
				// Postgresql Database has been implemented with the provided properties
				// returning the db which a *sql.DB
				rows.Close()
				log.Printf("Implemented Postgres Database connection | settings: %s\n", conStr)
				return db
			}
		}
	}
}
