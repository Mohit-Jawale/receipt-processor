package services

import "github.com/google/uuid"

// GenerateReceiptID creates a unique ID for a receipt
func GenerateReceiptID() string {
	return uuid.New().String()
}
