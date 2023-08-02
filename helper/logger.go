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

var date = time.Now().Format("2006-01-02")

func initLogger() *log.Logger {
	file, err := os.OpenFile(date+logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}

	logger := log.New(file, "", log.Ldate|log.Ltime)

	return logger
}

func LogOTPAction(action, otp, user, message string) {
	logger := initLogger()

	// Log the OTP action with the provided timestamp, OTP, user, and failure message.
	logMsg := fmt.Sprintf("%s %s - User: %s, OTP: %s, Message: %s", date, action, user, otp, message)
	logger.Println(logMsg)
}
