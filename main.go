package main

import (
	"redis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	routes.Operation(r)
	r.Run()
}
