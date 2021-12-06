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

// CreateShopBodyRequest defines all body attributes needed to add Shop.
type CreateShopBodyRequest struct {
	Credential_id uuid.UUID `json:"Credential_id"`
	Name          string    `json:"name"`
	Location      string    `json:"location"`
}

// ShopRowResponse defines all attributes needed to fulfill for Shop row entity.
type ShopRowResponse struct {
	Id            uuid.UUID `json:"id"`
	Credential_id uuid.UUID `json:"Credential_id"`
	Name          string    `json:"name"`
	Location      string    `json:"location"`
}

// ShopResponse defines all attributes needed to fulfill for pic Shop entity.
type ShopDetailResponse struct {
	Id            uuid.UUID `json:"id"`
	Credential_id uuid.UUID `json:"Credential_id"`
	Name          string    `json:"name"`
	Location      string    `json:"location"`
}

func buildShopRowResponse(Shop *entity.Shop) ShopRowResponse {
	form := ShopRowResponse{
		Id:            Shop.Id,
		Credential_id: Shop.Credential_id,
		Name:          Shop.Name,
		Location:      Shop.Location,
	}

	return form
}

func buildShopDetailResponse(Shop *entity.Shop) ShopDetailResponse {
	form := ShopDetailResponse{
		Id:            Shop.Id,
		Credential_id: Shop.Credential_id,
		Name:          Shop.Name,
		Location:      Shop.Location,
	}

	return form
}

// QueryParamsShop defines all attributes for input query params
type QueryParamsShop struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaShop define attributes needed for Meta
type MetaShop struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaShop creates an instance of Meta response.
func NewMetaShop(limit, offset int, total int64) *MetaShop {
	return &MetaShop{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// ShopHandler handles HTTP request related to Shop flow.
type ShopHandler struct {
	service service.ShopUseCase
}

// NewShopHandler creates an instance of ShopHandler.
func NewShopHandler(service service.ShopUseCase) *ShopHandler {
	return &ShopHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *ShopHandler) CreateShop(echoCtx echo.Context) error {
	var form CreateShopBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	ShopEntity := entity.NewShop(
		uuid.Nil,
		form.Credential_id,
		"",
		form.Name,
		form.Location,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), ShopEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", ShopEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ShopHandler) GetListShop(echoCtx echo.Context) error {
	var form QueryParamsShop
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Shop, err := handler.service.GetListShop(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Shop)
	return echoCtx.JSON(res.Status, res)

}

func (handler *ShopHandler) GetDetailShop(echoCtx echo.Context) error {
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

	Shop, err := handler.service.GetDetailShop(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Shop)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ShopHandler) SearchShop(echoCtx echo.Context) error {
	var form SearchBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	searchResult, err := handler.service.SearchShop(echoCtx.Request().Context(), form.Search)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidCredential)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", searchResult)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ShopHandler) UpdateShop(echoCtx echo.Context) error {
	var form CreateShopBodyRequest
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

	_, err = handler.service.GetDetailShop(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	ShopEntity := entity.NewShop(
		id,
		form.Credential_id,
		"",
		form.Name,
		form.Location,
	)

	if err := handler.service.UpdateShop(echoCtx.Request().Context(), ShopEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ShopHandler) DeleteShop(echoCtx echo.Context) error {
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

	err = handler.service.DeleteShop(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
