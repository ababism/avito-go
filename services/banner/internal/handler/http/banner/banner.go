package banner

import (
	"avito-go/pkg/xapp"
	"avito-go/services/banner/internal/domain"
	"avito-go/services/banner/internal/handler/http/banner/oapi"
	generated "avito-go/services/banner/internal/handler/http/banner/oapi"
	"avito-go/services/banner/internal/handler/http/models"
	"avito-go/services/banner/internal/service/ports"
	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
)

var _ oapi.ServerInterface = &Handler{}

type Handler struct {
	logger *zap.Logger
	banner ports.BannerService
}

func (h Handler) spanName(funcName string) string {
	return "banner/handler." + funcName
}

func NewHandler(logger *zap.Logger, bannerService ports.BannerService) *Handler {
	return &Handler{logger: logger, banner: bannerService}
}

func (h Handler) GetBanner(ginCtx *gin.Context, params generated.GetBannerParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, h.spanName("GetBanner"))
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	actor, err := NewActorFromToken(ginCtx, params.Token)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	banners, err := h.banner.GetList(ctx, actor, params.ToDomain())

	if banners == nil {
		err := xapp.NewError(http.StatusInternalServerError, "nil banners without error", "get banners returned nil lesson without error", nil)
		h.logger.Error("nil lesson", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
		return
	}

	resp := models.ToBannerListResponse(banners)

	ginCtx.JSON(http.StatusOK, resp)
}

func (h Handler) PostBanner(ginCtx *gin.Context, params generated.PostBannerParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, h.spanName("PostBanner"))
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	actor, err := NewActorFromToken(ginCtx, params.Token)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	var payload generated.PostBannerJSONBody
	h.bindRequestBody(ginCtx, &payload)

	bannerPayload, err := payload.ToDomain()
	if err != nil {
		h.abortWithAutoResponse(ginCtx, xapp.NewError(http.StatusBadRequest, "invalid request body", "request body", err))
		return
	}

	bannerID, err := h.banner.Create(ctx, actor, bannerPayload)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	resp := models.ToBannerResponse(domain.Banner{ID: bannerID})

	ginCtx.JSON(http.StatusOK, resp)
}

func (h Handler) DeleteBannerId(ginCtx *gin.Context, id int, params generated.DeleteBannerIdParams) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PatchBannerId(ginCtx *gin.Context, id int, params generated.PatchBannerIdParams) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetUserBanner(ginCtx *gin.Context, params generated.GetUserBannerParams) {
	//TODO implement me
	panic("implement me")
}

//
//func (h Handler) GetCoursesEditList(ginCtx *gin.Context, params generated.GetCoursesEditListParams) {
//	tr := global.Tracer(domain.ServiceName)
//	ctxTrace, span := tr.Start(ginCtx, "banner/handler.GetCoursesEditList")
//	defer span.End()
//
//	ctx := zapctx.WithLogger(ctxTrace, h.logger)
//
//	roles, err := h.banner.GetActorRoles(ctx, params.Actor.ToDomain())
//
//	course, err := h.banner.GetAllCoursesTemplates(ctx, params.Actor.ToDomainWithRoles(roles), params.Offset, params.Limit)
//	if err != nil {
//		h.abortWithAutoResponse(ginCtx, err)
//		return
//	}
//
//	if course == nil {
//		err := xapp.NewError(http.StatusInternalServerError, "nil banner without error", "GetCourse returned nil banner without error", nil)
//		h.logger.Error("nil course", zap.Error(err))
//		http2.AbortWithBadResponse(ginCtx, h.logger, http2.MapErrorToCode(err), err)
//	}
//	resp := models.ToCourseListResponse(course)
//
//	ginCtx.JSON(http.StatusOK, resp)
//}
//
//// PostCoursesEdit CREATE course
//func (h Handler) PostCoursesEdit(ginCtx *gin.Context, params generated.PostCoursesEditParams) {
//	tr := global.Tracer(domain.ServiceName)
//	ctxTrace, span := tr.Start(ginCtx, "banner/handler.PostCoursesEdit")
//	defer span.End()
//
//	ctx := zapctx.WithLogger(ctxTrace, h.logger)
//
//	var coursePayload generated.Course
//	h.bindRequestBody(ginCtx, &coursePayload)
//
//	roles, err := h.banner.GetActorRoles(ctx, params.Actor.ToDomain())
//
//	course, err := h.banner.CreateCourse(ctx, params.Actor.ToDomainWithRoles(roles), coursePayload.ToDomain())
//	if err != nil {
//		h.abortWithAutoResponse(ginCtx, err)
//		return
//	}
//
//	if course == nil {
//		err := xapp.NewError(http.StatusInternalServerError, "nil course without error", "GetCourse returned nil course without error", nil)
//		h.logger.Error("nil course", zap.Error(err))
//		http2.AbortWithBadResponse(ginCtx, h.logger, http2.MapErrorToCode(err), err)
//	}
//	resp := models.ToCourseResponse(*course)
//
//	ginCtx.JSON(http.StatusOK, resp)
//}
