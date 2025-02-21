package main

import (
	"log"
	"receipt-processor/internal/handlers"
	"receipt-processor/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	store := storage.NewInMemoryStorage()

	receiptHandler := handlers.NewReceiptHandler(store)

	server.POST("/receipts/process", receiptHandler.ProcessReceipt)
	server.GET("/receipts/:id/points", receiptHandler.GetReceiptPoints)

	log.Println("Server running on port 8080")
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
