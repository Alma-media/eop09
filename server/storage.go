package server

import (
	"sync"

	"github.com/Alma-media/eop09/proto"
)

// Storage is a port storage.
type Storage struct {
	mutex sync.RWMutex
	data  map[string]*proto.Port
}

// NewStorage creates a new in-memory port storage.
func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]*proto.Port),
	}
}

// Save a single port.
func (storage *Storage) Save(id string, port *proto.Port) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	storage.data[id] = port
}

// Each calls provided callback function for each available pair id/port.
func (storage *Storage) Each(fn func(key string, port *proto.Port) error) error {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	for key, value := range storage.data {
		if err := fn(key, value); err != nil {
			return err
		}
	}

	return nil
}
