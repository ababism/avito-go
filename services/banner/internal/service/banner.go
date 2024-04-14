package service

import (
	"avito-go/services/banner/internal/domain"
	"context"
)

func (s BannerService) Find(ctx context.Context, actor domain.Actor, tag, feature int, useLastRevision bool) (domain.CachedBanner, error) {
	//TODO implement me
	panic("implement me")
}

func (s BannerService) Create(ctx context.Context, actor domain.Actor, course domain.Banner) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s BannerService) GetList(ctx context.Context, actor domain.Actor, filter domain.BannerFilter) ([]domain.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (s BannerService) Update(ctx context.Context, actor domain.Actor, bannerID int, banner domain.Banner) error {
	//TODO implement me
	panic("implement me")
}

func (s BannerService) Delete(ctx context.Context, actor domain.Actor, bannerID int) error {
	//TODO implement me
	panic("implement me")
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
