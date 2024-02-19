package utils

import (
	"time"
)

func (ds *DataStore) Set(key string, value interface{}, expiry int, nx bool, xx bool) (string, bool) {
	ds.Lock()
	defer ds.Unlock()

	if xx {
		_, exists := ds.data[key]
		if !exists {
			return "invalid", false
		}
	}
	if nx {
		_, exists := ds.data[key]
		if exists {
			return "invalid", false
		}
	}
	oldEntry, exists := ds.data[key]
	if exists && oldEntry != nil && oldEntry.timer != nil {
		oldEntry.timer.Stop()
	}
	e := &entry{value: value}
	if expiry > 0 {
		e.timer = time.AfterFunc(time.Duration(expiry)*time.Second, func() {
			ds.Lock()
			delete(ds.data, key)
			ds.Unlock()
		})
	}
	ds.data[key] = e

	return "success", true
}
