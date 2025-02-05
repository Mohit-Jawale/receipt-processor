package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessReceipt(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"message": "Receipt processed successfully"})

}
