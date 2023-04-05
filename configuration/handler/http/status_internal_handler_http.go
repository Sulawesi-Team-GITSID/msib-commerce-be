package http

import (
	"backend-service/entity"
	"net/http"

	"github.com/labstack/echo"
)

// Status returns health check for the service.
func Status(echoCtx echo.Context) error {
	var res = entity.NewResponse(http.StatusOK, "It is work v3!", nil)
	return echoCtx.JSON(res.Status, res)
}
