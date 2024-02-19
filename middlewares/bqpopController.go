package middlewares

import (
	"fmt"
	"log"
	"redis/utils"
	"strconv"
)

func BQPopController(words []string, commands string) string {
	if len(words) < 3 {
		msg := "not enough arguments for BQPOP command"
		log.Println(msg)
		return msg
	} else if len(words) > 3 {
		msg := "enough arguments for BQPOP command"
		log.Println(msg)
		return msg
	}
	ds := utils.GetInstance()
	key := words[1]
	timeout, _ := strconv.Atoi(words[2])
	output, ok := ds.BQPOP(key, timeout)
	if !ok {
		log.Println("error while doing BQPOP command")
	}
	strValue := fmt.Sprintf("%v", output)
	return strValue
}
