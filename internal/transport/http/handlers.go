package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AdiKhoironHasan/go-kampus-auth/internal/services"
	authConst "github.com/AdiKhoironHasan/go-kampus-auth/pkg/common/const"
	"github.com/AdiKhoironHasan/go-kampus-auth/pkg/dto"
	matkulErrors "github.com/AdiKhoironHasan/go-kampus-auth/pkg/errors"
	"github.com/apex/log"

	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	// "github.com/dgrijalva/jwt-go"
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

	eJWT := e.Group("api/v1/mustlogin")
	eJWT.Use(middleware.JWT([]byte("the_secret_key")))
	eJWT.GET("/check", handler.Ping)
}

func (h *HttpHandler) Ping(c echo.Context) error {

	// middleware.
	// jwt.
	fmt.Println("jwt: ", middleware.JWT([]byte("the_secret_key")))

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
	var data *dto.UserResponse

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

	data, err = h.service.Login(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if data == nil {
		var resp = dto.ResponseDTO{
			Success: false,
			Message: authConst.LoginFailed,
			Data:    nil,
		}

		return c.JSON(http.StatusUnauthorized, resp)
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: authConst.LoginSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
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
