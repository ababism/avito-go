package postgre

import (
	"avito-go/pkg/xapp"
	"avito-go/pkg/xcommon"
	"avito-go/services/banner/internal/domain"
	"avito-go/services/banner/internal/repository/postgre/models"
	"avito-go/services/banner/internal/service/ports"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

const (
	spanBaseName = "banner/repository/postgres."
)

var _ ports.BannerRepository = &bannerRepository{}

func NewBannerRepository(db *sqlx.DB) ports.BannerRepository {
	return &bannerRepository{db: db,
		spanName: spanBaseName}
}

type bannerRepository struct {
	db       *sqlx.DB
	spanName string
}

func (r bannerRepository) Find(initialCtx context.Context, tag, feature int) (domain.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func IntSliceToPostgresArray(slice []int) string {
	return "{" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), ","), "[]") + "}"
}

func (r bannerRepository) Get(initialCtx context.Context, ID int) (domain.Banner, error) {
	logger := zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	_, span := tr.Start(initialCtx, r.spanName+"Create")
	defer span.End()

	q := `
	SELECT * FROM banner WHERE id = $1
	`
	logger.With(zap.String("PSQL query", formatQuery(q)))

	var banner models.Banner
	err := r.db.Get(&banner, q, ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Banner{}, xapp.NewError(http.StatusNotFound, "banner not found", "failed to get banner from DB", err)
		}
		return domain.Banner{}, xapp.NewError(http.StatusInternalServerError, "unknown error", "postgres internal error", err)
	}
	return banner.ToDomain(), nil
}

func (r bannerRepository) Create(initialCtx context.Context, banner domain.Banner) (domain.Banner, error) {
	logger := zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, r.spanName+"Create")
	defer span.End()

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return domain.Banner{}, xapp.NewError(http.StatusInternalServerError, "unknown error", "failed to start transaction", err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	q := `
	INSERT INTO banner (content, is_active)
	VALUES ($1, $2)
	RETURNING *;
	`
	logger.With(zap.String("PSQL query", formatQuery(q)))

	writeBanner := models.ToBannerModel(banner)
	var resBanner models.Banner
	//res, err := r.db.NamedExec(q, resBanner)
	row := tx.QueryRow(q, writeBanner.Content, writeBanner.IsActive)
	if err != nil {
		_ = tx.Rollback()
		return domain.Banner{}, xapp.NewError(http.StatusInternalServerError, "unknown error", "postgres internal error", err)
	}

	err = row.Scan(&resBanner.ID, &resBanner.Content, &resBanner.IsActive, &resBanner.CreatedAt, &resBanner.UpdatedAt)
	if err != nil {
		_ = tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Banner{}, xapp.NewError(http.StatusInternalServerError, "unknown error", "failed to create banner in psql", err)

		}
		return domain.Banner{}, xapp.NewError(http.StatusInternalServerError, "unknown error", "can't scan banner", err)
	}

	qFeature := `
	INSERT INTO banners_feature (banner_id, feature_id)
	VALUES ($1, $2);
	`

	//res, err := r.db.NamedExec(q, resBanner)
	_, err = tx.Exec(qFeature, resBanner.ID, banner.Feature)
	if err != nil {
		_ = tx.Rollback()
		return domain.Banner{}, xapp.NewError(http.StatusInternalServerError, "unknown error", "postgres internal error", err)
	}

	qTags := `
	INSERT INTO banners_tags (banner_id, tag_id)
	SELECT $1 AS banner_id, tag_id
	FROM unnest($2::int[]) AS tag_id;
	`
	tagIDs := IntSliceToPostgresArray(banner.Tags)
	//res, err := r.db.NamedExec(q, resBanner)
	_, err = tx.Exec(qTags, resBanner.ID, tagIDs)
	if err != nil {
		_ = tx.Rollback()
		return domain.Banner{}, xapp.NewError(http.StatusInternalServerError, "unknown error", "postgres internal error", err)
	}

	err = tx.Commit()
	if err != nil {
		return domain.Banner{}, xapp.NewError(http.StatusInternalServerError, "unknown error", "failed to commit transaction", err)
	}

	return resBanner.ToDomain(), nil
}

func mapGetBankAccountRequestParams(params *domain.BannerFilter) map[string]any {
	if params == nil {
		return map[string]any{}
	}
	paramsMap := make(map[string]any)
	if params.TagID != nil {
		paramsMap["tag_id"] = params.TagID
	}
	if params.Feature != nil {
		paramsMap["feature"] = params.Feature
	}
	return paramsMap
}

func (r bannerRepository) GetList(initialCtx context.Context, filter domain.BannerFilter) ([]domain.Banner, error) {
	logger := zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, r.spanName+"Create")
	defer span.End()

	query := `
	SELECT * FROM banner WHERE id = $1
	`

	paramsMap := mapGetBankAccountRequestParams(&filter)
	q, args := xcommon.QueryWhereAnd(
		query,
		paramsMap,
	)
	logger.With(zap.String("PSQL query", formatQuery(q)))

	var banners []models.Banner
	err := r.db.SelectContext(ctx, &banners, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, xapp.NewError(http.StatusNotFound, "banner not found", "failed to get banner from DB", err)
		}
		return nil, xapp.NewError(http.StatusInternalServerError, "unknown error", "postgres internal error", err)

	}

	domainBanners := make([]domain.Banner, 0, len(banners))
	for _, banner := range banners {
		domainBanners = append(domainBanners, banner.ToDomain())
	}

	return domainBanners, nil
}

func (r bannerRepository) ProcessDeleteQueue(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r bannerRepository) Update(ctx context.Context, bannerID int, banner domain.Banner) (domain.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (r bannerRepository) Delete(ctx context.Context, bannerID int) error {
	//TODO implement me
	panic("implement me")
}
