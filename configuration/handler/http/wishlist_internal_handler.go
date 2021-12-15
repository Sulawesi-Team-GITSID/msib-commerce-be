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

// CreateWishlistBodyRequest defines all body attributes needed to add Wishlist.
type CreateWishlistBodyRequest struct {
	Credential_id uuid.UUID `json:"credential_id"`
	Game_id       uuid.UUID `json:"game_id"`
}

// WishlistRowResponse defines all attributes needed to fulfill for Wishlist row entity.
type WishlistRowResponse struct {
	Credential_id uuid.UUID `json:"credential_id"`
	Game_id       uuid.UUID `json:"game_id"`
}

// WishlistResponse defines all attributes needed to fulfill for pic Wishlist entity.
type WishlistDetailResponse struct {
	Credential_id uuid.UUID `json:"credential_id"`
	Game_id       uuid.UUID `json:"game_id"`
}

func buildWishlistRowResponse(Wishlist *entity.Wishlist) WishlistRowResponse {
	form := WishlistRowResponse{
		Credential_id: Wishlist.Credential_id,
		Game_id:       Wishlist.Game_id,
	}

	return form
}

func buildWishlistDetailResponse(Wishlist *entity.Wishlist) WishlistDetailResponse {
	form := WishlistDetailResponse{
		Credential_id: Wishlist.Credential_id,
		Game_id:       Wishlist.Game_id,
	}

	return form
}

// QueryParamsWishlist defines all attributes for input query params
type QueryParamsWishlist struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaWishlist define attributes needed for Meta
type MetaWishlist struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaWishlist creates an instance of Meta response.
func NewMetaWishlist(limit, offset int, total int64) *MetaWishlist {
	return &MetaWishlist{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// WishlistHandler handles HTTP request related to Wishlist flow.
type WishlistHandler struct {
	service service.WishlistUseCase
}

// NewWishlistHandler creates an instance of WishlistHandler.
func NewWishlistHandler(service service.WishlistUseCase) *WishlistHandler {
	return &WishlistHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *WishlistHandler) CreateWishlist(echoCtx echo.Context) error {
	var form CreateWishlistBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	WishlistEntity := entity.NewWishlist(
		form.Credential_id,
		form.Game_id,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), WishlistEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", WishlistEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *WishlistHandler) GetListWishlist(echoCtx echo.Context) error {
	var form QueryParamsWishlist
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Wishlist, err := handler.service.GetListWishlist(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Wishlist)
	return echoCtx.JSON(res.Status, res)

}

func (handler *WishlistHandler) GetGame(echoCtx echo.Context) error {
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

	Wishlist, err := handler.service.GetGame(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Wishlist)
	return echoCtx.JSON(res.Status, res)

}

func (handler *WishlistHandler) GetDetailWishlist(echoCtx echo.Context) error {
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

	Wishlist, err := handler.service.GetDetailWishlist(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Wishlist)
	return echoCtx.JSON(res.Status, res)
}

func (handler *WishlistHandler) UpdateWishlist(echoCtx echo.Context) error {
	var form CreateWishlistBodyRequest
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

	_, err = handler.service.GetDetailWishlist(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	WishlistEntity := entity.NewWishlist(
		form.Credential_id,
		form.Game_id,
	)

	if err := handler.service.UpdateWishlist(echoCtx.Request().Context(), WishlistEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *WishlistHandler) DeleteWishlist(echoCtx echo.Context) error {
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

	err = handler.service.DeleteWishlist(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
