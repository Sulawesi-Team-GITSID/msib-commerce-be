package http

import (
	"backend-service/entity"
	"backend-service/service"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

// CreateGameBodyRequest defines all body attributes needed to add Game.
type CreateGameBodyRequest struct {
	Shop_id  uuid.UUID `json:"shop_id"`
	NamaGame string    `json:"nama_game"`
	Harga    int       `json:"harga"`
	Genre_id uuid.UUID `json:"genre_id"`
}

// GameRowResponse defines all attributes needed to fulfill for Game row entity.
type GameRowResponse struct {
	Id       uuid.UUID `json:"id"`
	Shop_id  uuid.UUID `json:"shop_id"`
	NamaGame string    `json:"nama_game"`
	Harga    int       `json:"harga"`
	Genre_id uuid.UUID `json:"genre_id"`
}

// GameResponse defines all attributes needed to fulfill for pic Game entity.
type GameDetailResponse struct {
	Id       uuid.UUID `json:"id"`
	Shop_id  uuid.UUID `json:"shop_id"`
	NamaGame string    `json:"nama_game"`
	Harga    int       `json:"harga"`
	Genre_id uuid.UUID `json:"genre_id"`
}

// Defines request form for search function
type SearchBodyRequest struct {
	Search string `json:"search" binding:"required"`
}

func buildGameRowResponse(Game *entity.Game) GameRowResponse {
	form := GameRowResponse{
		Id:       Game.Id,
		Shop_id:  Game.Shop_id,
		NamaGame: Game.NamaGame,
		Harga:    Game.Harga,
		Genre_id: Game.Genre_id,
	}

	return form
}

func buildGameDetailResponse(Game *entity.Game) GameDetailResponse {
	form := GameDetailResponse{
		Id:       Game.Id,
		Shop_id:  Game.Shop_id,
		NamaGame: Game.NamaGame,
		Harga:    Game.Harga,
		Genre_id: Game.Genre_id,
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
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	GameEntity := entity.NewGame(
		uuid.Nil,
		form.Shop_id,
		form.NamaGame,
		form.Harga,
		form.Genre_id,
		false,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), GameEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", GameEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *GameHandler) GetListGame(echoCtx echo.Context) error {
	var form QueryParamsGame
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Game, err := handler.service.GetListGame(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Game)
	return echoCtx.JSON(res.Status, res)

}

func (handler *GameHandler) GetListGameShop(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, nil, entity.ErrInvalidNullParam)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Game, err := handler.service.GetListGameShop(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Game)
	return echoCtx.JSON(res.Status, res)

}

func (handler *GameHandler) GetListGenre(echoCtx echo.Context) error {
	var form QueryParamsGame
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Game, err := handler.service.GetListGenre(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Game)
	return echoCtx.JSON(res.Status, res)

}

func (handler *GameHandler) GetListTrendGame(echoCtx echo.Context) error {
	var form QueryParamsGame
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Game, err := handler.service.GetListTrendGame(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Game)
	return echoCtx.JSON(res.Status, res)

}

func (handler *GameHandler) GetDetailGame(echoCtx echo.Context) error {
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

	Game, err := handler.service.GetDetailGame(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Game)
	return echoCtx.JSON(res.Status, res)
}

func (handler *GameHandler) SearchGame(echoCtx echo.Context) error {
	var form SearchBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	searchResult, err := handler.service.SearchGame(echoCtx.Request().Context(), form.Search)
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

func (handler *GameHandler) UpdateGame(echoCtx echo.Context) error {
	var form CreateGameBodyRequest
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

	_, err = handler.service.GetDetailGame(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	GameEntity := entity.NewGame(
		id,
		form.Shop_id,
		form.NamaGame,
		form.Harga,
		form.Genre_id,
		false,
	)

	if err := handler.service.UpdateGame(echoCtx.Request().Context(), GameEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *GameHandler) DeleteGame(echoCtx echo.Context) error {
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

	err = handler.service.DeleteGame(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
