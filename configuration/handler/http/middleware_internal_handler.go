package http

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
)

type Middlewarehandler struct{}

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

// func mailing(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Get("user").(*jwt.Token)
// 		err := Sendmail(user)
// 		if err == nil {
// 			return c.String(http.StatusOK, "Welcome ")
// 		}
// 		return next(c)
// 	}
// }

// func (handler *Middlewarehandler) mailHandler(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	Sendmail(user)
// 	return c.String(http.StatusOK, "Mail sent!")
// }

// func (handler *Middlewarehandler) Private(echoCtx echo.Context) error {
// 	user := echoCtx.Get("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	name := claims["name"].(string)
// 	return echoCtx.String(http.StatusOK, "Welcome "+name+"!")
// }
