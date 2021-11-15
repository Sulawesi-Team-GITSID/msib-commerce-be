package http

import (
	"github.com/labstack/echo"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.

func NewGinEngine(credentialHandler *CredentialHandler, profileHandler *ProfileHandler, gameHandler *GameHandler, voucherHandler *VoucherHandler, verificationHandler *VerificationHandler, Middlewarehandler *Middlewarehandler, reviewHandler *ReviewHandler, superAdminHandler *SuperAdminHandler, internalUsername, internalPassword string) *echo.Echo {
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

	//Superadmin
	engine.POST("/login-admin", superAdminHandler.LoginAdmin)
	engine.POST("/register", superAdminHandler.Register)
	//engine.GET("/get-user", superAdminHandler.GetProfile)

	//Profile
	engine.POST("/create-profile", profileHandler.CreateProfile)
	engine.GET("/list-profile", profileHandler.GetListProfile, IsLoggedIn, isSeller)
	engine.GET("/get-profile/:id", profileHandler.GetDetailProfile)
	engine.PUT("/update-profile/:id", profileHandler.UpdateProfile)
	engine.DELETE("/delete-profile/:id", profileHandler.DeleteProfile)

	//Credential
	engine.POST("/create-credential", credentialHandler.CreateCredential)
	engine.GET("/list-credential", credentialHandler.GetListCredential, IsLoggedIn)
	engine.POST("/login", credentialHandler.Login)
	engine.GET("/update-credential/:id", credentialHandler.UpdateCredential)
	engine.PUT("/reset-password/:id", credentialHandler.ForgotPassword)
	// engine.GET("/private", JWThandler.Private)

	//Game
	engine.POST("/create-game", gameHandler.CreateGame)
	engine.GET("/list-game", gameHandler.GetListGame)
	engine.GET("/list-genre", gameHandler.GetListGenre)
	engine.GET("/list-trend-game", gameHandler.GetListTrendGame)
	engine.GET("/get-game/:id", gameHandler.GetDetailGame)
	engine.PUT("/update-game/:id", gameHandler.UpdateGame)
	engine.DELETE("/delete-game/:id", gameHandler.DeleteGame)

	//Voucher
	engine.POST("/create-voucher", voucherHandler.CreateVoucher)
	engine.GET("/list-voucher", voucherHandler.GetListVoucher)
	engine.GET("/get-voucher/:id", voucherHandler.GetDetailVoucher)
	engine.PUT("/update-voucher/:id", voucherHandler.UpdateVoucher)
	engine.DELETE("/delete-voucher/:id", voucherHandler.DeleteVoucher)

	//Verification
	engine.POST("/create-verification", verificationHandler.CreateVerification)
	engine.GET("/list-verification", verificationHandler.GetListVerification)
	engine.GET("/get-verification/:id", verificationHandler.GetDetailVerification)
	engine.PUT("/update-verification/:id", verificationHandler.UpdateVerification)
	engine.DELETE("/delete-verification/:id", verificationHandler.DeleteVerification)
	engine.POST("/verify-mail", verificationHandler.Verify)

	//Review
	engine.POST("/create-review", reviewHandler.CreateReview)
	engine.GET("/list-review", reviewHandler.GetListReview)

	return engine
}
