package dynamodb

import (
	"github.com/guregu/dynamo"
)

// Client is an abstraction for dynamo.Client
type Client interface {
	GetTable(tableName string) dynamo.Table
}
