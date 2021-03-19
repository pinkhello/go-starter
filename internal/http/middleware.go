package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go-starter/internal/controller"
	"net/http"
	"strings"
)

type EchoMiddleware struct {
}

func (e *EchoMiddleware) CORS(h echo.HandlerFunc) echo.HandlerFunc {
	cors := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.DELETE, echo.GET, echo.POST, echo.PUT, echo.OPTIONS, echo.HEAD, echo.PATCH},
	})
	return cors(h)
}

func (e *EchoMiddleware) Recover(h echo.HandlerFunc) echo.HandlerFunc {
	recover := middleware.Recover()
	return recover(h)
}

func (e *EchoMiddleware) Logger(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().RequestURI, "swagger") {
			return h(c)
		}
		logEntry(c).Info()
		return h(c)
	}
}

func (e *EchoMiddleware) JWT(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uri := c.Request().RequestURI
		if strings.Compare(uri, "/") == 0 ||
			strings.Compare(uri, "/health") == 0 ||
			strings.Contains(uri, "swagger") {
			return hf(c)
		}
		jwtStr := c.Request().Header.Get("Authorization")
		auths := strings.Split(jwtStr, " ")
		if strings.ToUpper(auths[0]) != "BEARER" || auths[1] == "" {
			return c.JSON(http.StatusUnauthorized, controller.ResponseError{Message: "认证失败"})
		}
		// todo check jwt token
		// todo set jwt info in echo context
		return hf(c)
	}
}

func logEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{})
	}
	return logrus.WithFields(logrus.Fields{
		"method":    c.Request().Method,
		"uri":       c.Request().URL.String(),
		"userAgent": c.Request().UserAgent(),
	})
}

func (e *EchoMiddleware) ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	logEntry(c).Error(report.Message)
	c.Echo().DefaultHTTPErrorHandler(err, c)
}

func InitMiddleware() *EchoMiddleware {
	return &EchoMiddleware{}
}
