// File: `Server Logger File` source/utils/logger.go
package PhoeniciaDigitalUtils

import (
	"log"
	"os"
)

var _PhoeniciaDigitalLogger *log.Logger

// InitLogger initializes the logger and sets the log output to the given log file.
func init() {
	logFile, err := os.OpenFile("./Phoenicia-Digital.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error initializing logger:", err)
	}
	_PhoeniciaDigitalLogger = log.New(logFile, "- Time: ", log.Ldate|log.Ltime) // Include microseconds for more precision.
}

// Log logs the given message using the logger.
func Log(message string) {
	if _PhoeniciaDigitalLogger != nil {
		_PhoeniciaDigitalLogger.Printf("| LOG: %s", message)
	}
}
