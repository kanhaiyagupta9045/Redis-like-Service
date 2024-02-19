package routes

import (
	"redis/controllers"

	"github.com/gin-gonic/gin"
)

func Operation(router *gin.Engine) {
	router.POST("/operation", controllers.RunCommand())
}
