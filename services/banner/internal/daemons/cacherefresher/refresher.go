package cacherefresher

import (
	"avito-go/services/banner/internal/domain"
	"avito-go/services/banner/internal/domain/keys"
	"avito-go/services/banner/internal/service/ports"
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"runtime"
	"time"
)

type CacheRefresher struct {
	stop        chan bool
	bannerCache ports.BannerCache
	logger      *zap.Logger
}

func New(logger *zap.Logger, cache ports.BannerCache) *CacheRefresher {
	return &CacheRefresher{
		logger:      logger,
		bannerCache: cache,
		stop:        make(chan bool)}
}

func (s *CacheRefresher) stopCallback(ctx context.Context) error {
	s.stop <- true
	return nil
}

func (s CacheRefresher) StopFunc() func(context.Context) error {
	return s.stopCallback
}

func (s *CacheRefresher) Start(scrapeInterval time.Duration) {
	go func() {
		stop := s.stop
		go func() {
			for {
				s.refresh(scrapeInterval)
				runtime.Gosched()
			}
		}()
		<-stop
	}()
}

func generateRequestID() string {
	id := uuid.New()
	return id.String()
}

func WithRequestID(ctx context.Context) context.Context {
	requestID := generateRequestID()
	return context.WithValue(ctx, keys.KeyRequestID, requestID)
}

func (s *CacheRefresher) refresh(scrapeInterval time.Duration) {
	initialCtx := context.Background()

	requestIdCtx := WithRequestID(initialCtx)
	ctxLogger := zapctx.WithLogger(requestIdCtx, s.logger)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(ctxLogger, "banner/daemon/refresher.refresh", trace.WithNewRoot())
	defer span.End()

	err := s.bannerCache.UpdateMostViewed(ctx)
	if err != nil {
		s.logger.Error("failed to refresh cache", zap.Error(err))
		return
	}

	time.Sleep(scrapeInterval)
}
