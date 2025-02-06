package services

import (
	"math"
	"receipt-processor/internal/models"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateReceiptID() string {
	return uuid.New().String()
}

func CalculatePoints(receipt models.Receipt) int64 {
	points := 0

	// 1. Retailer name points (only alphanumeric characters count)
	reg := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(reg.FindAllString(receipt.Retailer, -1))

	// 2. Round dollar total (50 points)
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// 3. Multiple of 0.25 (25 points)
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 4. Every two items = 5 points
	points += (len(receipt.Items) / 2) * 5

	// 5. Description length multiple of 3
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)

			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6. Odd day bonus (6 points)
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// 7. Purchase time bonus (10 points if between 2 PM - 4 PM)
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 || purchaseTime.Hour() == 15 {
		points += 10
	}

	return int64(points)
}
