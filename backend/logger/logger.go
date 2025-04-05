package logger

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func InitLogger() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.Println("Logger initialized")
}

func Info(message string) {
	logger.Println("INFO: " + message)
}

func InfOf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	logger.Println("INFO: " + message)
}

func Error(message string) {
	logger.Println("ERROR: " + message)
}
