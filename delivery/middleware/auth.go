package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"online-store-application/utils/security"
	"strings"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var authHeader authHeader
		if err := c.Bind(&authHeader); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": fmt.Sprintf("unauthorized %v", err.Error()),
			})
		}

		// Substring
		token := strings.Replace(authHeader.AuthorizationHeader, "Bearer ", "", 1)
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "unauthorized",
			})
		}

		// Verifikasi Token
		claims, err := security.VerifyJwtToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": fmt.Sprintf("unauthorized %v", err.Error()),
			})
		}
		c.Set("claims", claims)
		return next(c)
	}
}

