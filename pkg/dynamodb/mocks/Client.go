package mocks

import (
	"github.com/guregu/dynamo"
	"github.com/jaegertracing/jaeger/pkg/dynamodb"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetTable provides a mock function with given fields: tableName
func (_m *Client) GetTable(tableName string) dynamo.Table {
	ret := _m.Called(tableName)

	var r0 dynamo.Table
	if rf, ok := ret.Get(0).(func(string, ...interface{}) dynamo.Table); ok {
		r0 = rf(tableName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamo.Table)
		}
	}

	return r0
}

var _ dynamodb.Client = (*Client)(nil)
