package PhoeniciaDigitalDatabase

import (
	PhoeniciaDigitalUtils "Phoenicia-Digital-Base-API/base/utils"
	PhoeniciaDigitalConfig "Phoenicia-Digital-Base-API/config"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client = implementRedisDB()

func implementRedisDB() *redis.Client {

	// Set a conStr Variable That will be used as the connection string to our Redis Database
	var conStr string

	if PhoeniciaDigitalConfig.Config.Redis.Redis_host != "" {
		conStr += PhoeniciaDigitalConfig.Config.Redis.Redis_host
	} else {
		conStr += "localhost"
	}

	if PhoeniciaDigitalConfig.Config.Redis.Redis_port != "" {
		if portNum, err := strconv.Atoi(PhoeniciaDigitalConfig.Config.Redis.Redis_port); err != nil {
			// In case the provided port is not an integer the API will consider it an error
			// And will Exist the process logging the error & issue
			PhoeniciaDigitalUtils.Log(fmt.Sprintf("Failed to implement Redis client | Invalid Port Number: %s", PhoeniciaDigitalConfig.Config.Redis.Redis_port))
			log.Fatalf("Failed to implement Redis client | Invalid Port Number: %s", PhoeniciaDigitalConfig.Config.Redis.Redis_port)
			return nil
		} else {
			if portNum < 0 || portNum > 65535 {
				// In case the provided port is not in range 0 <-> 65535 the API will consider it an error
				// And will Exist the process logging the error & issue
				PhoeniciaDigitalUtils.Log(fmt.Sprintf("Redis PORT: %s is OUT OF RANGE 0 --> 65535 | Change in ./config/.env", PhoeniciaDigitalConfig.Config.Redis.Redis_port))
				log.Fatalf("Redis PORT: %s is OUT OF RANGE 0 --> 65535 | Change in ./config/.env", PhoeniciaDigitalConfig.Config.Redis.Redis_port)
				return nil
			} else {
				// If all is good and checked append the specific Port to the MongoDB Connection string
				conStr += fmt.Sprintf(":%s", PhoeniciaDigitalConfig.Config.Redis.Redis_port)
			}
		}
	} else {
		conStr += ":6379"
	}

	return redis.NewClient(&redis.Options{
		Addr:     conStr,
		Password: PhoeniciaDigitalConfig.Config.Redis.Redis_password,
		DB:       0,
	})

}
