package http

import (
	"github.com/labstack/echo"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(credentialHandler *CredentialHandler, profileHandler *ProfileHandler, internalUsername, internalPassword string) *echo.Echo {

	engine := echo.New()

	// CORS
	// engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowAllOrigins: true,
	// 	AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "SVC_USER", "SVC_PASS"},
	// 	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH"},
	// }))

	engine.GET("/", Status)
	// engine.POST("/login", h.Login)
	// engine.GET("/private", h.Private, IsLoggedIn)
	// engine.GET("/admin", h.Private, IsLoggedIn, isAdmin)
	engine.GET("/version", Version)

	//Profile
	engine.POST("/create-Profile", profileHandler.CreateProfile)
	engine.GET("/list-Profile", profileHandler.GetListProfile)
	engine.GET("/get-Profile/:id", profileHandler.GetDetailProfile)
	engine.PUT("/update-Profile/:id", profileHandler.UpdateProfile)
	engine.DELETE("/delete-Profile/:id", profileHandler.DeleteProfile)

	//User
	engine.POST("/create-credential", credentialHandler.CreateCredential)

	return engine
}
