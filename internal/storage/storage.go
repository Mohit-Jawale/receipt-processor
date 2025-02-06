package storage

import (
	"errors"
	"log"
	"receipt-processor/internal/models"
	"sync"
)

type Storage interface {
	StoreReceipt(id string, receipt models.Receipt) error
	GetReceipt(id string) (models.Receipt, error)
}

type InMemoryStorage struct {
	sync.RWMutex
	data map[string]models.Receipt
}

func NewInMemoryStorage() *InMemoryStorage {

	return &InMemoryStorage{
		data: make(map[string]models.Receipt),
	}
}

func (s *InMemoryStorage) StoreReceipt(id string, receipt models.Receipt) error {

	if id == "" {
		log.Println("Attempted to store receipt with empty ID")
		return errors.New("receipt ID cannot be empty")
	}
	if receipt.Retailer == "" || receipt.Total == "" {
		log.Println("Attempted to store receipt with missing fields")
		return errors.New("invalid receipt: missing required fields")
	}
	s.Lock()
	defer s.Unlock()
	s.data[id] = receipt
	log.Printf("Receipt successfully stored with ID: %s", id)
	return nil

}

func (s *InMemoryStorage) GetReceipt(id string) (models.Receipt, error) {

	if id == "" {
		log.Println("Attempted to retrieve receipt with empty ID")
		return models.Receipt{}, errors.New("receipt ID cnnot be empty")
	}

	s.RLock()
	defer s.RUnlock()
	receipt, exists := s.data[id]
	if !exists {
		log.Printf("Receipt ID %s not found", id)
		return models.Receipt{}, errors.New("receipt not found")

	}
	log.Printf("Successfully retrieved receipt ID: %s", id)
	return receipt, nil

}
