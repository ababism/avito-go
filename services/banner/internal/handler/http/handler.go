package http

import (
	"avito-go/services/banner/internal/service/ports"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"

	"avito-go/services/banner/internal/config"
	"avito-go/services/banner/internal/handler/http/banner"
	oapigen "avito-go/services/banner/internal/handler/http/banner/oapi"
)

const (
	httpPrefix = "api"
	version    = "1"
)

type Handler struct {
	logger              *zap.Logger
	cfg                 *config.Config
	coursesHandler      *banner.Handler
	userServiceProvider ports.BannerService
}

// HandleError is a sample error handler function
func HandleError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func InitHandler(
	router gin.IRouter,
	logger *zap.Logger,
	middlewares []oapigen.MiddlewareFunc,
	coursesService ports.BannerService,
) {
	coursesHandler := banner.NewHandler(logger, coursesService)

	ginOpts := oapigen.GinServerOptions{
		BaseURL:      fmt.Sprintf("%s/%s", httpPrefix, getVersion()),
		Middlewares:  middlewares,
		ErrorHandler: HandleError,
	}
	oapigen.RegisterHandlersWithOptions(router, coursesHandler, ginOpts)
}

func getVersion() string {
	return fmt.Sprintf("v%s", strings.Split(version, ".")[0])
}
