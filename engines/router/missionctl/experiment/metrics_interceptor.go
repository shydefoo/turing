package experiment

import (
	"context"
	"time"

	"github.com/caraml-dev/turing/engines/experiment/runner"
	"github.com/caraml-dev/turing/engines/router/missionctl/instrumentation"
	"github.com/caraml-dev/turing/engines/router/missionctl/log"
	"github.com/gojek/mlp/api/pkg/instrumentation/metrics"
)

type ctxKey string

const (
	startTimeKey ctxKey = "startTimeKey"
)

// MetricsInterceptor is the structural interceptor used for capturing metrics
// from experiment runs
type MetricsInterceptor struct{}

// NewMetricsInterceptor is a creator for a MetricsInterceptor
func NewMetricsInterceptor() runner.Interceptor {
	return &MetricsInterceptor{}
}

// BeforeDispatch associates the start time to the context
func (i *MetricsInterceptor) BeforeDispatch(
	ctx context.Context,
) context.Context {
	return context.WithValue(ctx, startTimeKey, time.Now())
}

// AfterCompletion logs the time taken for the component to process the request,
// to the metrics collector
func (i *MetricsInterceptor) AfterCompletion(
	ctx context.Context,
	err error,
) {
	labels := map[string]string{
		"status": metrics.GetStatusString(err == nil),
		"engine": "",
	}

	if engine, ok := ctx.Value(runner.ExperimentEngineKey).(string); ok {
		labels["engine"] = engine
	}

	// Get start time
	if startTime, ok := ctx.Value(startTimeKey).(time.Time); ok {
		// Measure the time taken for the experiment run
		err := metrics.Glob().MeasureDurationMsSince(
			instrumentation.ExperimentEngineRequestMs,
			startTime,
			labels,
		)
		if err != nil {
			log.Glob().Errorf(err.Error())
		}
	}
}
