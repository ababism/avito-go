package deleter

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

type PostponeDeleter struct {
	stop             chan bool
	bannerRepository *ports.BannerRepository
	logger           *zap.Logger
}

func New(logger *zap.Logger, bannerRepository *ports.BannerRepository) *PostponeDeleter {
	return &PostponeDeleter{
		logger:           logger,
		bannerRepository: bannerRepository,
		stop:             make(chan bool)}
}

func (s *PostponeDeleter) stopCallback(ctx context.Context) error {
	s.stop <- true
	return nil
}

func (s PostponeDeleter) StopFunc() func(context.Context) error {
	return s.stopCallback
}

func (s *PostponeDeleter) Start(scrapeInterval time.Duration) {
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

func (s *PostponeDeleter) refresh(scrapeInterval time.Duration) {
	initialCtx := context.Background()

	requestIdCtx := WithRequestID(initialCtx)
	ctxLogger := zapctx.WithLogger(requestIdCtx, s.logger)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(ctxLogger, "banner/daemon/refresher.refresh", trace.WithNewRoot())
	defer span.End()

	err := s.bannerRepository.ProcessDeleteQueue(ctx)
	if err != nil {
		s.logger.Error("failed to refresh cache", zap.Error(err))
		return
	}

	time.Sleep(scrapeInterval)
}
