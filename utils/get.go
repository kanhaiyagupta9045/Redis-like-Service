package utils

func (ds *DataStore) Get(key string) (interface{}, bool) {
	ds.Lock()
	defer ds.Unlock()
	e, exists := ds.data[key]
	if !exists {
		return "not found", false
	}
	return e.value, true
}
