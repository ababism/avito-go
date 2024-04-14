package banner

import (
	"avito-go/pkg/xapp"
	"avito-go/services/banner/internal/domain"
	"avito-go/services/banner/internal/handler/http/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func parseRolesFromToken(c *gin.Context, token *string) ([]string, error) {
	if token == nil {
		return make([]string, 0), xapp.NewError(http.StatusUnauthorized, "token is required", "token is required", nil)
	}
	switch *token {
	case "user_token":
		return []string{domain.UserRole}, nil
	case "admin_token":
		return []string{domain.AdminRole}, nil
	default:
		return make([]string, 0), xapp.NewError(http.StatusUnauthorized, "invalid token", "invalid token", nil)
	}
}

func NewActorFromToken(c *gin.Context, token *string) (domain.Actor, error) {
	roles, err := parseRolesFromToken(c, token)
	if err != nil {
		return domain.Actor{}, err
	}
	return domain.NewActor(roles), nil

}

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
