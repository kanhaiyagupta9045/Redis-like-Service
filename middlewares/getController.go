package middlewares

import (
	"fmt"
	"log"
	"redis/utils"
)

func GetController(words []string, commands string) string {
	if len(words) < 2 {
		msg := "not enough arguments for GET command"
		log.Println(msg)
		return msg
	} else if len(words) > 2 {
		msg := "enough arguments for GET command"
		log.Println(msg)
		return msg
	}

	key := words[1]
	ds := utils.GetInstance()
	value, found := ds.Get(key)
	if !found {
		msg := fmt.Sprintf("Key '%s' not found in DataStore", key)
		log.Println(msg)
		return "key not found"
	}

	log.Printf("Retrieved for key '%s': %v\n", key, value)
	strValue := fmt.Sprintf("%v", value)
	return strValue
}
