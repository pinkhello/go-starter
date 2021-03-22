package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/nsqio/go-nsq"
	"go-starter/internal/service"
	"go-starter/utils"
	"net/http"
	"strconv"
)

type BusinessGroupController struct {
	BusinessGroupService service.BusinessGroupService
	producer *nsq.Producer
}

func InitBusinessGroupController(e *echo.Echo, groupService service.BusinessGroupService, producer *nsq.Producer) {
	controller := &BusinessGroupController{
		BusinessGroupService: groupService,
		producer: producer,
	}
	g := e.Group("/business_groups")
	g.GET("/:id", controller.GetByID)
	g.GET("/producer", controller.TestProduce)
}

// GetByID godoc
// @Summary Get BusinessGroup By ID
// @Description Get BusinessGroup By ID
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
func (b *BusinessGroupController) GetByID(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	id := int64(idParam)
	bg, err := b.BusinessGroupService.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, SimpleResponse{Data: bg, Message: "success"})
}

// GetByID godoc
// @Summary 测试 NSQ Producer
// @Description 测试 NSQ Producer
// @Tags BusinessGroup
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} SimpleResponse{data=string} "操作信息"
// @Failure 400,401,404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /business_groups/producer [get]
func (b *BusinessGroupController) TestProduce(c echo.Context) error {
	str := "hello-test"
	if err := b.producer.Publish("hello.test.topic", []byte(str)); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, SimpleResponse{Message: "success"})
}
