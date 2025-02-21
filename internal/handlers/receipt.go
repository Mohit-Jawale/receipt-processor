package handlers

import (
	"log"
	"net/http"
	"receipt-processor/internal/models"
	"receipt-processor/internal/services"
	"receipt-processor/internal/storage"

	"github.com/gin-gonic/gin"
)

type ReceiptHandler struct {
	store storage.Storage
}

func NewReceiptHandler(store storage.Storage) *ReceiptHandler {
	return &ReceiptHandler{store: store}
}

func (h *ReceiptHandler) GetReceiptPoints(ctx *gin.Context) {
	id := ctx.Param("id")

	points, err := h.store.GetPoints(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"points": points})
}

func (h *ReceiptHandler) ProcessReceipt(ctx *gin.Context) {

	var receipt models.Receipt

	log.Println("Received a request to process a receipt")

	if err := ctx.ShouldBindJSON(&receipt); err != nil {

		log.Printf("The receipt is invalid. %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid."})
		return
	}

	if err := receipt.Validate(); err != nil {
		log.Printf("The receipt is invalid.%v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid."})
		return
	}

	receiptID := services.GenerateReceiptID()

	points := services.CalculatePoints(receipt)

	if err := h.store.StorePoints(receiptID, points); err != nil {
		log.Printf("Failed to Store points:%v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store points"})
	}

	log.Printf("Points stored successfully with ID: %s", receiptID)

	///ToDo
	// create a SHA 256 check if its in redis
	// If not then store recipt if then it duplicate return Error
	// inside this I can call the calcuation points service and store the point against the UUID

	if err := h.store.StoreReceipt(receiptID, receipt); err != nil {

		log.Printf("Failed to store receipt: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store receipt"})
		return

	}

	log.Printf("Receipt stored successfully with ID: %s", receiptID)

	ctx.JSON(http.StatusOK, gin.H{"id": receiptID})

}
