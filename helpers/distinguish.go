package helpers

import (
	middlewares "redis/middlewares"
	"strings"
)

func Distinguish(commands string) string {
	operator, err := Validate(commands)
	if err != nil {
		return err.Error()
	}
	words := strings.Fields(commands)
	if operator == "SET" {

		output := middlewares.SetController(words, commands)
		return output

	} else if operator == "GET" {

		output := middlewares.GetController(words, commands)
		return output
	} else if operator == "QPUSH" {

		output := middlewares.QpushController(words, commands)
		return output
	} else if operator == "QPOP" {

		output := middlewares.QpopController(words, commands)
		return output
	} else if operator == "BQPOP" {
		output := middlewares.BQPopController(words, commands)
		return output
	}

	return "invalid command"

}
