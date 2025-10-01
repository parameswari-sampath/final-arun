package logger

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func Init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	if InfoLogger == nil {
		Init()
	}
	InfoLogger.Println(message)
}

func Error(message string) {
	if ErrorLogger == nil {
		Init()
	}
	ErrorLogger.Println(message)
}
