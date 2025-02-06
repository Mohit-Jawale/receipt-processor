package handlers

import (
	"log"
	"net/http"
	"receipt-processor/internal/models"
	"receipt-processor/internal/services"
	"receipt-processor/internal/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReceiptHandler struct {
	Storage storage.Storage
}

func NewReceiptHandler(store storage.Storage) *ReceiptHandler {
	return &ReceiptHandler{Storage: store}
}

func (h *ReceiptHandler) GetReceiptPoints(ctx *gin.Context) {
	id := ctx.Param("id")

	receipt, err := h.Storage.GetReceipt(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No receipt found"})
		return
	}

	points := strconv.Itoa(services.CalculatePoints(receipt))

	ctx.JSON(http.StatusOK, models.PointsResponse{Points: points})
}

func (h *ReceiptHandler) ProcessReceipt(ctx *gin.Context) {

	var receipt models.Receipt

	log.Println("Received a request to process a receipt")

	if err := ctx.ShouldBindJSON(&receipt); err != nil {

		log.Printf("Invalid receipt format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format"})
		return
	}

	receiptID := services.GenerateReceiptID()

	if err := h.Storage.StoreReceipt(receiptID, receipt); err != nil {

		log.Printf("Failed to store receipt: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store receipt"})
		return

	}

	log.Printf("Receipt stored successfully with ID: %s", receiptID)

	ctx.JSON(http.StatusOK, gin.H{"id": receiptID})

}
