package storage

import (
	"errors"
	"log"
	"receipt-processor/internal/models"
	"sync"
)

type Storage interface {
	StorageReciept
	StorageUUID
}

type StorageReciept interface {
	StoreReceipt(id string, receipt models.Receipt) error
	GetReceipt(id string) (models.Receipt, error)
}

type StorageUUID interface {
	StorePoints(id string, points int64) error
	GetPoints(id string) (int64, error)
}

type InMemoryStorage struct {
	sync.RWMutex
	data  map[string]models.Receipt
	score map[string]int64
}

func NewInMemoryStorage() *InMemoryStorage {

	return &InMemoryStorage{
		data:  make(map[string]models.Receipt),
		score: make(map[string]int64),
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

func (s *InMemoryStorage) StorePoints(id string, points int64) error {

	if id == "" {
		log.Println("Attempted to store receipt with empty ID")
		return errors.New("receipt ID cannot be empty")
	}
	if points < 0 {
		log.Println("Point has to be positive something is wrong")
		return errors.New("problem with Point calculation")

	}

	s.Lock()
	defer s.Unlock()
	s.score[id] = points

	log.Println("UUID store successfully with points")
	return nil

}

func (s *InMemoryStorage) GetPoints(id string) (int64, error) {

	if id == "" {
		log.Println("Attempted to store receipt with empty ID")
		return 0, errors.New("receipt ID cannot be empty")
	}

	s.RLock()
	defer s.RUnlock()

	points, exists := s.score[id]

	if !exists {
		log.Printf("Receipt ID %s not found", id)
		return 0, errors.New("receipt not found")

	}

	return points, nil

}
