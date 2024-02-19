package middlewares

import (
	"log"
	"redis/utils"
)

func QpushController(words []string, command string) string {
	if len(words) < 3 {
		msg := "Error: not enough arguments for QPUSH command"
		log.Println(msg)
		return msg
	}

	key := words[1]
	var values []interface{}
	for _, v := range words[2:] {
		values = append(values, v)
	}
	ds := utils.GetInstance()
	response, err := ds.QPUSH(key, values...)
	if err != nil {
		log.Printf("Error in QPUSH operation: %v\n", err)
		return err.Error()
	}
	return response
}
