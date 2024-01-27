package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ValidateAPIKey() gin.HandlerFunc {
	apiKey := os.Getenv("API_KEY") // Asegúrate de establecer esta variable de entorno
	return func(c *gin.Context) {
		providedApiKey := c.GetHeader("X-API-Key")
		if providedApiKey != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API key no válida o ausente"})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
