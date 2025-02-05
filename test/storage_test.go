package test

import (
	"receipt-processor/internal/models"
	"receipt-processor/internal/storage"
	"testing"
)

func TestStoreAndRetirveReceipt(t *testing.T) {

	store := storage.NewInMemoryStorage()

	receipt := models.Receipt{
		Retailer:     "Test Store",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:00",
		Total:        "10.00",
		Items: []models.Item{
			{"Item 1", "5.00"},
		},
	}

	store.StorageReceipt("test-id", receipt)

	retrieved, exists := store.GetReceipt("test-id")

	if !exists {
		t.Errorf("Receipt not found")
	}

	if retrieved.Retailer != "Test Store" {
		t.Errorf("Expected retailer 'Test Store', got '%s'", retrieved.Retailer)
	}

}

func TestRetrieveNonExistentReceipt(t *testing.T) {
	store := storage.NewInMemoryStorage()

	_, exists := store.GetReceipt("non-existent-id")

	if exists {
		t.Errorf("Expected no receipt, but one was found")
	}
}
