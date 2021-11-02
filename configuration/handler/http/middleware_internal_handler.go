package http

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
)

type JWThandler struct{}

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	// SigningKey: os.Getenv("JWT_SECRET_KEY"),
	SigningKey: []byte("secret"),
})

func isSeller(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isSeller := claims["seller"].(bool)
		if !isSeller {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

// func (handler *JWThandler) Private(echoCtx echo.Context) error {
// 	user := echoCtx.Get("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	name := claims["name"].(string)
// 	return echoCtx.String(http.StatusOK, "Welcome "+name+"!")
// }
