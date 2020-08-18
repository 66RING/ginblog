package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Type", "Authorization", "Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})
}
