package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"database/sql"
)

// CORSMiddleware sets the Access-Control-Allow-Origin header to allow CORS requests.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") 
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, nil)
			return
		}
		c.Next()
	}
}

// JSONContentTypeMiddleware sets the Content-Type header to application/json
func JSONContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}


// AuthMiddleware extracts the user ID from the JWT token and sets it in the context.
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"].(string)
			c.Set("userID", userID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AuthAdminMiddleware extracts the user ID from the JWT token, fetches the roles from the database, and sets them in the context.
func AuthAdminMiddleware(secretKey string, db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Parse token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract user ID, roles, and role types from token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["user_id"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}
			c.Set("userID", userID)

			// Extract roles from token claims
			roles, ok := claims["roles"].([]interface{})
			if ok {
				roleIDs := make([]string, len(roles))
				for i, role := range roles {
					roleIDs[i] = role.(string)
				}
				c.Set("roles", roleIDs)
			}

			// Extract role types from token claims
			roleTypes, ok := claims["roleTypes"].([]interface{})
			if ok {
				roleTypeStrings := make([]string, len(roleTypes))
				for i, roleType := range roleTypes {
					roleTypeStrings[i] = roleType.(string)
				}
				c.Set("roleTypes", roleTypeStrings)
			}

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// FormContentTypeMiddleware sets the Content-Type header to multipart/form-data
func FormContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "multipart/form-data")
		c.Next()
	}
}