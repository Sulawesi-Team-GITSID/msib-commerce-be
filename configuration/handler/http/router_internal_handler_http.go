package http

import (
	"github.com/labstack/echo"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(credentialHandler *CredentialHandler, profileHandler *ProfileHandler, gameHandler *GameHandler, internalUsername, internalPassword string) *echo.Echo {

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
	engine.POST("/create-profile", profileHandler.CreateProfile)
	engine.GET("/list-profile", profileHandler.GetListProfile)
	engine.GET("/get-profile/:id", profileHandler.GetDetailProfile)
	engine.PUT("/update-profile/:id", profileHandler.UpdateProfile)
	engine.DELETE("/delete-profile/:id", profileHandler.DeleteProfile)

	//User
	engine.POST("/create-credential", credentialHandler.CreateCredential)
	engine.GET("/list-credential", credentialHandler.GetListCredential)

	//Game
	engine.POST("/create-game", gameHandler.CreateGame)
	engine.GET("/list-game", gameHandler.GetListGame)
	engine.GET("/get-game/:id", gameHandler.GetDetailGame)
	engine.PUT("/update-game/:id", gameHandler.UpdateGame)
	engine.DELETE("/delete-game/:id", gameHandler.DeleteGame)

	return engine
}
