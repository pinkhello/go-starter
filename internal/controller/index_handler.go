package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type IndexController struct {
}

func InitIndexController(e *echo.Echo) {
	controller := &IndexController{}
	e.GET("/health", controller.Health)
	e.GET("/", controller.Health)
}

func (index *IndexController) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, SimpleResponse{
		Message: "success",
		Data:    time.Now().Format("2006-01-02 15:04:05"),
	})
}
