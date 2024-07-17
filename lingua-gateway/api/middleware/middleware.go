package middleware

import (
	"fmt"
	"net/http"

	"gateway-service/api/token"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		valid, err := token.ValidateToken(authHeader)
		if err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		claims, err := token.ExtractClaim(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims", "details": err.Error()})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		jwtClaims := claims.(jwt.MapClaims)
		sub := jwtClaims["role"].(string)

		obj := c.FullPath()
		act := c.Request.Method
		fmt.Println(sub, obj, act)

		allowed, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during authorization"})
			return
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
