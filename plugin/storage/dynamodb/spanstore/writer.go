package spanstore

import (
	"time"

	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/pkg/dynamodb"
	"github.com/uber/jaeger-lib/metrics"
)

// SpanWriter handles all writes to Dynamodb for the Jaeger data model
type SpanWriter struct {
	client dynamodb.Client
	logger *zap.Logger
}

// NewSpanWriter returns a SpanWriter
func NewSpanWriter(
	client dynamodb.Client,
	writeCacheTTL time.Duration,
	metricsFactory metrics.Factory,
	logger *zap.Logger,
	// TODO options ...Option,
) *SpanWriter {
	// TODO opts := applyOptions(options...)
	return &SpanWriter{
		client: client,
		logger: logger,
	}
}

// WriteSpan saves the span into Dynamodb
func (s *SpanWriter) WriteSpan(span *model.Span) error {
	println(span)

	// table := client.GetTable("spanstore_services_operations")
	// descr, err := table.Describe().Run()

	// if err != nil {
	// 	logger.Error(err.Error())
	// }

	// logger.Info("----")
	// logger.Info(descr.Name)
	// logger.Info(descr.HashKey)
	// logger.Info(descr.RangeKey)
	// logger.Info("----")

	// TODO do something with span
	return nil
}
