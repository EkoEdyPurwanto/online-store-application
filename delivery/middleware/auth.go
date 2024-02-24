package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"online-store-application/utils/security"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Unauthorized",
			})
		}

		// Check if the Authorization header starts with "Bearer"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Invalid token format",
			})
		}

		// Extract token from Authorization header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify Token
		claims, err := security.VerifyJwtToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": fmt.Sprintf("Unauthorized: %v", err.Error()),
			})
		}

		// Set claims in the context for later use
		c.Set("claims", claims)
		return next(c)
	}
}
