package helper

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	logFilePath = "otp_logs.log"
)

func initLogger() *log.Logger {
	// Open the log file in append mode. Create the file if it doesn't exist.
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}

	// Create a logger with a custom log format and set the output to the log file.
	logger := log.New(file, "", log.Ldate|log.Ltime)

	return logger
}

// func LogOTPAction(action, otp string, user uint) {
// 	date := time.Now().Format("2006-01-02")
// 	logger := initLogger()

// 	// Log the OTP action with the provided timestamp, OTP, and action.
// 	logMsg := fmt.Sprintf("%s user_id=%d %s - %s", date, user, action, otp)
// 	logger.Println(logMsg)
// }

func LogOTPAction(action, otp, user, message string) {
	logger := initLogger()
	date := time.Now().Format("2006-01-02")
	// Log the OTP action with the provided timestamp, OTP, user, and failure message.
	logMsg := fmt.Sprintf("%s %s - User: %s, OTP: %s, Message: %s", date, action, user, otp, message)
	logger.Println(logMsg)
}
