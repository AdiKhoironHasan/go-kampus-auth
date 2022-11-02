package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AdiKhoironHasan/matkul/internal/services"
	"github.com/AdiKhoironHasan/matkul/pkg/dto"
	matkulErrors "github.com/AdiKhoironHasan/matkul/pkg/errors"
	"github.com/apex/log"

	"github.com/labstack/echo"
)

type HttpHandler struct {
	service services.Services
}

func NewHttpHandler(e *echo.Echo, srv services.Services) {
	handler := &HttpHandler{
		srv,
	}
	g := e.Group("api/v1/auth")
	g.GET("/ping", handler.Ping)
	g.POST("/login", handler.Login)
}

func (h *HttpHandler) Ping(c echo.Context) error {

	version := os.Getenv("VERSION")
	fmt.Println(version)
	if version == "" {
		version = "pong"
	}

	data := version

	return c.JSON(http.StatusOK, data)

}

func (h *HttpHandler) Login(c echo.Context) error {
	postDTO := dto.UserLoginReqDTO{}

	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveMatkul(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// var resp = dto.ResponseDTO{
	// 	Success: true,
	// 	Message: matkulConst.SaveSuccess,
	// 	Data:    nil,
	// }

	return c.JSON(http.StatusOK, nil)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case matkulErrors.ErrInternalServerError:
		return http.StatusInternalServerError
	case matkulErrors.ErrNotFound:
		return http.StatusNotFound
	case matkulErrors.ErrConflict:
		return http.StatusConflict
	case matkulErrors.ErrInvalidRequest:
		return http.StatusBadRequest
	case matkulErrors.ErrFailAuth:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
