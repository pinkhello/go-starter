package controller

import (
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
	g := e.Group("/business_groups")
	g.GET("/:id", controller.GetById)
}

// GetById godoc
// @Summary Get BusinessGroup By Id
// @Description Get BusinessGroup By Id
// @Tags BusinessGroup
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Security ApiKeyAuth
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} SimpleResponse{data=models.BusinessGroup} "BusinessGroup Info"
// @Failure 400,401,404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /business_groups/{id} [get]
func (b *BusinessGroupController) GetById(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	id := int64(idParam)
	bg, err := b.BusinessGroupService.GetById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, SimpleResponse{Data: bg, Message: "success"})
}
