package http

import (
	"backend-service/entity"
	"backend-service/service"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

// CreateReviewBodyRequest defines all body attributes needed to add Review.
type CreateReviewBodyRequest struct {
	Game_id uuid.UUID `json:"game_id"`
	Rating  float64   `json:"rating"`
	Comment string    `json:"comment"`
}

// ReviewRowResponse defines all attributes needed to fulfill for Review row entity.
type ReviewRowResponse struct {
	Id      uuid.UUID `json:"id"`
	Game_id uuid.UUID `json:"game_id"`
	Rating  float64   `json:"rating"`
	Comment string    `json:"comment"`
}

// ReviewResponse defines all attributes needed to fulfill for pic Review entity.
type ReviewDetailResponse struct {
	Id      uuid.UUID `json:"id"`
	Game_id uuid.UUID `json:"game_id"`
	Rating  float64   `json:"rating"`
	Comment string    `json:"comment"`
}

func buildReviewRowResponse(Review *entity.Review) ReviewRowResponse {
	form := ReviewRowResponse{
		Id:      Review.Id,
		Game_id: Review.Game_id,
		Rating:  Review.Rating,
		Comment: Review.Comment,
	}

	return form
}

func buildReviewDetailResponse(Review *entity.Review) ReviewDetailResponse {
	form := ReviewDetailResponse{
		Id:      Review.Id,
		Game_id: Review.Game_id,
		Rating:  Review.Rating,
		Comment: Review.Comment,
	}

	return form
}

// QueryParamsReview defines all attributes for input query params
type QueryParamsReview struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaReview define attributes needed for Meta
type MetaReview struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaReview creates an instance of Meta response.
func NewMetaReview(limit, offset int, total int64) *MetaReview {
	return &MetaReview{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// ReviewHandler handles HTTP request related to Review flow.
type ReviewHandler struct {
	service service.ReviewUseCase
}

// NewReviewHandler creates an instance of ReviewHandler.
func NewReviewHandler(service service.ReviewUseCase) *ReviewHandler {
	return &ReviewHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *ReviewHandler) CreateReview(echoCtx echo.Context) error {
	var form CreateReviewBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	ReviewEntity := entity.NewReview(
		uuid.Nil,
		form.Game_id,
		form.Rating,
		form.Comment,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), ReviewEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", ReviewEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ReviewHandler) GetListReview(echoCtx echo.Context) error {
	var form QueryParamsReview
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Review, err := handler.service.GetListReview(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Review)
	return echoCtx.JSON(res.Status, res)

}
