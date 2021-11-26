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

// CreateGenreBodyRequest defines all body attributes needed to add Genre.
type CreateGenreBodyRequest struct {
	Name string `json:"name"`
}

// GenreRowResponse defines all attributes needed to fulfill for Genre row entity.
type GenreRowResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// GenreResponse defines all attributes needed to fulfill for pic Genre entity.
type GenreDetailResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func buildGenreRowResponse(Genre *entity.Genre) GenreRowResponse {
	form := GenreRowResponse{
		Id:   Genre.Id,
		Name: Genre.Name,
	}

	return form
}

func buildGenreDetailResponse(Genre *entity.Genre) GenreDetailResponse {
	form := GenreDetailResponse{
		Id:   Genre.Id,
		Name: Genre.Name,
	}

	return form
}

// QueryParamsGenre defines all attributes for input query params
type QueryParamsGenre struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaGenre define attributes needed for Meta
type MetaGenre struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaGenre creates an instance of Meta response.
func NewMetaGenre(limit, offset int, total int64) *MetaGenre {
	return &MetaGenre{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// GenreHandler handles HTTP request related to Genre flow.
type GenreHandler struct {
	service service.GenreUseCase
}

// NewGenreHandler creates an instance of GenreHandler.
func NewGenreHandler(service service.GenreUseCase) *GenreHandler {
	return &GenreHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *GenreHandler) CreateGenre(echoCtx echo.Context) error {
	var form CreateGenreBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	GenreEntity := entity.NewGenre(
		uuid.Nil,
		form.Name,
		false,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), GenreEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", GenreEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *GenreHandler) GetListGenre(echoCtx echo.Context) error {
	var form QueryParamsGenre
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Genre, err := handler.service.GetListGenre(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Genre)
	return echoCtx.JSON(res.Status, res)

}

func (handler *GenreHandler) GetDetailGenre(echoCtx echo.Context) error {
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

	Genre, err := handler.service.GetDetailGenre(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Genre)
	return echoCtx.JSON(res.Status, res)
}

func (handler *GenreHandler) UpdateGenre(echoCtx echo.Context) error {
	var form CreateGenreBodyRequest
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

	_, err = handler.service.GetDetailGenre(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	GenreEntity := entity.NewGenre(
		id,
		form.Name,
		false,
	)

	if err := handler.service.UpdateGenre(echoCtx.Request().Context(), GenreEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *GenreHandler) DeleteGenre(echoCtx echo.Context) error {
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

	err = handler.service.DeleteGenre(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
