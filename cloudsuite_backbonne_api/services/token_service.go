package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type TokenService struct {
	LogFile string
}

type TokenLog struct {
	Date               string `json:"date"`
	ProcessingStatus   string `json:"processing_status"`
	ProcessingDuration int64  `json:"processing_duration"` // in milliseconds
}

func NewTokenService(logFile string) *TokenService {
	return &TokenService{LogFile: logFile}
}

func (ts *TokenService) generateToken() string {
	randomData := fmt.Sprintf("random-%d", rand.Int63())
	hash := sha256.Sum256([]byte(randomData))
	return hex.EncodeToString(hash[:])
}

func (ts *TokenService) ProcessToken() error {
	success := false

	for !success {
		startTime := time.Now()

		success = rand.Float32() > 0.5
		token := ts.generateToken()

		duration := time.Since(startTime).Milliseconds()

		status := "success"
		if !success {
			status = "failure"
		}

		logEntry := TokenLog{
			Date:               time.Now().Format(time.RFC3339),
			ProcessingStatus:   status,
			ProcessingDuration: duration,
		}

		if err := LogResult(ts.LogFile, logEntry); err != nil {
			log.Printf("Failed to log token processing result: %v", err)
			return err
		}

		if !success {
			log.Println("Retrying token generation due to failure...")
		} else {
			log.Printf("Token processed: %s | Status: %s | Duration: %d ms", token, status, duration)
		}
	}
	return nil
}
