package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, err string) {
	c.JSON(http.StatusNotFound, gin.H{"error": err})
}
