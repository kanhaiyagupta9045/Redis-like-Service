package middlewares

import (
	"fmt"
	"log"
	"redis/utils"
)

func QpopController(words []string, commands string) string {
	if len(words) < 2 {
		log.Println("Error: not enough arguments for QPOP command")
	}
	key := words[1]

	ds := utils.GetInstance()
	output, _ := ds.QPOP(key)

	strValue := fmt.Sprintf("%v", output)
	return strValue
}
