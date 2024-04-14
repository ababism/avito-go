package banner

import (
	"avito-go/services/banner/internal/handler/http/banner/oapi"
	generated "avito-go/services/banner/internal/handler/http/banner/oapi"
	"avito-go/services/banner/internal/service/ports"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ oapi.ServerInterface = &Handler{}

type Handler struct {
	logger *zap.Logger
	banner ports.BannerService
}

func NewHandler(logger *zap.Logger, bannerService ports.BannerService) *Handler {
	return &Handler{logger: logger, banner: bannerService}
}

func (h Handler) GetBanner(c *gin.Context, params generated.GetBannerParams) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostBanner(c *gin.Context, params generated.PostBannerParams) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) DeleteBannerId(c *gin.Context, id int, params generated.DeleteBannerIdParams) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PatchBannerId(c *gin.Context, id int, params generated.PatchBannerIdParams) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetUserBanner(c *gin.Context, params generated.GetUserBannerParams) {
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
