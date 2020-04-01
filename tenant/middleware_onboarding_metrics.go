package tenant

import (
	"context"

	"github.com/influxdata/influxdb"
	"github.com/influxdata/influxdb/kit/metric"
	"github.com/influxdata/influxdb/kit/prom"
)

var _ influxdb.OnboardingService = (*OnboardingMetrics)(nil)

type OnboardingMetrics struct {
	// RED metrics
	rec *metric.REDClient

	onboardingService influxdb.OnboardingService
}

// NewOnboardingMetrics returns a metrics service middleware for the User Service.
func NewOnboardingMetrics(reg *prom.Registry, s influxdb.OnboardingService, opts ...MetricsOption) *OnboardingMetrics {
	o := applyOpts(opts...)
	return &OnboardingMetrics{
		rec:               metric.New(reg, o.applySuffix("user")),
		onboardingService: s,
	}
}

func (m *OnboardingMetrics) IsOnboarding(ctx context.Context) (bool, error) {
	rec := m.rec.Record("is_onboarding")
	available, err := m.onboardingService.IsOnboarding(ctx)
	return available, rec(err)
}

func (m *OnboardingMetrics) OnboardInitialUser(ctx context.Context, req *influxdb.OnboardingRequest) (*influxdb.OnboardingResults, error) {
	rec := m.rec.Record("onboard_initial_user")
	res, err := m.onboardingService.OnboardInitialUser(ctx, req)
	return res, rec(err)
}
