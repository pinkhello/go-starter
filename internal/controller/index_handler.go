package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type IndexController struct {
}

func InitIndexController(e *echo.Echo) {
	controller := &IndexController{}
	e.GET("/health", controller.Health)
}

func (index *IndexController) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
