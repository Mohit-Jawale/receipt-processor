package storage

import (
	"receipt-processor/internal/models"
	"sync"
)

type Storage interface {
	StoreReceipt(id string, receipt models.Receipt)
	GetReceipt(id string) (models.Receipt, bool)
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

func (s *InMemoryStorage) StorageReceipt(id string, receipt models.Receipt) {

	s.Lock()
	defer s.Unlock()
	s.data[id] = receipt

}

func (s *InMemoryStorage) GetReceipt(id string) (models.Receipt, bool) {

	s.RLock()
	defer s.RUnlock()
	receipt, exists := s.data[id]
	return receipt, exists

}
