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

// CreateVoucherBodyRequest defines all body attributes needed to add Voucher.
type CreateVoucherBodyRequest struct {
	Game_id     uuid.UUID `json:"game_id"`
	VoucherName string    `json:"nama_voucher"`
	Harga       int       `json:"harga"`
}

// VoucherRowResponse defines all attributes needed to fulfill for Voucher row entity.
type VoucherRowResponse struct {
	Id          uuid.UUID `json:"id"`
	Game_id     uuid.UUID `json:"game_id"`
	VoucherName string    `json:"nama_voucher"`
	Harga       int       `json:"harga"`
}

// VoucherResponse defines all attributes needed to fulfill for pic Voucher entity.
type VoucherDetailResponse struct {
	Id          uuid.UUID `json:"id"`
	Game_id     uuid.UUID `json:"game_id"`
	VoucherName string    `json:"nama_voucher"`
	Harga       int       `json:"harga"`
}

func buildVoucherRowResponse(Voucher *entity.Voucher) VoucherRowResponse {
	form := VoucherRowResponse{
		Id:          Voucher.Id,
		Game_id:     Voucher.Game_id,
		VoucherName: Voucher.VoucherName,
		Harga:       Voucher.Harga,
	}

	return form
}

func buildVoucherDetailResponse(Voucher *entity.Voucher) VoucherDetailResponse {
	form := VoucherDetailResponse{
		Id:          Voucher.Id,
		Game_id:     Voucher.Game_id,
		VoucherName: Voucher.VoucherName,
		Harga:       Voucher.Harga,
	}

	return form
}

// QueryParamsVoucher defines all attributes for input query params
type QueryParamsVoucher struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaVoucher define attributes needed for Meta
type MetaVoucher struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaVoucher creates an instance of Meta response.
func NewMetaVoucher(limit, offset int, total int64) *MetaVoucher {
	return &MetaVoucher{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// VoucherHandler handles HTTP request related to Voucher flow.
type VoucherHandler struct {
	service service.VoucherUseCase
}

// NewVoucherHandler creates an instance of VoucherHandler.
func NewVoucherHandler(service service.VoucherUseCase) *VoucherHandler {
	return &VoucherHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *VoucherHandler) CreateVoucher(echoCtx echo.Context) error {
	var form CreateVoucherBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	VoucherEntity := entity.NewVoucher(
		uuid.Nil,
		form.Game_id,
		form.VoucherName,
		form.Harga,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), VoucherEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", VoucherEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *VoucherHandler) GetListVoucher(echoCtx echo.Context) error {
	var form QueryParamsVoucher
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Voucher, err := handler.service.GetListVoucher(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Voucher)
	return echoCtx.JSON(res.Status, res)

}

func (handler *VoucherHandler) GetDetailVoucher(echoCtx echo.Context) error {
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

	Voucher, err := handler.service.GetDetailVoucher(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Voucher)
	return echoCtx.JSON(res.Status, res)
}

func (handler *VoucherHandler) UpdateVoucher(echoCtx echo.Context) error {
	var form CreateVoucherBodyRequest
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

	_, err = handler.service.GetDetailVoucher(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	VoucherEntity := entity.NewVoucher(
		id,
		form.Game_id,
		form.VoucherName,
		form.Harga,
	)

	if err := handler.service.UpdateVoucher(echoCtx.Request().Context(), VoucherEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *VoucherHandler) DeleteVoucher(echoCtx echo.Context) error {
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

	err = handler.service.DeleteVoucher(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
