package app

import (
	"log/slog"
	"net/http"
	"time"

	"gitdev.devops.krungthai.com/techcoach/template/goapi.git/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Context interface {
	Bind(v any) error
	OK(v any)
	BadRequest(err error)
	StoreError(err error)
}

type context struct {
	*gin.Context
	logHandler slog.Handler
}

func NewContext(c *gin.Context, logHandler slog.Handler) Context {
	return &context{Context: c, logHandler: logHandler}
}

func (c *context) Bind(v any) error {
	return c.Context.ShouldBindJSON(v)
}

func (c *context) OK(v any) {
	if v == nil {
		c.Context.Status(http.StatusOK)
		return
	}
	c.Context.JSON(http.StatusOK, Response{
		Status: Success,
		Data:   v,
	})
}

func (c *context) BadRequest(err error) {
	logger.AppErrorf(c.logHandler, "%s", err)
	c.Context.JSON(http.StatusBadRequest, Response{
		Status:  Fail,
		Message: err.Error(),
	})
}

func (c *context) StoreError(err error) {
	logger.AppErrorf(c.logHandler, "%s", err)
	c.Context.JSON(storeErrorStutas, Response{
		Status:  Fail,
		Message: err.Error(),
	})
}

func NewGinHandler(handler func(Context), logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewContext(c, logger.Handler().WithAttrs([]slog.Attr{slog.String("transaction-id", c.Request.Header.Get("transaction-id"))})))
		// handler(NewContext(c, logger.With(zap.String("transaction-id", c.Request.Header.Get("transaction-id")))))
	}
}

type Router struct {
	*gin.Engine
	logger *slog.Logger
}

func NewRouter(logger *slog.Logger) *Router {
	r := gin.Default()

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type", "TransactionID"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	return &Router{Engine: r, logger: logger}
}

func (r *Router) GET(path string, handler func(Context)) {
	r.Engine.GET(path, NewGinHandler(handler, r.logger))
}

func (r *Router) POST(path string, handler func(Context)) {
	r.Engine.POST(path, NewGinHandler(handler, r.logger))
}
