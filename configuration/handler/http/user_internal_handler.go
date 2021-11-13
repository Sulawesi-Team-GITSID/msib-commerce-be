package http

import (
	"backend-service/entity"
	"backend-service/service"
	"log"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type RegisterBodyRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// type LoginBodyRequest struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

type UsersRowResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type UsersDetailResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type QueryParamsUsers struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// type JWTCustomClaims struct {
// 	Nama  string    `json:"nama"`
// 	Email string    `json:"email"`
// 	Id    uuid.UUID `json:"id"`
// 	jwt.StandardClaims
// }

// type Result struct {
// 	Token string    `json:"token"`
// 	Id    uuid.UUID `json:"id"`
// }

type MetaUsers struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// UsersHandler handles HTTP request related to user flow.
type UsersHandler struct {
	service service.UsersUseCase
}

// NewMetaUsers creates an instance of Meta response.
func NewMetaUsers(limit, offset int, total int64) *MetaUsers {
	return &MetaUsers{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// NewUsersHandler creates an instance of UsersHandler.
func NewUsersHandler(service service.UsersUseCase) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

// Create handles users creation.
// It will reject the request if the request doesn't have required data,
// func (handler *UsersHandler) LoginUser(echoCtx echo.Context) error {
// 	var form LoginBodyRequest
// 	if err := echoCtx.Bind(&form); err != nil {
// 		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
// 		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

// 	}

// 	userData, err := handler.service.Login(echoCtx.Request().Context(), form.Email, form.Password)
// 	if err != nil {
// 		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidCredential)
// 		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
// 	}

// 	claims := &JWTCustomClaims{
// 		userData.Nama,
// 		userData.Email,
// 		userData.ID,
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

// 	if err != nil {
// 		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
// 		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
// 	}

// 	result := &Result{
// 		tokenString,
// 		userData.ID,
// 	}

// 	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", result)
// 	return echoCtx.JSON(res.Status, res)
// }

// Create handles users creation.
// It will reject the request if the request doesn't have required data,
func (handler *UsersHandler) Register(echoCtx echo.Context) error {
	var form RegisterBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	usersEntity := entity.NewUsers(
		uuid.Nil,
		form.Name,
		form.Email,
		form.Password,
	)

	if err := handler.service.Register(echoCtx.Request().Context(), usersEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", true)
	return echoCtx.JSON(res.Status, res)
}

func (handler *UsersHandler) GetProfile(echoCtx echo.Context) error {

	// get user from JWT
	userClaim := echoCtx.Request().Header
	log.Print("Test : ", userClaim)

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", userClaim)
	return echoCtx.JSON(res.Status, res)
}
