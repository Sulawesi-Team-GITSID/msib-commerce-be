package http

import (
	"backend-service/configuration/config"
	"backend-service/entity"
	"backend-service/service"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	nethttp "net/http"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type UploadBodyRequest struct {
	File      []byte `form:"file" validName:"file"`
	Folder    string `form:"folder" validName:"folder"`
	Entity_id string `json:"entity_id" form:"entity_id" validName:"entity_id"`
}

type FileRowResponse struct {
	ID   uuid.UUID `json:"id"`
	File string    `json:"file" binding:"required"`
}

type FileDetailResponse struct {
	ID   uuid.UUID `json:"id"`
	File string    `json:"file" binding:"required"`
}

type FileHandler struct {
	service service.FileUseCase
}

// NewFileHandler creates an instance of FileHandler.
func NewFileHandler(service service.FileUseCase) *FileHandler {
	return &FileHandler{
		service: service,
	}
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (handler *FileHandler) CreateFile(echoCtx echo.Context) error {
	var base64Encoding string
	var form UploadBodyRequest

	_, err := echoCtx.FormFile("file")
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidFileFormat)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	if err := echoCtx.Bind(&form); err != nil {
		entitya, err := echoCtx.FormFile("entity_id")
		fmt.Printf("%+v\n", entitya)
		fmt.Printf("%+v\n", form)
		errorResponse := buildErrorResponse(nethttp.StatusBadRequest, err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	dir, err := os.Getwd()
	byteFile, filename, _ := GetByteFile(echoCtx)

	// kind, err := filetype.Match(byteFile)

	// log.Printf("Kind : ", kind)

	fileLocation := filepath.Join(dir, "files", *filename)
	newuuid := uuid.New()
	fileName := newuuid.String()

	cfg, _ := config.NewConfig(".env")
	cld, err := cloudinary.NewFromParams(cfg.Cloudinary.Name, cfg.Cloudinary.Key, cfg.Cloudinary.Secret)

	fileEntity := entity.NewFile(
		newuuid,
		uuid.MustParse(form.Entity_id),
		form.Folder+"/"+fileName,
	)

	// Determine the content type of the image file
	mimeType := http.DetectContentType(byteFile)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(byteFile)
	// fmt.Println(base64Encoding)

	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}
	fmt.Println(fileLocation)
	resp, err := cld.Upload.Upload(echoCtx.Request().Context(), base64Encoding, uploader.UploadParams{
		PublicID:     form.Folder + "/" + fileName,
		ResourceType: "image",
	})
	if err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	// log.Printf("Error : ", err)
	if err := handler.service.Create(echoCtx.Request().Context(), fileEntity); err != nil {
		errorResponse := buildErrorResponse(nethttp.StatusInternalServerError, err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", resp)
	return echoCtx.JSON(res.Status, res)
}

func GetByteFile(ctx echo.Context) ([]byte, *string, error) {
	fileInput, err := ctx.FormFile("file")
	if err != nil {
		return nil, nil, errors.Wrap(err, "[FileHandler-GetByteFile] FormFile")
	}
	src, err := fileInput.Open()
	if err != nil {
		return nil, nil, errors.Wrap(err, "[FileHandler-GetByteFile] Open")
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return nil, nil, errors.Wrap(err, "[FileHandler-GetByteFile] NewBuffer")
	}
	return buf.Bytes(), &fileInput.Filename, nil
}
