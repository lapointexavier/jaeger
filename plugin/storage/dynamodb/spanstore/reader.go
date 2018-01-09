package spanstore

import (
	"github.com/guregu/dynamo"
	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/pkg/dynamodb"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
)

type serviceNamesReader func() ([]string, error)

type operationNamesReader func(service string) ([]string, error)

// SpanReader can query for and load traces from Dynamodb.
type SpanReader struct {
	client               *dynamo.DB
	serviceNamesReader   serviceNamesReader
	operationNamesReader operationNamesReader
	logger               *zap.Logger
}

// NewSpanReader returns a new SpanReader.
func NewSpanReader(
	client dynamodb.Client,
	metricsFactory metrics.Factory,
	logger *zap.Logger,
) *SpanReader {
	return &SpanReader{}
}

// GetServices returns all services traced by Jaeger
func (s *SpanReader) GetServices() ([]string, error) {
	return s.serviceNamesReader()

}

// GetOperations returns all operations for a specific service traced by Jaeger
func (s *SpanReader) GetOperations(service string) ([]string, error) {
	return s.operationNamesReader(service)
}

// GetTrace takes a traceID and returns a Trace associated with that traceID
func (s *SpanReader) GetTrace(traceID model.TraceID) (*model.Trace, error) {
	var retMe *model.Trace
	return retMe, nil
}

// FindTraces retrieves traces that match the traceQuery
func (s *SpanReader) FindTraces(traceQuery *spanstore.TraceQueryParameters) ([]*model.Trace, error) {
	var retMe []*model.Trace
	return retMe, nil
}
