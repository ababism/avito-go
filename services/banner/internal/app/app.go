package app

import (
	"avito-go/pkg/metrics"
	"avito-go/pkg/mylogger"
	"avito-go/pkg/xapp"
	"avito-go/pkg/xdb/xpostgres"
	"avito-go/pkg/xshutdown"
	"avito-go/pkg/xtracer"
	"avito-go/services/banner/internal/config"
	"avito-go/services/banner/internal/daemons/cacherefresher"
	"avito-go/services/banner/internal/repository/cache"
	"avito-go/services/banner/internal/repository/postgre"
	"avito-go/services/banner/internal/service"
	"avito-go/services/banner/internal/service/ports"
	"context"
	"fmt"
	"github.com/juju/zaputil/zapctx"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

type App struct {
	cfg            *config.Config
	address        string
	logger         *zap.Logger
	tracerProvider *trace.TracerProvider
	service        ports.BannerService
	daemon         *cacherefresher.CacheRefresher
}

func NewApp(cfg *config.Config) (*App, error) {
	startCtx := context.Background()
	// INFRASTRUCTURE ----------------------------------------------------------------------

	// Инициализируем logger
	//logger, err := xlogger.Init(cfg.Logger, cfg.App)
	logger, err := mylogger.InitLogger(cfg.Logger, cfg.App.Name)
	if err != nil {
		return nil, err
	}
	// Чистим кэш logger при shutdown
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name: "ZapLoggerCacheWipe",
			FnCtx: func(ctx context.Context) error {
				return logger.Sync()
			},
		})
	logger.Info("Init Logger – success")

	logger.Debug("Logger Debug test")
	logger.Info("Logger Info test")
	logger.Error("Logger Error test")

	// сохраняем logger в контекст
	_ = zapctx.WithLogger(startCtx, logger)

	// Инициализируем обработку ошибок
	err = xapp.InitAppError(cfg.App)
	if err != nil {
		logger.Fatal("while initializing App Error handling package", zap.Error(err))
	}

	//logger.Info("Importing constants from driver openApi – success")

	// Инициализируем трассировку
	tp, err := xtracer.Init(cfg.Tracer, cfg.App)
	if err != nil {
		return nil, err
	}
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name: "OpenTelemetryShutdown",
			FnCtx: func(ctx context.Context) error {
				if err := tp.Shutdown(context.Background()); err != nil {
					logger.Error("Error shutting down tracer provider: %v", zap.Error(err))
					return err
				}
				return nil
			},
		})
	logger.Info("Init Tracer – success")

	// Инициализируем Prometheus
	metrics.InitOnce(cfg.Metrics, logger, metrics.AppInfo{
		Name:        cfg.App.Name,
		Environment: string(cfg.App.Environment),
		Version:     cfg.App.Version,
	})
	logger.Info("Init Metrics – success")

	// REPOSITORY ----------------------------------------------------------------------

	// Инициализация PostgreSQL
	PostgreSQL, PsqlClose, err := xpostgres.NewDB(cfg.Postgres)
	if err != nil {
		logger.Fatal("Error init Postgres DB:", zap.Error(err))
		return nil, errors.Wrap(err, "Init Postgres DB")
	}
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name: "MongoClientDisconnect",
			FnCtx: func(ctx context.Context) error {
				return PsqlClose()
			},
		})
	logger.Info("PosgtreSQL connect – success")

	// Инициализируем кэш

	// All repositories
	bannerRepository := postgre.NewBannerRepository(PostgreSQL)
	bannerCache := cache.New(cfg.Cache)

	// SERVICE LAYER ----------------------------------------------------------------------

	// Service layer
	bannerService := service.NewBannerService(bannerRepository, bannerCache)

	logger.Info(fmt.Sprintf("Init %s – success", cfg.App.Name))

	//CacheRefresher for cache refreshing
	daemon := cacherefresher.New(logger, bannerCache)
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name:  "Data scraper stop",
			FnCtx: daemon.StopFunc(),
		})

	daemonInterval, err := cfg.CacheRefresher.GetIterationInterval()
	if err != nil {
		logger.Fatal("can't parse time from daemon config string:", zap.Error(err))
	}
	logger.Info("CacheRefresher interval – ", zap.Duration("interval", daemonInterval))

	logger.Info("Init CacheRefresher – success")

	// TRANSPORT LAYER ----------------------------------------------------------------------

	// инициализируем адрес сервера
	address := fmt.Sprintf(":%d", cfg.Http.Port)

	return &App{
		cfg:            cfg,
		logger:         logger,
		service:        bannerService,
		address:        address,
		tracerProvider: tp,
		daemon:         daemon,
	}, nil
}
