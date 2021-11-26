package http

import (
	"backend-service/entity"
	"backend-service/service"
	nethttp "net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type RegisterBodyRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginAdminBodyRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SuperAdminRowResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type SuperAdminDetailResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type QueryParamsSuperAdmin struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

type JWTAdminClaims struct {
	Nama  string    `json:"nama"`
	Email string    `json:"email"`
	Id    uuid.UUID `json:"id"`
	jwt.StandardClaims
}

type ResultAdmin struct {
	Token string    `json:"token"`
	Id    uuid.UUID `json:"id"`
}

type MetaSuperAdmin struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// SuperAdminHandler handles HTTP request related to user flow.
type SuperAdminHandler struct {
	service service.SuperAdminUseCase
}

// NewMetaSuperAdmin creates an instance of Meta response.
func NewMetaSuperAdmin(limit, offset int, total int64) *MetaSuperAdmin {
	return &MetaSuperAdmin{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// NewSuperAdminHandler creates an instance of SuperAdminHandler.
func NewSuperAdminHandler(service service.SuperAdminUseCase) *SuperAdminHandler {
	return &SuperAdminHandler{
		service: service,
	}
}

// Create handles SuperAdmin creation.
// It will reject the request if the request doesn't have required data,
func (handler *SuperAdminHandler) LoginAdmin(echoCtx echo.Context) error {
	var form LoginBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	userData, err := handler.service.LoginAdmin(echoCtx.Request().Context(), form.Email, form.Password)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidCredential)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	claims := &JWTAdminClaims{
		userData.Nama,
		userData.Email,
		userData.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	result := &ResultAdmin{
		tokenString,
		userData.ID,
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", result)
	return echoCtx.JSON(res.Status, res)
}

// Create handles SuperAdmin creation.
// It will reject the request if the request doesn't have required data,
func (handler *SuperAdminHandler) Register(echoCtx echo.Context) error {
	var form RegisterBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	SuperAdminEntity := entity.NewSuperAdmin(
		uuid.Nil,
		form.Name,
		form.Email,
		form.Password,
	)

	if err := handler.service.Register(echoCtx.Request().Context(), SuperAdminEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", true)
	return echoCtx.JSON(res.Status, res)
}

// func (handler *SuperAdminHandler) GetProfile(echoCtx echo.Context) error {

// 	// get user from JWT
// 	userClaim := echoCtx.Request().Header
// 	log.Print("Test : ", userClaim)

// 	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", userClaim)
// 	return echoCtx.JSON(res.Status, res)
// }
