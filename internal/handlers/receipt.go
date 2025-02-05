package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessReceipt(ctx *gin.Context) {

	var receipt models.Receipt

	ctx.JSON(http.StatusOK, gin.H{"message": "Receipt processed successfully"})

	if err := ctx.ShouldBindJSON(&receipt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format"})
		return
	}

	receiptID := services.GenerateReceiptID()

	ctx.JSON(http.StatusOK, gin.H{"id": receiptID})

}
