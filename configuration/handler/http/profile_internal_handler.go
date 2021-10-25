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

// CreateProfileBodyRequest defines all body attributes needed to add Profile.
type CreateProfileBodyRequest struct {
	Credential_id uuid.UUID `json:"credential_id"`
	First_name    string    `json:"first_name"`
	Last_name     string    `json:"last_name"`
	Phone         string    `json:"phone"`
	Gender        string    `json:"gender"`
	Birthday      string    `json:"birthday"`
}

// ProfileRowResponse defines all attributes needed to fulfill for Profile row entity.
type ProfileRowResponse struct {
	Id            uuid.UUID `json:"id"`
	Credential_id uuid.UUID `json:"credential_id"`
	First_name    string    `json:"first_name"`
	Last_name     string    `json:"last_name"`
	Phone         string    `json:"phone"`
	Gender        string    `json:"gender"`
	Birthday      string    `json:"birthday"`
}

// ProfileResponse defines all attributes needed to fulfill for pic Profile entity.
type ProfileDetailResponse struct {
	Id            uuid.UUID `json:"id"`
	Credential_id uuid.UUID `json:"credential_id"`
	First_name    string    `json:"first_name"`
	Last_name     string    `json:"last_name"`
	Phone         string    `json:"phone"`
	Gender        string    `json:"gender"`
	Birthday      string    `json:"birthday"`
}

func buildProfileRowResponse(Profile *entity.Profile) ProfileRowResponse {
	form := ProfileRowResponse{
		Id:            Profile.Id,
		Credential_id: Profile.Credential_id,
		First_name:    Profile.First_name,
		Last_name:     Profile.Last_name,
		Phone:         Profile.Phone,
		Gender:        Profile.Gender,
		Birthday:      Profile.Birthday,
	}

	return form
}

func buildProfileDetailResponse(Profile *entity.Profile) ProfileDetailResponse {
	form := ProfileDetailResponse{
		Id:            Profile.Id,
		Credential_id: Profile.Credential_id,
		First_name:    Profile.First_name,
		Last_name:     Profile.Last_name,
		Phone:         Profile.Phone,
		Gender:        Profile.Gender,
		Birthday:      Profile.Birthday,
	}

	return form
}

// QueryParamsProfile defines all attributes for input query params
type QueryParamsProfile struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaProfile define attributes needed for Meta
type MetaProfile struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaProfile creates an instance of Meta response.
func NewMetaProfile(limit, offset int, total int64) *MetaProfile {
	return &MetaProfile{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// ProfileHandler handles HTTP request related to Profile flow.
type ProfileHandler struct {
	service service.ProfileUseCase
}

// NewProfileHandler creates an instance of ProfileHandler.
func NewProfileHandler(service service.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *ProfileHandler) CreateProfile(echoCtx echo.Context) error {
	var form CreateProfileBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	ProfileEntity := entity.NewProfile(
		uuid.Nil,
		form.Credential_id,
		form.First_name,
		form.Last_name,
		form.Phone,
		form.Gender,
		form.Birthday,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), ProfileEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", ProfileEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ProfileHandler) GetListProfile(echoCtx echo.Context) error {
	var form QueryParamsProfile
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Profile, err := handler.service.GetListProfile(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Profile)
	return echoCtx.JSON(res.Status, res)

}

func (handler *ProfileHandler) GetDetailProfile(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	Profile, err := handler.service.GetDetailProfile(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Profile)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ProfileHandler) UpdateProfile(echoCtx echo.Context) error {
	var form CreateProfileBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	idParam := echoCtx.Param("id")

	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	_, err = handler.service.GetDetailProfile(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	ProfileEntity := entity.NewProfile(
		id,
		form.Credential_id,
		form.First_name,
		form.Last_name,
		form.Phone,
		form.Gender,
		form.Birthday,
	)

	if err := handler.service.UpdateProfile(echoCtx.Request().Context(), ProfileEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ProfileHandler) DeleteProfile(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	err = handler.service.DeleteProfile(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
