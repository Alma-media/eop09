package storage

import (
	"sync"

	"github.com/Alma-media/eop09/proto"
)

// InMemory is in-memory port storage.
type InMemory struct {
	mutex sync.RWMutex
	data  map[string]*proto.Port
}

// NewInMemory creates a new in-memory port storage.
func NewInMemory() *InMemory {
	return &InMemory{
		data: make(map[string]*proto.Port),
	}
}

// Save a single port.
func (storage *InMemory) Save(id string, port *proto.Port) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	storage.data[id] = port
}

// Each calls provided callback function for each available pair id/port.
func (storage *InMemory) Each(fn func(key string, port *proto.Port) error) error {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	for key, value := range storage.data {
		if err := fn(key, value); err != nil {
			return err
		}
	}

	return nil
}
