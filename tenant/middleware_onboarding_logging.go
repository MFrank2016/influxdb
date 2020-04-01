package tenant

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influxdb"
	"go.uber.org/zap"
)

type OnboardingLogger struct {
	logger            *zap.Logger
	onboardingService influxdb.OnboardingService
}

// NewOnboardingLogger returns a logging service middleware for the Bucket Service.
func NewOnboardingLogger(log *zap.Logger, s influxdb.OnboardingService) *OnboardingLogger {
	return &OnboardingLogger{
		logger:            log,
		onboardingService: s,
	}
}

var _ influxdb.OnboardingService = (*OnboardingLogger)(nil)

func (l *OnboardingLogger) IsOnboarding(ctx context.Context) (available bool, err error) {
	defer func(start time.Time) {
		dur := zap.Duration("took", time.Since(start))
		if err != nil {
			l.logger.Error("failed to check onboarding", zap.Error(err), dur)
			return
		}
		l.logger.Info("bucket create", dur)
	}(time.Now())
	return l.onboardingService.IsOnboarding(ctx)
}

func (l *OnboardingLogger) OnboardInitialUser(ctx context.Context, req *influxdb.OnboardingRequest) (res *influxdb.OnboardingResults, err error) {
	defer func(start time.Time) {
		dur := zap.Duration("took", time.Since(start))
		if err != nil {
			msg := fmt.Sprintf("failed to onboard user %s", req.User)
			l.logger.Error(msg, zap.Error(err), dur)
			return
		}
		l.logger.Info("bucket find by ID", dur)
	}(time.Now())
	return l.onboardingService.OnboardInitialUser(ctx, req)
}
