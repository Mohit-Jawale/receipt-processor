package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProcessReceipt(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Post("/receipts/process", handlers.ProcessReceipt)

	receipt := `{"retailer":"Target","purchaseDate":"2022-01-01","purchaseTime":"13:01","items":[{"shortDescription":"Item 1","price":"5.00"}],"total":"10.00"}`

	req, _ := http.NewRequest("Post", "/receipts/process", bytes.NewBuffer(([]byte(receipt))))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServerHTTP(resp, http.Request)

	if resp.Code != http.StatusOK {
		t.Errorf("Excepted 200 ok got %d", resp.Code)
	}

}
