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

// CreateTagsBodyRequest defines all body attributes needed to add Tags.
type CreateTagsBodyRequest struct {
	Name string `json:"name"`
}

// TagsRowResponse defines all attributes needed to fulfill for Tags row entity.
type TagsRowResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// TagsResponse defines all attributes needed to fulfill for pic Tags entity.
type TagsDetailResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func buildTagsRowResponse(Tags *entity.Tags) TagsRowResponse {
	form := TagsRowResponse{
		Id:   Tags.Id,
		Name: Tags.Name,
	}

	return form
}

func buildTagsDetailResponse(Tags *entity.Tags) TagsDetailResponse {
	form := TagsDetailResponse{
		Id:   Tags.Id,
		Name: Tags.Name,
	}

	return form
}

// QueryParamsTags defines all attributes for input query params
type QueryParamsTags struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaTags define attributes needed for Meta
type MetaTags struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaTags creates an instance of Meta response.
func NewMetaTags(limit, offset int, total int64) *MetaTags {
	return &MetaTags{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// TagsHandler handles HTTP request related to Tags flow.
type TagsHandler struct {
	service service.TagsUseCase
}

// NewTagsHandler creates an instance of TagsHandler.
func NewTagsHandler(service service.TagsUseCase) *TagsHandler {
	return &TagsHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *TagsHandler) CreateTags(echoCtx echo.Context) error {
	var form CreateTagsBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	TagsEntity := entity.NewTags(
		uuid.Nil,
		form.Name,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), TagsEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", TagsEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *TagsHandler) GetListTags(echoCtx echo.Context) error {
	var form QueryParamsTags
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Tags, err := handler.service.GetListTags(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Tags)
	return echoCtx.JSON(res.Status, res)

}

func (handler *TagsHandler) GetDetailTags(echoCtx echo.Context) error {
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

	Tags, err := handler.service.GetDetailTags(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Tags)
	return echoCtx.JSON(res.Status, res)
}

func (handler *TagsHandler) UpdateTags(echoCtx echo.Context) error {
	var form CreateTagsBodyRequest
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

	_, err = handler.service.GetDetailTags(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	TagsEntity := entity.NewTags(
		id,
		form.Name,
	)

	if err := handler.service.UpdateTags(echoCtx.Request().Context(), TagsEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *TagsHandler) DeleteTags(echoCtx echo.Context) error {
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

	err = handler.service.DeleteTags(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
