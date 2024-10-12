package services

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateLogDirectory() error {
	return os.MkdirAll("logs", 0755)
}
func LogResult(logFile string, entry TokenLog) error {

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	logEntryJSON, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("failed to marshal log entry: %w", err)
	}

	if _, err := file.WriteString(string(logEntryJSON) + "\n"); err != nil {
		return fmt.Errorf("failed to write log entry to file: %w", err)
	}

	return nil
}
