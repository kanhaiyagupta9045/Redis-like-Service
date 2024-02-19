package utils

import (
	"errors"
)

var queueUpdated = make(chan string, 10)

func (ds *DataStore) QPUSH(key string, values ...interface{}) (string, error) {
	ds.Lock()
	defer ds.Unlock()

	queueEntry, exists := ds.data[key]
	if !exists {
		queueEntry = &entry{value: make([]interface{}, 0)}
		ds.data[key] = queueEntry
	}
	queueSlice, ok := queueEntry.value.([]interface{})
	if !ok {
		return "invalid data", errors.New("value associated with the key is not a queue")
	}

	queueEntry.value = append(queueSlice, values...)
	select {
	case queueUpdated <- key:
	default:
	}

	return "data inserted", nil
}
