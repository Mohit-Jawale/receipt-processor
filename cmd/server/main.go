package main

import (
	"log"
	"receipt-processor/internal/handlers"
	"receipt-processor/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	store := storage.NewInMemoryStorage()
	receiptHandler := handlers.NewReceiptHandler(store)

	// ✅ Register API routes
	r.POST("/receipts/process", receiptHandler.ProcessReceipt)
	r.GET("/receipts/:id/points", receiptHandler.GetReceiptPoints)

	// ✅ Start server
	log.Println("🚀 Server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
