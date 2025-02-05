package handlers

import (
	"net/http"
	"receipt-processor/internal/models"
	"receipt-processor/internal/services"
	"receipt-processor/internal/storage"

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

	points := services.CalculatePoints(receipt)

	ctx.JSON(http.StatusOK, models.PointsResponse{Points: points})
}

func (h *ReceiptHandler) ProcessReceipt(ctx *gin.Context) {

	var receipt models.Receipt

	if err := ctx.ShouldBindJSON(&receipt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format"})
		return
	}

	receiptID := services.GenerateReceiptID()

	ctx.JSON(http.StatusOK, gin.H{"id": receiptID})

}
