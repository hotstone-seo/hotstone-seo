package metric

import (
	"context"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/hotstone-seo/hotstone-seo/app/config"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.uber.org/dig"
)

type contextKey string

var startTimeKey = contextKey("startTime")

var (
	// MLatencyMs records the time it took for matching operation
	MLatencyMs = stats.Float64("matching/latency", "Latency of matching operation", "ms")

	// KeyIsMatched holds matching operation result. It will be either "matched" or "mismatched"
	KeyIsMatched, _ = tag.NewKey("is_matched")

	// KeyMismatchedPath holds information of the mismatched path
	KeyMismatchedPath, _ = tag.NewKey("mismatched_path")
)

func sinceInMilliseconds(startTime time.Time) float64 {
	return float64(time.Since(startTime).Nanoseconds()) / 1e6
}

// InitializeLatencyTracking ...
func InitializeLatencyTracking(ctx context.Context) context.Context {
	return context.WithValue(ctx, startTimeKey, time.Now())
}

// recordLatency ...
func recordLatency(ctx context.Context) {
	startTimeVal := ctx.Value(startTimeKey)
	if startTime, ok := startTimeVal.(time.Time); ok {
		stats.Record(ctx, MLatencyMs.M(sinceInMilliseconds(startTime)))
	}
}

func recordIsMatched(ctx context.Context, isMatched string) {
	// TODO
}

func recordMismatchedPath(ctx context.Context, mismatchedPath string) {
	// TODO
}

func RecordMatched(ctx context.Context) {
	recordIsMatched(ctx, "matched")
}

func RecordMismatched(ctx context.Context, mismatchedPath string) {
	recordIsMatched(ctx, "mismatched")
}

// AddIsMatchedTag ...
func AddIsMatchedTag(ctx context.Context, target string) (context.Context, error) {
	return tag.New(ctx, tag.Upsert(KeyIsMatched, target))
}

// AddMismatchedTag ...
func AddMismatchedTag(ctx context.Context, reason string) (context.Context, error) {
	return tag.New(ctx, tag.Upsert(KeyMismatchedPath, reason))
}

type MetricServer interface {
	Start() error
}

type MetricServerImpl struct {
	dig.In
	config.Config
}

func NewMetricServer(server MetricServerImpl) MetricServer {
	return &server
}

func (s *MetricServerImpl) Start() (err error) {

	var (
		// MatchingOperationCountView provide View for request count grouped by target and reason
		MatchingOperationCountView = &view.View{
			Name:        "matching_operation/count",
			Measure:     MLatencyMs,
			Description: "The count of matching operation",
			Aggregation: view.Count(),
			TagKeys:     []tag.Key{KeyIsMatched},
		}

		// MatchingOperationLatencyView provide view for latency count distribution
		MatchingOperationLatencyView = &view.View{
			Name:        "matching_operation/latency",
			Measure:     MLatencyMs,
			Description: "The latency distribution per matching operation",

			// Latency in buckets:
			// [>=0ms, >=25ms, >=50ms, >=75ms, >=100ms, >=200ms, >=400ms, >=600ms, >=800ms, >=1s, >=2s, >=4s, >=6s]
			Aggregation: view.Distribution(0, 25, 50, 75, 100, 200, 400, 600, 800, 1000, 2000, 4000, 6000),
			TagKeys:     []tag.Key{KeyIsMatched},
		}

		views = []*view.View{MatchingOperationCountView, MatchingOperationLatencyView}
	)

	if err = view.Register(views...); err != nil {
		return
	}

	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "hotstone_seo",
	})
	if err != nil {
		return
	}

	view.RegisterExporter(pe)
	addr := s.MetricAddress
	mux := http.NewServeMux()
	mux.Handle("/metrics", pe)
	log.Warnf("Metrics endpoint will be running at: %s", addr)

	go func(mux *http.ServeMux, addr string) {
		if err = http.ListenAndServe(addr, mux); err != nil {
			log.Warnf("Failed to run Prometheus scrape endpoint: %v", err)
			return
		}
	}(mux, addr)

	return nil
}
