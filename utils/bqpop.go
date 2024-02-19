package utils

import (
	"time"
)

func (ds *DataStore) BQPOP(key string, timeout int) (interface{}, bool) {
	ds.Lock()

	queueEntry, exists := ds.data[key]
	if !exists || len(queueEntry.value.([]interface{})) == 0 {
		if timeout == 0 {
			ds.Unlock()
			return nil, false
		}

		timer := time.NewTimer(time.Duration(timeout) * time.Second)
		defer timer.Stop()

		ds.Unlock()

		for {
			select {
			case updatedKey := <-queueUpdated:
				if updatedKey == key {
					ds.Lock()
					queueEntry, exists = ds.data[key]
					if exists && len(queueEntry.value.([]interface{})) > 0 {
						value := queueEntry.value.([]interface{})[0]
						queueEntry.value = queueEntry.value.([]interface{})[1:]
						ds.Unlock()
						return value, true
					}
					ds.Unlock()
				}
			case <-timer.C:
				return nil, false
			}
		}
	}

	value := queueEntry.value.([]interface{})[0]
	queueEntry.value = queueEntry.value.([]interface{})[1:]
	ds.Unlock()
	return value, true
}
