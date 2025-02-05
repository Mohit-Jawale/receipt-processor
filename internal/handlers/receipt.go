package handlers

import (
	"net/http"
	"receipt-processor/internal/models"
	"receipt-processor/internal/services"

	"github.com/gin-gonic/gin"
)

func ProcessReceipt(ctx *gin.Context) {

	var receipt models.Receipt

	if err := ctx.ShouldBindJSON(&receipt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format"})
		return
	}

	receiptID := services.GenerateReceiptID()

	ctx.JSON(http.StatusOK, gin.H{"id": receiptID})

}
