package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"receipt-processor/internal/handlers"
	"receipt-processor/internal/models"
	"receipt-processor/internal/services"
	"receipt-processor/internal/storage"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProcessReceipt(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	store := storage.NewInMemoryStorage()
	handler := handlers.NewReceiptHandler(store)
	router.POST("/receipts/process", handler.ProcessReceipt)

	receipt := `{"retailer":"Target","purchaseDate":"2022-01-01","purchaseTime":"13:01","items":[{"shortDescription":"Item 1","price":"5.00"}],"total":"10.00"}`

	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(([]byte(receipt))))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Excepted 200 ok got %d", resp.Code)
	}

}

func TestGetReceiptPoints(t *testing.T) {

	gin.SetMode(gin.TestMode)

	store := storage.NewInMemoryStorage()
	handler := handlers.NewReceiptHandler(store)

	router := gin.Default()
	router.GET("/receipts/:id/points", handler.GetReceiptPoints)

	receipt := models.Receipt{
		Retailer:     "Store123",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "14:30",
		Total:        "10.25",
		Items: []models.Item{
			{ShortDescription: "Example Item", Price: "5.00"},
		},
	}

	receiptID := services.GenerateReceiptID()
	store.StoreReceipt(receiptID, receipt)

	req, _ := http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", resp.Code)
	}

}
