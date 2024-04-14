package service

import (
	"avito-go/pkg/xapp"
	"avito-go/services/banner/internal/domain"
	"context"
	"fmt"
	global "go.opentelemetry.io/otel"
	"net/http"
)

func (s BannerService) Create(initialCtx context.Context, actor domain.Actor, banner domain.Banner) (int, error) {
	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, s.spanName("GetList"))
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasRole(domain.AdminRole) {
		return 0, xapp.NewError(http.StatusForbidden, "user can't create banner",
			fmt.Sprintf("actor do not have %s role", domain.AdminRole), nil)
	}

	bannerInserted, err := s.repository.Create(ctx, banner)
	if err != nil {
		return 0, err
	}
	bannerID := bannerInserted.ID

	return bannerID, nil
}

func (s BannerService) GetList(initialCtx context.Context, actor domain.Actor, filter domain.BannerFilter) ([]domain.Banner, error) {
	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, s.spanName("GetList"))
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.AdminRole, domain.UserRole) {
		return nil, xapp.NewError(http.StatusForbidden, "user does not roles to get banner list",
			fmt.Sprintf("actor do not have %s or %s roles", domain.AdminRole, domain.UserRole), nil)
	}

	banners, err := s.repository.GetList(ctx, filter)
	if err != nil {
		return nil, err
	}

	bannerRes := make([]domain.Banner, 0, len(banners))

	for _, banner := range banners {
		if banner.IsActive || actor.HasRole(domain.AdminRole) {
			bannerRes = append(bannerRes, banner)
		}
	}

	return bannerRes, nil
}

func (s BannerService) Update(initialCtx context.Context, actor domain.Actor, bannerID int, banner domain.Banner) error {
	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, s.spanName("GetList"))
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasRole(domain.AdminRole) {
		return xapp.NewError(http.StatusForbidden, "user can't update banner",
			fmt.Sprintf("actor do not have %s role", domain.AdminRole), nil)
	}

	_, err := s.repository.Update(ctx, bannerID, banner)
	if err != nil {
		return err
	}

	return nil
}

func (s BannerService) Delete(initialCtx context.Context, actor domain.Actor, bannerID int) error {
	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, s.spanName("GetList"))
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasRole(domain.AdminRole) {
		return xapp.NewError(http.StatusForbidden, "user can't delete banner",
			fmt.Sprintf("actor do not have %s role", domain.AdminRole), nil)
	}

	err := s.repository.Delete(ctx, bannerID)
	if err != nil {
		return err
	}

	return nil
}

func (s BannerService) Find(initialCtx context.Context, actor domain.Actor, tag, feature int, useLastRevision bool) (domain.CachedBanner, error) {
	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, s.spanName("GetList"))
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.AdminRole, domain.UserRole) {
		return domain.CachedBanner{}, xapp.NewError(http.StatusForbidden, "user can't find banner",
			fmt.Sprintf("actor do not have %s or %s roles", domain.AdminRole, domain.UserRole), nil)
	}

	cachedBanner, found := s.cache.Get(tag, feature)
	if found {
		return *cachedBanner, nil

	}

	banner, err := s.repository.Find(ctx, tag, feature)
	if err != nil {
		return domain.CachedBanner{}, err
	}

	return domain.CachedBanner{
		ID:       banner.ID,
		Content:  banner.Content,
		IsActive: banner.IsActive,
	}, nil

}

//func (c BannerService) CreateCourse(initialCtx context.Context, actor domain.Actor, course *domain.Course) (*domain.Course, error) {
//	_ = zapctx.Logger(initialCtx)
//
//	tr := global.Tracer(domain.ServiceName)
//	ctx, span := tr.Start(initialCtx, "banner/service.CreateCourse")
//	defer span.End()
//
//	ToSpan(&span, actor)
//
//	if !actor.HasRole(domain.AdminRole) {
//		return nil, xapp.NewError(http.StatusForbidden, "user does not have cache rights to create course",
//			fmt.Sprintf("user do not have %s roles", domain.AdminRole), nil)
//	}
//
//	course.TeacherID = actor.ID
//
//	//if course.ID == uuid.Nil || (course.ID == uuid.UUID{}) {
//	//	course.ID = uuid.New()
//	//}
//
//	err := course.Validate()
//	if err != nil {
//		return nil, err
//	}
//
//	newCourse, err := c.courseEdit.Create(ctx, course)
//	if err != nil {
//		return nil, err
//	}
//
//	span.AddEvent("course created", trace.WithAttributes(attribute.String(keys.CourseIDAttributeKey, newCourse.ID.String())))
//	return newCourse, nil
//}
//
//func (c BannerService) GetAllCoursesTemplates(initialCtx context.Context, actor domain.Actor, offset, limit int) ([]*domain.Course, error) {
//	_ = zapctx.Logger(initialCtx)
//
//	tr := global.Tracer(domain.ServiceName)
//	ctx, span := tr.Start(initialCtx, "banner/service.GetAllCoursesTemplates")
//	defer span.End()
//
//	ToSpan(&span, actor)
//
//	if !actor.HasRole(domain.AdminRole) {
//		return nil, xapp.NewError(http.StatusForbidden, "user does not have rights to read all banner templates",
//			fmt.Sprintf("no rights to read all banner templates with no %s role", domain.AdminRole), nil)
//	}
//
//	courses, err := c.courseEdit.GetAll(ctx, offset, limit)
//	if err != nil {
//		return nil, err
//	}
//
//	return courses, nil
//}
