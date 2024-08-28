package PhoeniciaDigitalDatabase

import (
	PhoeniciaDigitalUtils "Phoenicia-Digital-Base-API/base/utils"
	PhoeniciaDigitalConfig "Phoenicia-Digital-Base-API/config"
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Phoenicia Digital's MongoDB Database client struct that will be used through our applciations
type mongodb struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// Implement a MongoDB Database Client to a global variable that will be used to manipulate and manage
// Our MongoDB Database Client
var Mongo *mongodb = implementMongoDB()

// Function Used to Implement a MongoDB Client
func implementMongoDB() *mongodb {
	// mongoDB variable that will returned to the Global MongoDB Database Client Variable
	var mongoDB *mongodb = &mongodb{}
	// Set a conStr Variable That will be used as the connection string to our MongoDB Database
	var conStr string

	// Check That DB field are not empty in .env file since logically it cant be Empty -
	// If implementing a MongoDB Database Connection
	// If fields are empty the API Will consider it intentional since the Specific API will not use
	// A MongoDB Database & Will return nil & Log The Warning that it will continue Running Withought
	// A MongoDB Database Connection
	if PhoeniciaDigitalConfig.Config.Mongo.Mongo_db != "" {
		// Start Implementing the Connection String For MongoDB if database field not empty in .env
		conStr = "mongodb://"
	} else {
		log.Printf("Continued with No MongoDB Database implementation! | In case expected a db connection MONGODB_DATABASE field REQUIRED ./config/.env\n")
		PhoeniciaDigitalUtils.Log("Continued with No MongoDB Database implementation! | In case expected a db connection MONGODB_DATABASE field REQUIRED ./config/.env\n")
		return nil
	}

	// Check If MongoDB user field is no empty and add the username:password@ field to the connection string
	// In case there is no username field filledout it will ignore this and continue to implement the
	// Connection String disregarding the user and password
	if PhoeniciaDigitalConfig.Config.Mongo.Mongo_user != "" {
		conStr += fmt.Sprintf("%s:%s@", PhoeniciaDigitalConfig.Config.Mongo.Mongo_user, PhoeniciaDigitalConfig.Config.Mongo.Mongo_password)
	}

	// Check if a spcific MongoDB host is provided and append that to the Connection String
	if PhoeniciaDigitalConfig.Config.Mongo.Mongo_host != "" {

		conStr += PhoeniciaDigitalConfig.Config.Mongo.Mongo_host
	} else if PhoeniciaDigitalConfig.Config.Project_Name != "" {
		// If a Mongodb Host Is NOT Provided revert back to the Project Name
		// Due to how our Dockerfile & docker-compose is set up
		// the Mongodb Container Service is called {Project Name}-Mongodb
		// then Get the Project Name and Append if to conStr wit the proper format

		conStr += fmt.Sprintf("%s-Mongodb", PhoeniciaDigitalConfig.Config.Project_Name)
	} else {
		// If a host is not provided for For the MongoDB Connection it will result in appending
		// Phoenicia-Digital-Mongodb <MONGODB DEFAULT HOST> according to
		// Dockerfile & docker-compose Setup to the Connection String

		conStr += "Phoenicia-Digital-Mongodb"
	}

	// Check if a specific Port is provided for MongoDB coonection
	if PhoeniciaDigitalConfig.Config.Mongo.Mongo_port != "" {
		// Make sure the provided Port is an actual Port <Intiger in range 0 <-> 65535>
		if portNum, err := strconv.Atoi(PhoeniciaDigitalConfig.Config.Mongo.Mongo_port); err != nil {
			// In case the provided port is not an integer the API will consider it an error
			// And will Exist the process logging the error & issue
			PhoeniciaDigitalUtils.Log(fmt.Sprintf("Failed to implement MongoDB client | Invalid Port Number: %s", PhoeniciaDigitalConfig.Config.Mongo.Mongo_port))
			log.Fatalf("Failed to implement MongoDB client | Invalid Port Number: %s", PhoeniciaDigitalConfig.Config.Mongo.Mongo_port)
			return nil
		} else {
			if portNum < 0 || portNum > 65535 {
				// In case the provided port is not in range 0 <-> 65535 the API will consider it an error
				// And will Exist the process logging the error & issue
				PhoeniciaDigitalUtils.Log(fmt.Sprintf("MONGODB PORT: %s is OUT OF RANGE 0 --> 65535 | Change in ./config/.env", PhoeniciaDigitalConfig.Config.Mongo.Mongo_port))
				log.Fatalf("MONGODB PORT: %s is OUT OF RANGE 0 --> 65535 | Change in ./config/.env", PhoeniciaDigitalConfig.Config.Mongo.Mongo_port)
				return nil
			} else {
				// If all is good and checked append the specific Port to the MongoDB Connection string
				conStr += fmt.Sprintf(":%s", PhoeniciaDigitalConfig.Config.Mongo.Mongo_port)
			}
		}
	} else {
		// If no specific Port is provided it will append the port 27017 to MongoDB Connection String
		// Which is <MONGODB DEFAULT PORT>
		conStr += ":27017"

	}

	// Check the ssl type for The MongoDB Client Connection
	if PhoeniciaDigitalConfig.Config.Mongo.Mongo_ssl == "true" {
		// if the sll is set to true
		conStr += "/?ssl=true"
	}

	// Try and connect to the MongoDB Client with the generated Connection String With all fields specified
	if clientConnection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conStr)); err != nil {
		// If there was an error connecting Exist the process Logging the Error
		PhoeniciaDigitalUtils.Log(fmt.Sprintf("Failed to create MongoDB client | Verify MONGODB_HOST: %s | Verify MONGODB_PORT: %s", PhoeniciaDigitalConfig.Config.Mongo.Mongo_host, PhoeniciaDigitalConfig.Config.Mongo.Mongo_port))
		log.Fatalf("Failed to create MongoDB client | Verify MONGODB_HOST: %s | Verify MONGODB_PORT: %s", PhoeniciaDigitalConfig.Config.Mongo.Mongo_host, PhoeniciaDigitalConfig.Config.Mongo.Mongo_port)
		return nil
	} else {
		// In case Connection returned a *mongo.Client & No Errors occured set mongoDB.Client to struct mongodb
		// This Struct is specific to Phoenicia Digital Check Line 15 To Line 17
		mongoDB.Client = clientConnection
	}

	// Try and Ping the MongoDB Client Making Sure that a positive connection has been established
	if err := mongoDB.Client.Ping(context.Background(), nil); err != nil {
		// In case there was an error returned EXIST the process Logging the Error
		PhoeniciaDigitalUtils.Log(fmt.Sprintf("Failed to connect MongoDB client: %s | Service might be down or WRONG PORT | ERROR: %s", conStr, err.Error()))
		log.Fatalf("Failed to connect MongoDB client: %s | Service might be down or WRONG PORT | ERROR: %s", conStr, err.Error())
		return nil
	} else {
		// If the Ping was successful Return a database handle and set it to mongoDB.DB
		mongoDB.DB = mongoDB.Client.Database(PhoeniciaDigitalConfig.Config.Mongo.Mongo_db)
	}

	log.Printf("Implemented Mongodb Database connection | settings: %s\n", conStr)
	// In case all fields are populated as needed & Pings were successful return the Client Struct
	return mongoDB
}
