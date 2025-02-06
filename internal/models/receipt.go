package models

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
	Points int64
}

func (r *Receipt) Validate() error {
	// 1. Retailer: Basic check for non-empty and alphanumeric retailer name
	if strings.TrimSpace(r.Retailer) == "" {
		return fmt.Errorf("retailer cannot be empty")
	}

	// 2. Total: Check for non-negative and valid numeric format
	total, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		return fmt.Errorf("invalid total format: %w", err)
	}
	if total < 0 {
		return fmt.Errorf("total cannot be negative")
	}

	// 3. Items: Ensure at least one item exists
	if len(r.Items) == 0 {
		return fmt.Errorf("receipt must contain at least one item")
	}
	for _, item := range r.Items {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("invalid item: %w", err)
		}
	}

	// 4. PurchaseDate: Validate format (YYYY-MM-DD)
	parsedDate, err := time.Parse("2006-01-02", r.PurchaseDate)
	if err != nil {
		return fmt.Errorf("invalid purchase date format (YYYY-MM-DD): %w", err)
	}
	// Check if the date is in the future
	if parsedDate.After(time.Now()) {
		return fmt.Errorf("purchase date cannot be in the future")
	}

	// 5. PurchaseTime: Validate 24-hour format (HH:MM)
	if match, _ := regexp.MatchString(`^(?:[01]?\d|2[0-3]):[0-5]\d$`, r.PurchaseTime); !match {
		return fmt.Errorf("invalid purchase time format (HH:MM, 24-hour)")
	}

	return nil // Validation successful
}

func (i *Item) Validate() error {

	// Validate Price format
	price, err := strconv.ParseFloat(i.Price, 64)
	if err != nil {
		return fmt.Errorf("invalid price format: %w", err)
	}
	if price < 0 {
		return fmt.Errorf("price cannot be negative")
	}

	return nil // Validation successful
}
