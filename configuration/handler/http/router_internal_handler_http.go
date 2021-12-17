package http

import (
	"github.com/labstack/echo"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.

func NewGinEngine(credentialHandler *CredentialHandler, profileHandler *ProfileHandler, gameHandler *GameHandler, voucherHandler *VoucherHandler, verificationHandler *VerificationHandler, Middlewarehandler *Middlewarehandler, reviewHandler *ReviewHandler, superAdminHandler *SuperAdminHandler, shopHandler *ShopHandler, genreHandler *GenreHandler, tagsHandler *TagsHandler, tags_detailHandler *Tags_detailHandler, fileHandler *FileHandler, wishlistHandler *WishlistHandler, internalUsername, internalPassword string) *echo.Echo {
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
	engine.GET("/get-credential/:id", credentialHandler.GetListCredential)
	engine.POST("/login", credentialHandler.Login)
	engine.GET("/verify-account/:id", credentialHandler.UpdateCredentialVerify)
	engine.POST("/search-email", credentialHandler.EmailSearch)
	engine.PUT("/reset-password/:id", credentialHandler.ForgotPassword)
	// engine.GET("/private", JWThandler.Private)

	//Shop
	engine.POST("/create-shop", shopHandler.CreateShop)
	engine.GET("/list-shop", shopHandler.GetListShop)
	engine.GET("/get-shop/:id", shopHandler.GetDetailShop)
	engine.PUT("/update-shop/:id", shopHandler.UpdateShop)
	engine.DELETE("/delete-shop/:id", shopHandler.DeleteShop)

	//Game
	engine.POST("/create-game", gameHandler.CreateGame)
	engine.GET("/list-game", gameHandler.GetListGame)
	engine.GET("/list-game-shop/:id", gameHandler.GetListGameShop)
	engine.GET("/sort-game/:order/:sort", gameHandler.SortGame)
	engine.GET("/sort-game-shop/:order/:sort/:id", gameHandler.SortGameByShop)
	engine.GET("/list-trend-game", gameHandler.GetListTrendGame)
	engine.GET("/get-game/:id", gameHandler.GetDetailGame)
	engine.PUT("/update-game/:id", gameHandler.UpdateGame)
	engine.DELETE("/delete-game/:id", gameHandler.DeleteGame)

	//Voucher
	engine.POST("/create-voucher", voucherHandler.CreateVoucher)
	engine.GET("/list-voucher", voucherHandler.GetListVoucher)
	engine.GET("/list-voucher-shop/:id", voucherHandler.GetListVoucherShop)
	engine.GET("/sort-voucher/:order/:sort", voucherHandler.SortVoucher)
	engine.GET("/sort-voucher-shop/:order/:sort/:id", voucherHandler.SortVoucherByShop)
	engine.GET("/get-voucher/:id", voucherHandler.GetDetailVoucher)
	engine.PUT("/update-voucher/:id", voucherHandler.UpdateVoucher)
	engine.DELETE("/delete-voucher/:id", voucherHandler.DeleteVoucher)

	//Genre
	engine.POST("/create-genre", genreHandler.CreateGenre)
	engine.GET("/list-genre", genreHandler.GetListGenre)
	engine.GET("/get-genre/:id", genreHandler.GetDetailGenre)
	engine.PUT("/update-genre/:id", genreHandler.UpdateGenre)
	engine.DELETE("/delete-genre/:id", genreHandler.DeleteGenre)

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

	//Tags
	engine.POST("/create-tags", tagsHandler.CreateTags)
	engine.GET("/list-tags", tagsHandler.GetListTags)
	engine.GET("/get-tags/:id", tagsHandler.GetDetailTags)
	engine.PUT("/update-tags/:id", tagsHandler.UpdateTags)
	engine.DELETE("/delete-tags/:id", tagsHandler.DeleteTags)

	//Tags_detail
	engine.POST("/create-gametags", tags_detailHandler.CreateTags_detail)
	engine.GET("/list-gametags", tags_detailHandler.GetListTags_detail)
	// engine.GET("/get-gametags/:id", tags_detailHandler.GetDetailTags_detail)
	engine.GET("/get-gametags/:id", tags_detailHandler.GetGameTags)
	engine.PUT("/update-gametags/:id", tags_detailHandler.UpdateTags_detail)
	engine.DELETE("/delete-gametags/:id", tags_detailHandler.DeleteTags_detail)

	//wishlist
	engine.POST("/create-wishlist", wishlistHandler.CreateWishlist)
	engine.GET("/list-wishlist", wishlistHandler.GetListWishlist)
	// engine.GET("/get-wishlist/:id", wishlistHandler.GetDetailwishlist)
	engine.GET("/get-wishlist-game/:id", wishlistHandler.GetGame)
	engine.PUT("/update-wishlist/:id", wishlistHandler.UpdateWishlist)
	engine.DELETE("/delete-wishlist/:credential_id/:game", wishlistHandler.DeleteWishlist)

	//searching
	engine.POST("/search-game", gameHandler.SearchGame)
	engine.POST("/search-shop", shopHandler.SearchShop)
	engine.POST("/search-voucher", voucherHandler.SearchVoucher)

	engine.POST("/upload-file", fileHandler.CreateFile)

	return engine
}
