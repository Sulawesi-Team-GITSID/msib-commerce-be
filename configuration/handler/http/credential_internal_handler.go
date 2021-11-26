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
	_ "github.com/lib/pq"
)

// CreateCredentialBodyRequest defines all body attributes needed to add Credential.
type CreateCredentialBodyRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Seller   bool   `json:"seller"`
	Verified bool   `json:"verified"`
}

type UpdatePassword struct {
	Password string `json:"password"`
}

// CredentialRowResponse defines all attributes needed to fulfill for Credential row entity.
type CredentialRowResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Seller   bool      `json:"seller"`
	Verified bool      `json:"verified"`
}

// CredentialResponse defines all attributes needed to fulfill for pic Credential entity.
type CredentialDetailResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Seller   bool      `json:"seller"`
	Verified bool      `json:"verified"`
}

// Defines request form for login function
type LoginBodyRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SearchEmailRequest struct {
	Id    uuid.UUID `json:"id" binding:"required"`
	Email string    `json:"email" binding:"required"`
}

type SearchEmailResult struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type JWTCustomClaims struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Id       uuid.UUID `json:"id"`
	Seller   bool      `json:"seller"`
	jwt.StandardClaims
}

type Result struct {
	Token string    `json:"token"`
	Id    uuid.UUID `json:"id"`
}

func buildCredentialRowResponse(Credential *entity.Credential) CredentialRowResponse {
	form := CredentialRowResponse{
		Id:       Credential.Id,
		Username: Credential.Username,
		Email:    Credential.Email,
		Password: Credential.Password,
		Seller:   Credential.Seller,
		Verified: Credential.Verified,
	}

	return form
}

func buildCredentialDetailResponse(Credential *entity.Credential) CredentialDetailResponse {
	form := CredentialDetailResponse{
		Id:       Credential.Id,
		Username: Credential.Username,
		Email:    Credential.Email,
		Password: Credential.Password,
		Seller:   Credential.Seller,
		Verified: Credential.Verified,
	}

	return form
}

// QueryParamsCredential defines all attributes for input query params
type QueryParamsCredential struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaCredential define attributes needed for Meta
type MetaCredential struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaCredential creates an instance of Meta response.
func NewMetaCredential(limit, offset int, total int64) *MetaCredential {
	return &MetaCredential{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// CredentialHandler handles HTTP request related to Credential flow.
type CredentialHandler struct {
	service service.CredentialUseCase
}

// NewCredentialHandler creates an instance of CredentialHandler.
func NewCredentialHandler(service service.CredentialUseCase) *CredentialHandler {
	return &CredentialHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *CredentialHandler) CreateCredential(echoCtx echo.Context) error {
	var form CreateCredentialBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}
	uuidsaver := uuid.New()
	CredentialEntity := entity.NewCredential(
		uuidsaver,
		form.Username,
		form.Email,
		form.Password,
		form.Seller,
		form.Verified,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), CredentialEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", CredentialEntity)
	Sendmail(form)
	return echoCtx.JSON(res.Status, res)
}

func (handler *CredentialHandler) GetListCredential(echoCtx echo.Context) error {
	var form QueryParamsCredential
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Credential, err := handler.service.GetListCredential(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Credential)
	return echoCtx.JSON(res.Status, res)

}

func (handler *CredentialHandler) Login(echoCtx echo.Context) error {
	var form LoginBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	userData, err := handler.service.Login(echoCtx.Request().Context(), form.Email, form.Password)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidCredential)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	claims := &JWTCustomClaims{
		userData.Username,
		userData.Email,
		userData.Id,
		userData.Seller,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = userData.Username
	// claims["email"] = userData.Email
	// claims["id"] = userData.Id
	// claims["seller"] = userData.Seller
	// claims["exp"] = time.Now().Add(time.Minute * 45).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	result := &Result{
		tokenString,
		userData.Id,
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", result)
	return echoCtx.JSON(res.Status, res)
}

func (handler *CredentialHandler) UpdateCredentialVerify(echoCtx echo.Context) error {
	var form CreateCredentialBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}
	idParam := echoCtx.Param("id")

	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, nil, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	_, err = handler.service.GetDetailCredential(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	CredentialEntity := entity.NewCredential(
		id,
		form.Username,
		form.Email,
		form.Password,
		form.Seller,
		form.Verified,
	)

	if err := handler.service.UpdateCredentialVerify(echoCtx.Request().Context(), CredentialEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
func (handler *CredentialHandler) EmailSearch(echoCtx echo.Context) error {
	var form SearchEmailRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	emailResult, err := handler.service.EmailSearch(echoCtx.Request().Context(), form.Email)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidCredential)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}
	result := &SearchEmailResult{
		emailResult.Id,
		emailResult.Email,
	}
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", result)
	forgot_mail(emailResult.Id, emailResult.Email)
	return echoCtx.JSON(res.Status, res)
}
func (handler *CredentialHandler) ForgotPassword(echoCtx echo.Context) error {
	var form CreateCredentialBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}
	idParam := echoCtx.Param("id")

	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, nil, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	_, err = handler.service.GetDetailCredential(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	CredentialEntity := entity.NewCredential(
		id,
		form.Username,
		form.Email,
		form.Password,
		form.Seller,
		form.Verified,
	)

	if err := handler.service.ForgotPassword(echoCtx.Request().Context(), CredentialEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", form)
	return echoCtx.JSON(res.Status, res)
}
