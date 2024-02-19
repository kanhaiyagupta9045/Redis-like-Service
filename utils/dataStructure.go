package utils

import (
	"sync"
	"time"
)

type entry struct {
	value interface{}
	timer *time.Timer
}

type DataStore struct {
	sync.Mutex
	data map[string]*entry
}

var instance *DataStore
var once sync.Once

func GetInstance() *DataStore {
	once.Do(func() {
		instance = &DataStore{
			data: make(map[string]*entry),
		}
	})
	return instance
}
