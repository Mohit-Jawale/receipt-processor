package storage

import (
	"errors"
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
		return errors.New("receipt ID cannot be empty")
	}
	if receipt.Retailer == "" || receipt.Total == "" {
		return errors.New("invalid receipt: missing required fields")
	}
	s.Lock()
	defer s.Unlock()
	s.data[id] = receipt

	return nil

}

func (s *InMemoryStorage) GetReceipt(id string) (models.Receipt, error) {

	if id == "" {
		return models.Receipt{}, errors.New("receipt ID cnnot be empty")
	}

	s.RLock()
	defer s.RUnlock()
	receipt, exists := s.data[id]
	if !exists {
		return models.Receipt{}, errors.New("receipt not found")

	}
	return receipt, nil

}
