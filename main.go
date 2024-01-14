package main

import (
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2/app"
	"github.com/stenstromen/textstorm/gui"
)

func main() {
	logFile, err := setupLogging()
	if err != nil {
		log.Fatalf("Unable to set up logging: %v", err)
	}
	defer logFile.Close()

	textStorm := app.New()
	gui.Gui(&textStorm)
	textStorm.Run()
}

func setupLogging() (*os.File, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	logDir := filepath.Join(homeDir, "Textstorm")
	err = os.MkdirAll(logDir, 0755)
	if err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile(filepath.Join(logDir, "app.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	// Redirect output to log file
	os.Stdout = logFile
	os.Stderr = logFile
	log.SetOutput(logFile)

	return logFile, nil
}
