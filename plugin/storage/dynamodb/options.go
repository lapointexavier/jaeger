package dynamodb

import (
	"flag"
	"time"

	"github.com/jaegertracing/jaeger/pkg/dynamodb/config"
	"github.com/spf13/viper"
)

// Options contains various type of Dynamodb configs and provides the ability
// to bind them to command line flag and apply overlays, so that some configurations
// (e.g. archive) may be underspecified and infer the rest of its parameters from primary.
type Options struct {
	primary                *namespaceConfig
	others                 map[string]*namespaceConfig
	SpanStoreWriteCacheTTL time.Duration
	DepStoreDataFrequency  time.Duration
}

// the Servers field in config.Configuration is a list, which we cannot represent with flags.
// This struct adds a plain string field that can be bound to flags and is then parsed when
// preparing the actual config.Configuration.
type namespaceConfig struct {
	config.Configuration
	servers   string
	namespace string
}

// NewOptions creates a new Options struct.
func NewOptions(primaryNamespace string, otherNamespaces ...string) *Options {
	options := &Options{
		primary: &namespaceConfig{
			Configuration: config.Configuration{},
			servers:       "127.0.0.1",
			namespace:     primaryNamespace,
		},
		others:                 make(map[string]*namespaceConfig, len(otherNamespaces)),
		SpanStoreWriteCacheTTL: time.Hour * 12,
		DepStoreDataFrequency:  time.Hour * 24,
	}

	for _, namespace := range otherNamespaces {
		options.others[namespace] = &namespaceConfig{namespace: namespace}
	}

	return options
}

// AddFlags adds flags for Options
func (opt *Options) AddFlags(flagSet *flag.FlagSet) {
	addFlags(flagSet, opt.primary)
	for _, cfg := range opt.others {
		addFlags(flagSet, cfg)
	}
}

func addFlags(flagSet *flag.FlagSet, nsConfig *namespaceConfig) {
	flagSet.String(
		nsConfig.namespace+".region",
		nsConfig.Region,
		"The AWS region",
	)
	flagSet.String(
		nsConfig.namespace+".endpoint",
		nsConfig.Endpoint,
		"The Dynamodb endpoint (eg: dynamodb.us-west-1.amazonaws.com)",
	)
}

// InitFromViper initializes Options with properties from viper
func (opt *Options) InitFromViper(v *viper.Viper) {
	initFromViper(opt.primary, v)
	for _, cfg := range opt.others {
		initFromViper(cfg, v)
	}
}

func initFromViper(cfg *namespaceConfig, v *viper.Viper) {
	cfg.Endpoint = v.GetString(cfg.namespace + ".endpoint")
	cfg.Region = v.GetString(cfg.namespace + ".region")
}

// GetPrimary returns primary configuration.
func (opt *Options) GetPrimary() *config.Configuration {
	return &opt.primary.Configuration
}

// Get returns auxiliary named configuration.
func (opt *Options) Get(namespace string) *config.Configuration {
	nsCfg, ok := opt.others[namespace]
	if !ok {
		nsCfg = &namespaceConfig{}
		opt.others[namespace] = nsCfg
	}
	nsCfg.Configuration.ApplyDefaults(&opt.primary.Configuration)

	return &nsCfg.Configuration
}
