package middlewares

import (
	"log"
	"redis/utils" 
	"strconv"
	"strings"
)

func SetController(words []string, commands string) string {
	if len(words) < 3 {
		msg := "not enough arguments for SET command"
		log.Println(msg)
		return msg
	}

	key := words[1]
	value := words[2]
	expiry := 0 
	nx := strings.Contains(commands, "NX")
	xx := strings.Contains(commands, "XX")
	if strings.Contains(commands, "EX") {
		if len(words) < 5 {
			msg := "not enough arguments for EX option"
			log.Println(msg)
			return msg
		}
		var err error
		expiry, err = strconv.Atoi(words[4])
		if err != nil {
			log.Printf("Error parsing expiry: %v\n", err)
			return "Error parsing expiry"
		}
	}

	ds := utils.GetInstance()
	response, success := ds.Set(key, value, expiry, nx, xx)
	if !success {
		return "SET command failed"
	}
	return response
}
