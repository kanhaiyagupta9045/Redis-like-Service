package utils

func (ds *DataStore) QPOP(key string) (interface{}, bool) {
	ds.Lock()
	defer ds.Unlock()

	queueEntry, exists := ds.data[key]
	if !exists {
		return nil, false
	}
	queue, ok := queueEntry.value.([]interface{})
	if !ok || len(queue) == 0 {

		return "null", false
	}

	size := len(queue)
	lastElement := queue[size-1]
	queue = queue[:size-1]
	queueEntry.value = queue
	return lastElement, true
}
