package test

import (
	"receipt-processor/internal/models"
	"receipt-processor/internal/services"
	"testing"
)

func TestCalculatePoints(t *testing.T) {

	receipt := models.Receipt{
		Retailer:     "Store123",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "14:30",
		Total:        "10.25",
		Items: []models.Item{
			{ShortDescription: "Example Item", Price: "5.00"},
		},
	}

	expectedPoints := 8 + 6 + 25 + 10 + 1

	actualPoints := services.CalculatePoints(receipt)

	if actualPoints != expectedPoints {
		t.Errorf("Expected %d points, got %d", expectedPoints, actualPoints)
	}

}
