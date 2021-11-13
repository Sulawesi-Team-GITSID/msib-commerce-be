package http

import (
	"backend-service/entity"
	"backend-service/service"
	nethttp "net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

// CreateVerificationBodyRequest defines all body attributes needed to add Verification.
type CreateVerificationBodyRequest struct {
	Credential_id uuid.UUID `json:"credential_id"`
	Code          string    `json:"code"`
	Expiresat     time.Time `json:"expiresat"`
}

// VerificationRowResponse defines all attributes needed to fulfill for Verification row entity.
type VerificationRowResponse struct {
	Id            uuid.UUID `json:"id"`
	Credential_id uuid.UUID `json:"credential_id"`
	Code          string    `json:"code"`
	Expiresat     time.Time `json:"expiresat"`
}

// VerificationResponse defines all attributes needed to fulfill for pic Verification entity.
type VerificationDetailResponse struct {
	Id            uuid.UUID `json:"id"`
	Credential_id uuid.UUID `json:"credential_id"`
	Code          string    `json:"code"`
	Expiresat     time.Time `json:"expiresat"`
}

type VerifyBodyRequest struct {
	Credential_id string    `json:"credential_id" binding:"required"`
	Code          string    `json:"code" binding:"required"`
	Expiresat     time.Time `json:"expiresat" binding:"required"`
}

type VerifyResult struct {
	Credential_id uuid.UUID `json:"credential_id"`
	Code          string    `json:"code"`
	Expiresat     time.Time `json:"expiresat"`
	Email         string    `json:"email"`
}

func buildVerificationRowResponse(Verification *entity.Verification) VerificationRowResponse {
	form := VerificationRowResponse{
		Id:            Verification.Id,
		Credential_id: Verification.Credential_id,
		Code:          Verification.Code,
		Expiresat:     Verification.Expiresat,
	}

	return form
}

func buildVerificationDetailResponse(Verification *entity.Verification) VerificationDetailResponse {
	form := VerificationDetailResponse{
		Id:            Verification.Id,
		Credential_id: Verification.Credential_id,
		Code:          Verification.Code,
		Expiresat:     Verification.Expiresat,
	}

	return form
}

// QueryParamsVerification defines all attributes for input query params
type QueryParamsVerification struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaVerification define attributes needed for Meta
type MetaVerification struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaVerification creates an instance of Meta response.
func NewMetaVerification(limit, offset int, total int64) *MetaVerification {
	return &MetaVerification{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// VerificationHandler handles HTTP request related to Verification flow.
type VerificationHandler struct {
	service service.VerificationUseCase
}

// NewVerificationHandler creates an instance of VerificationHandler.
func NewVerificationHandler(service service.VerificationUseCase) *VerificationHandler {
	return &VerificationHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *VerificationHandler) CreateVerification(echoCtx echo.Context) error {
	var form CreateVerificationBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	VerificationEntity := entity.NewVerification(
		uuid.Nil,
		form.Credential_id,
		form.Code,
		form.Expiresat,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), VerificationEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", VerificationEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *VerificationHandler) GetListVerification(echoCtx echo.Context) error {
	var form QueryParamsVerification
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Verification, err := handler.service.GetListVerification(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Verification)
	return echoCtx.JSON(res.Status, res)

}

func (handler *VerificationHandler) GetDetailVerification(echoCtx echo.Context) error {
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

	Verification, err := handler.service.GetDetailVerification(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Verification)
	return echoCtx.JSON(res.Status, res)
}

func (handler *VerificationHandler) UpdateVerification(echoCtx echo.Context) error {
	var form CreateVerificationBodyRequest
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

	_, err = handler.service.GetDetailVerification(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	VerificationEntity := entity.NewVerification(
		id,
		form.Credential_id,
		form.Code,
		form.Expiresat,
	)

	if err := handler.service.UpdateVerification(echoCtx.Request().Context(), VerificationEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *VerificationHandler) DeleteVerification(echoCtx echo.Context) error {
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

	err = handler.service.DeleteVerification(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *VerificationHandler) Verify(echoCtx echo.Context) error {
	var form VerifyBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	userData, err := handler.service.Verify(echoCtx.Request().Context(), form.Credential_id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrAccessDenied)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	result := &VerifyResult{
		userData.Credential_id,
		userData.Code,
		userData.Expiresat,
		userData.Email,
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", result)
	verify_mail(*result)
	return echoCtx.JSON(res.Status, res)
}
