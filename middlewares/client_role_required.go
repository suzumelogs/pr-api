package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientRoleRequired(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Role not found in context",
		})
		c.Abort()
		return
	}

	if role != "client" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied. Client role required.",
		})
		c.Abort()
		return
	}

	c.Next()
}
