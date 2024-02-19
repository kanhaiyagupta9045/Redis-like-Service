package controllers

import (
	"net/http"
	"redis/helpers"
	"redis/models"

	"github.com/gin-gonic/gin"
)
func RunCommand() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmd models.Command
		if err := c.BindJSON(&cmd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		body := cmd.Operation

		output := helpers.Distinguish(body)
		c.JSON(http.StatusOK, gin.H{"value": output})
	}
}
