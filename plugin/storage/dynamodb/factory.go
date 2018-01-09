package dynamodb

import (
	"flag"

	"github.com/spf13/viper"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/pkg/dynamodb"
	"github.com/jaegertracing/jaeger/pkg/dynamodb/config"
	dSpanStore "github.com/jaegertracing/jaeger/plugin/storage/dynamodb/spanstore"
	"github.com/jaegertracing/jaeger/storage/dependencystore"
	"github.com/jaegertracing/jaeger/storage/spanstore"
)

// Factory implements storage.Factory for Dynamodb backend.
type Factory struct {
	Options *Options

	metricsFactory metrics.Factory
	logger         *zap.Logger

	primaryConfig config.ClientBuilder
	primaryClient dynamodb.Client
}

// NewFactory creates a new Factory.
func NewFactory() *Factory {
	return &Factory{
		Options: NewOptions("dynamodb"), // TODO add "dynamodb-archive" once supported
	}
}

// AddFlags implements plugin.Configurable
func (f *Factory) AddFlags(flagSet *flag.FlagSet) {
	f.Options.AddFlags(flagSet)
}

// InitFromViper implements plugin.Configurable
func (f *Factory) InitFromViper(v *viper.Viper) {
	f.Options.InitFromViper(v)
	f.primaryConfig = f.Options.GetPrimary()
}

// Initialize implements storage.Factory
func (f *Factory) Initialize(metricsFactory metrics.Factory, logger *zap.Logger) error {
	f.metricsFactory, f.logger = metricsFactory, logger

	primaryClient, err := f.primaryConfig.NewClient()
	if err != nil {
		return err
	}
	f.primaryClient = primaryClient
	// TODO init archive (cf. https://github.com/jaegertracing/jaeger/pull/604)
	return nil
}

// CreateSpanReader implements storage.Factory
func (f *Factory) CreateSpanReader() (spanstore.Reader, error) {
	return dSpanStore.NewSpanReader(f.primaryClient, f.metricsFactory, f.logger), nil
}

// CreateSpanWriter implements storage.Factory
func (f *Factory) CreateSpanWriter() (spanstore.Writer, error) {
	return dSpanStore.NewSpanWriter(f.primaryClient, f.Options.SpanStoreWriteCacheTTL, f.metricsFactory, f.logger), nil
}

// CreateDependencyReader implements storage.Factory
func (f *Factory) CreateDependencyReader() (dependencystore.Reader, error) {
	return nil, nil //cDepStore.NewDependencyStore(f.primarySession, f.Options.DepStoreDataFrequency, f.metricsFactory, f.logger), nil
}
