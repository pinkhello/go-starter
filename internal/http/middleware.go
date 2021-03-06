package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type EchoMiddleware struct {
}

func (em *EchoMiddleware) CORS(h echo.HandlerFunc) echo.HandlerFunc {
	cors := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.DELETE, echo.GET, echo.POST, echo.PUT, echo.OPTIONS, echo.HEAD, echo.PATCH},
	})
	return cors(h)
}

func (em *EchoMiddleware) Recover(h echo.HandlerFunc) echo.HandlerFunc {
	recover := middleware.Recover()
	return recover(h)
}

func (em *EchoMiddleware) Logger(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().RequestURI, "swagger") {
			return h(c)
		}
		makeLogEntry(c).Info("incoming request")
		return h(c)
	}
}

func makeLogEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	return logrus.WithFields(logrus.Fields{
		"at":        time.Now().Format("2006-01-02 15:04:05"),
		"method":    c.Request().Method,
		"uri":       c.Request().URL.String(),
		"ip":        c.Request().RemoteAddr,
		"userAgent": c.Request().UserAgent(),
	})
}

func (em *EchoMiddleware) ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	makeLogEntry(c).Error(report.Message)
	c.Echo().DefaultHTTPErrorHandler(err, c)
}

func InitMiddleware() *EchoMiddleware {
	return &EchoMiddleware{}
}
