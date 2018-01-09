package dynamodb

import "github.com/guregu/dynamo"

// DDBClient is a wrapper around dynamo.DB
type DDBClient struct {
	client *dynamo.DB
}

// WrapDDBClient creates a DDBClient out of *dynamo.DB.
func WrapDDBClient(client *dynamo.DB) DDBClient {
	return DDBClient{client: client}
}

// GetTable returns a dynamo.Table
func (c DDBClient) GetTable(tableName string) dynamo.Table {
	return c.client.Table(tableName)
}
