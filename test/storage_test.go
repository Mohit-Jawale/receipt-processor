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

	err := store.StoreReceipt("test-id", receipt)
	if err != nil {
		t.Errorf("Unexpected error while storing receipt: %v", err)
	}

	retrieved, err := store.GetReceipt("test-id")

	if err != nil {
		t.Errorf("Unexpected error while retrieving receipt: %v", err)
	}

	if retrieved.Retailer != "Test Store" {
		t.Errorf("Expected retailer 'Test Store', got '%s'", retrieved.Retailer)
	}

}

func TestRetrieveNonExistentReceipt(t *testing.T) {
	store := storage.NewInMemoryStorage()

	_, err := store.GetReceipt("non-existent-id")

	if err == nil {
		t.Errorf("Expected no receipt, but one was found")
	}
}

func TestStoreWithEmptyID(t *testing.T) {
	store := storage.NewInMemoryStorage()

	receipt := models.Receipt{
		Retailer:     "Test Store",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:00",
		Total:        "10.00",
		Items:        []models.Item{},
	}

	err := store.StoreReceipt("", receipt)
	if err == nil {
		t.Errorf("Expected error for empty receipt ID, but got none")
	}
}
