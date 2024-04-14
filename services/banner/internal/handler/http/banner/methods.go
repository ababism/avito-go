package banner

import (
	"avito-go/pkg/xapp"
	"avito-go/services/banner/internal/handler/http/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AbortWithBadResponse(c *gin.Context, logger *zap.Logger, statusCode int, err error) {
	logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, xapp.GetLastMessage(err)))
	c.AbortWithStatusJSON(statusCode, models.Error{Message: xapp.GetLastMessage(err)})
}

func AbortWithErrorResponse(c *gin.Context, logger *zap.Logger, statusCode int, message string) {
	logger.Error(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, message))
	c.AbortWithStatusJSON(statusCode, models.Error{Message: message})
}

func MapErrorToCode(err error) int {
	return xapp.GetCode(err)
}

func (h Handler) abortWithBadResponse(c *gin.Context, statusCode int, err error) {
	h.logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, xapp.GetLastMessage(err)))
	c.AbortWithStatusJSON(statusCode, models.Error{Message: xapp.GetLastMessage(err)})
}

func (h Handler) abortWithAutoResponse(c *gin.Context, err error) {
	h.logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, xapp.GetCode(err), xapp.GetLastMessage(err)))
	c.AbortWithStatusJSON(xapp.GetCode(err), models.Error{Message: xapp.GetLastMessage(err)})
}

func (h Handler) bindRequestBody(c *gin.Context, obj any) bool {
	if err := c.BindJSON(obj); err != nil {
		AbortWithBadResponse(c, h.logger, http.StatusBadRequest, err)
		return false
	}
	return true
}
