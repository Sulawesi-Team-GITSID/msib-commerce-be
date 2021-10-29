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

// CreateGameBodyRequest defines all body attributes needed to add Game.
type CreateGameBodyRequest struct {
	NamaGame string `json:"nama_game"`
	Harga    int    `json:"harga"`
	Genre    string `json:"genre"`
}

// GameRowResponse defines all attributes needed to fulfill for Game row entity.
type GameRowResponse struct {
	Id       uuid.UUID `json:"id"`
	NamaGame string    `json:"nama_game"`
	Harga    int       `json:"harga"`
	Genre    string    `json:"genre"`
}

// GameResponse defines all attributes needed to fulfill for pic Game entity.
type GameDetailResponse struct {
	Id       uuid.UUID `json:"id"`
	NamaGame string    `json:"nama_game"`
	Harga    int       `json:"harga"`
	Genre    string    `json:"genre"`
}

func buildGameRowResponse(Game *entity.Game) GameRowResponse {
	form := GameRowResponse{
		Id:       Game.Id,
		NamaGame: Game.NamaGame,
		Harga:    Game.Harga,
		Genre:    Game.Genre,
	}

	return form
}

func buildGameDetailResponse(Game *entity.Game) GameDetailResponse {
	form := GameDetailResponse{
		Id:       Game.Id,
		NamaGame: Game.NamaGame,
		Harga:    Game.Harga,
		Genre:    Game.Genre,
	}

	return form
}

// QueryParamsGame defines all attributes for input query params
type QueryParamsGame struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaGame define attributes needed for Meta
type MetaGame struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaGame creates an instance of Meta response.
func NewMetaGame(limit, offset int, total int64) *MetaGame {
	return &MetaGame{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// GameHandler handles HTTP request related to Game flow.
type GameHandler struct {
	service service.GameUseCase
}

// NewGameHandler creates an instance of GameHandler.
func NewGameHandler(service service.GameUseCase) *GameHandler {
	return &GameHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *GameHandler) CreateGame(echoCtx echo.Context) error {
	var form CreateGameBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	GameEntity := entity.NewGame(
		uuid.Nil,
		form.NamaGame,
		form.Harga,
		form.Genre,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), GameEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", GameEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *GameHandler) GetListGame(echoCtx echo.Context) error {
	var form QueryParamsGame
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Game, err := handler.service.GetListGame(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Game)
	return echoCtx.JSON(res.Status, res)

}

func (handler *GameHandler) GetDetailGame(echoCtx echo.Context) error {
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

	Game, err := handler.service.GetDetailGame(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Game)
	return echoCtx.JSON(res.Status, res)
}

func (handler *GameHandler) UpdateGame(echoCtx echo.Context) error {
	var form CreateGameBodyRequest
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

	_, err = handler.service.GetDetailGame(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	GameEntity := entity.NewGame(
		id,
		form.NamaGame,
		form.Harga,
		form.Genre,
	)

	if err := handler.service.UpdateGame(echoCtx.Request().Context(), GameEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *GameHandler) DeleteGame(echoCtx echo.Context) error {
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

	err = handler.service.DeleteGame(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}