package server

import (
	"net/http"
	"time"

	"github.com/caturarp/laporplat/handler"
	"github.com/caturarp/laporplat/logger"
	"github.com/caturarp/laporplat/middleware"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	UserHandler   *handler.UserHandler
	AuthHandler   *handler.AuthHandler
	ReportHandler *handler.ReportHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.ContextWithFallback = true
	router.Use(gin.Recovery())

	router.GET("/hello", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"data": "hello world",
		})
	})
	router.Use(middleware.CorsHandler())
	router.Use(middleware.Logger(logger.NewLogger()))

	report := router.Group("/reports")
	report.GET("/", opts.ReportHandler.ListReport)
	report.GET("/:id", opts.ReportHandler.FindReportByID)
	report.POST("/", opts.ReportHandler.CreateReport)
	report.PUT("/:id", opts.ReportHandler.UpdateReport)
	report.DELETE("/:id", opts.ReportHandler.DeleteReport)

	router.Use(middleware.AuthorizeHandler())
	router.Use(middleware.ErrorHandler())

	users := router.Group("/users")
	users.GET("/", opts.UserHandler.ListUser)
	users.GET("/find", opts.UserHandler.FindUser)
	users.GET("/detail", opts.UserHandler.GetUserDetail)
	users.PUT("/detail", opts.UserHandler.UpdateUserDetail)

	auth := router.Group("/auth")
	auth.POST("/login", opts.AuthHandler.Login)
	auth.POST("/register", opts.AuthHandler.RequestRegister)
	auth.POST("/verify", opts.AuthHandler.VerifyRegister)

	return router
}
