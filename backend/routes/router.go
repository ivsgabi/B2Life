package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/health", health)
}
