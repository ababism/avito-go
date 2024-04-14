package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	xorerror "xor-go/pkg/xerror"
)

type xorErrorHandler func(ctx *gin.Context, code int, err error)

func (r *HttpResponseWrapper) HandleXorError(ctx *gin.Context, err error) {
	handleXorError(ctx, err, r.HandleError, r.HandleError)
}

func (r *HttpResponseWrapper) HandleXorErrorWithMessage(ctx *gin.Context, err error) {
	handleXorError(ctx, err, r.HandleErrorWithMessage, r.HandleError)
}

func handleXorError(ctx *gin.Context, err error, handler xorErrorHandler, defaultHandler xorErrorHandler) {
	switch {
	case errors.As(err, &xorerror.ValueError{}):
		handler(ctx, http.StatusBadRequest, err)
	default:
		defaultHandler(ctx, http.StatusInternalServerError, err)
	}
}
