package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-starter/internal/service"
	"go-starter/utils"
	"net/http"
	"strconv"
)

type BusinessGroupController struct {
	BusinessGroupService service.BusinessGroupService
}

func InitBusinessGroupController(e *echo.Echo, groupService service.BusinessGroupService) {
	controller := &BusinessGroupController{
		BusinessGroupService: groupService,
	}
	e.GET("/business_groups", controller.FetchBusinessGroup)
	e.GET("/business_groups/:id", controller.GetById)
}

type ResponseError struct {
	Message string `json:"message"`
}

func (b *BusinessGroupController) FetchBusinessGroup(c echo.Context) error {
	nums := c.QueryParam("num")
	num, _ := strconv.Atoi(nums)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listBg, nextCursor, err := b.BusinessGroupService.Fetch(ctx, cursor, int64(num))
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listBg)
}

func (b *BusinessGroupController) GetById(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	id := int64(idParam)
	bg, err := b.BusinessGroupService.GetById(ctx, id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, bg)
}
