package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/jaegertracing/jaeger/pkg/dynamodb"
)

// Configuration describes the configuration properties needed to connect to a Cassandra cluster
type Configuration struct {
	Region   string `validate:"nonzero"`
	Endpoint string `validate:"nonzero"`
}

// ApplyDefaults copies settings from source unless its own value is non-zero.
func (c *Configuration) ApplyDefaults(source *Configuration) {

}

// ClientBuilder creates new Dynamo DB instance
type ClientBuilder interface {
	NewClient() (dynamodb.Client, error)
}

// NewClient creates a new Dynamo DB instance
func (c *Configuration) NewClient() (dynamodb.Client, error) {

	awsConf := aws.NewConfig()
	awsConf.WithRegion(c.Region).WithEndpoint(c.Endpoint)

	db := dynamo.New(session.New(), awsConf)

	return dynamodb.WrapDDBClient(db), nil
}
