package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoleRequired(c *gin.Context) {

	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Role not found in context",
		})
		c.Abort()
		return
	}
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied. Admin role required.",
		})
		c.Abort()
		return
	}

	c.Next()
}