package middleware

import (
	"encoding/base64"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ValidateAPIKey() gin.HandlerFunc {
	encodedApiKey := os.Getenv("API_KEY") // Asegúrate de establecer esta variable de entorno
	apiKey, err := base64.StdEncoding.DecodeString(encodedApiKey)
	if err != nil {
		panic("Error al decodificar la API Key: " + err.Error())
	}
	return func(c *gin.Context) {
		providedApiKey := c.GetHeader("X-API-Key")
		if providedApiKey != string(apiKey) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API key no válida o ausente"})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
