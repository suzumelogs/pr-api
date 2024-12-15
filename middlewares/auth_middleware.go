package middlewares

import (
	"net/http"
	"os"
	"pr-api/initializers"
	"pr-api/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginRequired(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header missing",
		})
		c.Abort()
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Authorization header format",
		})
		c.Abort()
		return
	}

	tokenString := authHeaderParts[1]

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not valid",
		})
		c.Abort()
		return
	}

	userID := claims["sub"].(float64)
	var user models.User
	result := initializers.DB.First(&user, int(userID))
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Set("role", user.Role)

	c.Next()
}
