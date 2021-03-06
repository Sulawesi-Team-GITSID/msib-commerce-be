package main

import (
	"context"
	"fmt"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"

	"backend-service/configuration/config"
	"backend-service/configuration/handler/http"
	"backend-service/configuration/repository"
	"backend-service/service"
)

func main() {
	log.Info().Msg("backend-service starting")
	cfg, err := config.NewConfig(".env")
	config.CheckError(err)

	// tool.ErrorClient = setupErrorReporting(context.Background(), cfg)

	var db *gorm.DB = config.OpenDatabase(cfg)

	defer func() {
		if sqlDB, err := db.DB(); err != nil {
			log.Fatal().Err(err)
			panic(err)
		} else {
			_ = sqlDB.Close()
		}
	}()
	// LoginHandler := &http.Loginhandler{}
	CredentialHandler := buildCredentialHandler(db)
	ProfileHandler := buildProfileHandler(db)
	GameHandler := buildGameHandler(db)
	VoucherHandler := buildVoucherHandler(db)
	VerificationHandler := buildVerificationHandler(db)
	Middlewarehandler := &http.Middlewarehandler{}
	ReviewHandler := buildReviewHandler(db)
	SuperAdminHandler := buildSuperAdminHandler(db)
	ShopHandler := buildShopHandler(db)
	GenreHandler := buildGenreHandler(db)
	TagsHandler := buildTagsHandler(db)
	Tags_detailHandler := buildTags_detailHandler(db)
	fileHandler := buildFileHandler(db)
	WishlistHandler := buildWishlistHandler(db)
	// usersHandler := buildUsersHandler(db)
	engine := http.NewGinEngine(CredentialHandler, ProfileHandler, GameHandler, VoucherHandler, VerificationHandler, Middlewarehandler, ReviewHandler, SuperAdminHandler, ShopHandler, GenreHandler, TagsHandler, Tags_detailHandler, fileHandler, WishlistHandler, cfg.InternalConfig.Username, cfg.InternalConfig.Password)

	server := &nethttp.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: engine,
	}
	// setGinMode(cfg.Env)
	// authenticate()
	runServer(server)
	waitForShutdown(server)
}

func runServer(srv *nethttp.Server) {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != nethttp.ErrServerClosed {
			log.Fatal().Err(err)
		}
	}()
}

func waitForShutdown(server *nethttp.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("shutting down backend-service")

	// The context is used to inform the server it has 2 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("backend-service forced to shutdown")
	}

	log.Info().Msg("backend-service exiting")
}

// func openDatabase(config *config.Config) *gorm.DB {
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		config.Database.Host,
// 		config.Database.Port,
// 		config.Database.Username,
// 		config.Database.Password,
// 		config.Database.Name)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	checkError(err)
// 	return db
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func buildCredentialHandler(db *gorm.DB) *http.CredentialHandler {
	repo := repository.NewCredentialRepository(db)
	CredentialService := service.NewCredentialService(repo)
	return http.NewCredentialHandler(CredentialService)
}

func buildProfileHandler(db *gorm.DB) *http.ProfileHandler {
	repo := repository.NewProfileRepository(db)
	ProfileService := service.NewProfileService(repo)
	return http.NewProfileHandler(ProfileService)
}
func buildGameHandler(db *gorm.DB) *http.GameHandler {
	repo := repository.NewGameRepository(db)
	GameService := service.NewGameService(repo)
	return http.NewGameHandler(GameService)
}

func buildVoucherHandler(db *gorm.DB) *http.VoucherHandler {
	repo := repository.NewVoucherRepository(db)
	VoucherService := service.NewVoucherService(repo)
	return http.NewVoucherHandler(VoucherService)
}

func buildVerificationHandler(db *gorm.DB) *http.VerificationHandler {
	repo := repository.NewVerificationRepository(db)
	VerificationService := service.NewVerificationService(repo)
	return http.NewVerificationHandler(VerificationService)
}

func buildReviewHandler(db *gorm.DB) *http.ReviewHandler {
	repo := repository.NewReviewRepository(db)
	ReviewService := service.NewReviewService(repo)
	return http.NewReviewHandler(ReviewService)
}

func buildSuperAdminHandler(db *gorm.DB) *http.SuperAdminHandler {
	repo := repository.NewSuperAdminRepository(db)
	superAdminService := service.NewSuperAdminService(repo)
	return http.NewSuperAdminHandler(superAdminService)
}

func buildShopHandler(db *gorm.DB) *http.ShopHandler {
	repo := repository.NewShopRepository(db)
	ShopService := service.NewShopService(repo)
	return http.NewShopHandler(ShopService)
}

func buildGenreHandler(db *gorm.DB) *http.GenreHandler {
	repo := repository.NewGenreRepository(db)
	GenreService := service.NewGenreService(repo)
	return http.NewGenreHandler(GenreService)
}

func buildTagsHandler(db *gorm.DB) *http.TagsHandler {
	repo := repository.NewTagsRepository(db)
	TagsService := service.NewTagsService(repo)
	return http.NewTagsHandler(TagsService)
}

func buildTags_detailHandler(db *gorm.DB) *http.Tags_detailHandler {
	repo := repository.NewTags_detailRepository(db)
	Tags_detailService := service.NewTags_detailService(repo)
	return http.NewTags_detailHandler(Tags_detailService)
}
func buildFileHandler(db *gorm.DB) *http.FileHandler {
	repo := repository.NewFileRepository(db)
	fileService := service.NewFileService(repo)
	return http.NewFileHandler(fileService)
}
func buildWishlistHandler(db *gorm.DB) *http.WishlistHandler {
	repo := repository.NewWishlistRepository(db)
	WishlistService := service.NewWishlistService(repo)
	return http.NewWishlistHandler(WishlistService)
}
