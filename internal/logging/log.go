package logging

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Read-only FS or other error: fallback to stdout
		Logger = log.New(os.Stdout, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
		Logger.Println("warning: could not open app.log; using stdout logger:", err)
		return
	}
	Logger = log.New(file, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Println("Logger initialized")
}