package http

import (
	"net/http"
	"os"

	"github.com/labstack/echo"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
})

func isSeller(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["seller"].(bool)
		if !isAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
func (handler *CredentialHandler) Private(echoCtx echo.Context) error {
	user := echoCtx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return echoCtx.String(http.StatusOK, "Welcome "+name+"!")
}
