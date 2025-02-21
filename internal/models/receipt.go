package models

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

// Receipt represents a store receipt
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item represents an item on the receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Points
type PointsResponse struct {
	Points int64 `json:"points"`
}

func isValidRegex(value, pattern string) bool {
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

func isValidPrice(priceStr string) error {
	if !isValidRegex(priceStr, `^\d+\.\d{2}$`) {
		return fmt.Errorf("price must be in the format X.XX with two decimal places")
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return fmt.Errorf("invalid price format: %w", err)
	}
	if price < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	return nil
}

func (i *Item) Validate() error {

	if !isValidRegex(i.ShortDescription, `^[\w\s\-]+$`) {
		return fmt.Errorf("short description must be alphanumeric, spaces, or hyphens")
	}

	if err := isValidPrice(i.Price); err != nil {
		return fmt.Errorf("invalid item price: %w", err)
	}

	return nil
}
func (r *Receipt) Validate() error {

	if !isValidRegex(r.Retailer, `^[\w\s\-&]+$`) {
		return fmt.Errorf("retailer must only contain letters, numbers, spaces, hyphens, or ampersands")
	}

	parsedDate, err := time.Parse("2006-01-02", r.PurchaseDate)
	if err != nil {
		return fmt.Errorf("invalid purchase date format (YYYY-MM-DD): %w", err)
	}
	if parsedDate.After(time.Now()) {
		return fmt.Errorf("purchase date cannot be in the future")
	}

	if !isValidRegex(r.PurchaseTime, `^(?:[01]?\d|2[0-3]):[0-5]\d$`) {
		return fmt.Errorf("invalid purchase time format (HH:MM, 24-hour)")
	}

	if len(r.Items) == 0 {
		return fmt.Errorf("receipt must contain at least one item")
	}

	var itemSum float64
	for _, item := range r.Items {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("invalid item: %w", err)
		}
		itemPrice, _ := strconv.ParseFloat(item.Price, 64)
		itemSum += itemPrice
	}
	if err := isValidPrice(r.Total); err != nil {
		return fmt.Errorf("invalid total: %w", err)
	}
	totalValue, _ := strconv.ParseFloat(r.Total, 64)
	roundeditemSum := math.Trunc(itemSum*100) / 100
	if totalValue != roundeditemSum {
		return fmt.Errorf("total does not match sum of item prices (expected: %.2f, got: %.2f)", roundeditemSum, totalValue)
	}

	return nil
}
