package http

import (
	"backend-service/entity"
	"backend-service/service"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

// CreateCredentialBodyRequest defines all body attributes needed to add Credential.
type CreateCredentialBodyRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Seller   bool   `json:"seller"`
	Verified bool   `json:"verified"`
}

// CredentialRowResponse defines all attributes needed to fulfill for Credential row entity.
type CredentialRowResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Seller   bool      `json:"seller"`
	Verified bool      `json:"verified"`
}

// CredentialResponse defines all attributes needed to fulfill for pic Credential entity.
type CredentialDetailResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Seller   bool      `json:"seller"`
	Verified bool      `json:"verified"`
}

func buildCredentialRowResponse(Credential *entity.Credential) CredentialRowResponse {
	form := CredentialRowResponse{
		Id:       Credential.Id,
		Username: Credential.Username,
		Password: Credential.Password,
		Seller:   Credential.Seller,
		Verified: Credential.Verified,
	}

	return form
}

func buildCredentialDetailResponse(Credential *entity.Credential) CredentialDetailResponse {
	form := CredentialDetailResponse{
		Id:       Credential.Id,
		Username: Credential.Username,
		Password: Credential.Password,
		Seller:   Credential.Seller,
		Verified: Credential.Verified,
	}

	return form
}

// QueryParamsCredential defines all attributes for input query params
type QueryParamsCredential struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaCredential define attributes needed for Meta
type MetaCredential struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaCredential creates an instance of Meta response.
func NewMetaCredential(limit, offset int, total int64) *MetaCredential {
	return &MetaCredential{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// CredentialHandler handles HTTP request related to Credential flow.
type CredentialHandler struct {
	service service.CredentialUseCase
}

// NewCredentialHandler creates an instance of CredentialHandler.
func NewCredentialHandler(service service.CredentialUseCase) *CredentialHandler {
	return &CredentialHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *CredentialHandler) CreateCredential(echoCtx echo.Context) error {
	var form CreateCredentialBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	CredentialEntity := entity.NewCredential(
		uuid.Nil,
		form.Username,
		form.Password,
		form.Seller,
		form.Verified,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), CredentialEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", CredentialEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *CredentialHandler) GetListCredential(echoCtx echo.Context) error {
	var form QueryParamsCredential
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	Credential, err := handler.service.GetListCredential(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", Credential)
	return echoCtx.JSON(res.Status, res)

}

// func GetCredentialdata() []CreateCredentialBodyRequest {
// 	db := config.CreateConnection()

// 	// kita tutup koneksinya di akhir proses
// 	defer db.Close()

// 	var Credentials []CreateCredentialBodyRequest

// 	// kita buat select query
// 	sqlStatement := `SELECT Username, password, Seller FROM public."Credential"`

// 	// mengeksekusi sql query
// 	rows, err := db.Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Query could not be executed. %v", err)
// 	}

// 	// kita tutup eksekusi proses sql qeurynya
// 	defer rows.Close()

// 	// kita iterasi mengambil datanya
// 	for rows.Next() {
// 		var Credential CreateCredentialBodyRequest

// 		// kita ambil datanya dan unmarshal ke structnya
// 		err = rows.Scan(&Credential.Username, &Credential.Password, &Credential.Seller)

// 		if err != nil {
// 			log.Fatalf("No data. %v", err)
// 		}

// 		// masukkan kedalam slice bukus
// 		Credentials = append(Credentials, Credential)

// 	}

// 	// return empty buku atau jika error
// 	return Credentials
// }
