package http

import (
	"backend-service/entity"
	"backend-service/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

// CreateTags_detailBodyRequest defines all body attributes needed to add Tags_detail.
type CreateTags_detailBodyRequest struct {
	Game_id uuid.UUID `json:"game_id"`
	Tags_id uuid.UUID `json:"tags_id"`
}

// Tags_detailRowResponse defines all attributes needed to fulfill for Tags_detail row entity.
type Tags_detailRowResponse struct {
	Game_id uuid.UUID `json:"game_id"`
	Tags_id uuid.UUID `json:"tags_id"`
}

// Tags_detailResponse defines all attributes needed to fulfill for pic Tags_detail entity.
type Tags_detailDetailResponse struct {
	Game_id uuid.UUID `json:"game_id"`
	Tags_id uuid.UUID `json:"tags_id"`
}

func buildTags_detailRowResponse(Tags_detail *entity.Tags_detail) Tags_detailRowResponse {
	form := Tags_detailRowResponse{
		Game_id: Tags_detail.Game_id,
		Tags_id: Tags_detail.Tags_id,
	}

	return form
}

func buildTags_detailDetailResponse(Tags_detail *entity.Tags_detail) Tags_detailDetailResponse {
	form := Tags_detailDetailResponse{
		Game_id: Tags_detail.Game_id,
		Tags_id: Tags_detail.Tags_id,
	}

	return form
}

// QueryParamsTags_detail defines all attributes for input query params
type QueryParamsTags_detail struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaTags_detail define attributes needed for Meta
type MetaTags_detail struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaTags_detail creates an instance of Meta response.
func NewMetaTags_detail(limit, offset int, total int64) *MetaTags_detail {
	return &MetaTags_detail{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// Tags_detailHandler handles HTTP request related to Tags_detail flow.
type Tags_detailHandler struct {
	service service.Tags_detailUseCase
}

// NewTags_detailHandler creates an instance of Tags_detailHandler.
func NewTags_detailHandler(service service.Tags_detailUseCase) *Tags_detailHandler {
	return &Tags_detailHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *Tags_detailHandler) CreateTags_detail(echoCtx echo.Context) error {
	var form CreateTags_detailBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	Tags_detailEntity := entity.NewTags_detail(
		form.Game_id,
		form.Tags_id,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), Tags_detailEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", Tags_detailEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *Tags_detailHandler) GetListTags_detail(echoCtx echo.Context) error {
	var form QueryParamsTags_detail
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Tags_detail, err := handler.service.GetListTags_detail(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Tags_detail)
	return echoCtx.JSON(res.Status, res)

}

func (handler *Tags_detailHandler) GetGameTags(echoCtx echo.Context) error {
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

	Tags_detail, err := handler.service.GetGameTags(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Tags_detail)
	return echoCtx.JSON(res.Status, res)

}

func (handler *Tags_detailHandler) GetDetailTags_detail(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	Tags_detail, err := handler.service.GetDetailTags_detail(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Tags_detail)
	return echoCtx.JSON(res.Status, res)
}

func (handler *Tags_detailHandler) UpdateTags_detail(echoCtx echo.Context) error {
	var form CreateTags_detailBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	idParam := echoCtx.Param("id")

	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	_, err = handler.service.GetDetailTags_detail(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	Tags_detailEntity := entity.NewTags_detail(
		form.Game_id,
		form.Tags_id,
	)

	if err := handler.service.UpdateTags_detail(echoCtx.Request().Context(), Tags_detailEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *Tags_detailHandler) DeleteTags_detail(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	err = handler.service.DeleteTags_detail(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
