// Code generated by mockery v1.0.0. DO NOT EDIT.

package client_test

import (
	access "github.com/onflow/flow/protobuf/go/flow/access"

	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockRPCClient is an autogenerated mock type for the RPCClient type
type MockRPCClient struct {
	mock.Mock
}

// ExecuteScriptAtBlockHeight provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) ExecuteScriptAtBlockHeight(ctx context.Context, in *access.ExecuteScriptAtBlockHeightRequest, opts ...grpc.CallOption) (*access.ExecuteScriptResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.ExecuteScriptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.ExecuteScriptAtBlockHeightRequest, ...grpc.CallOption) *access.ExecuteScriptResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ExecuteScriptResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.ExecuteScriptAtBlockHeightRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteScriptAtBlockID provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) ExecuteScriptAtBlockID(ctx context.Context, in *access.ExecuteScriptAtBlockIDRequest, opts ...grpc.CallOption) (*access.ExecuteScriptResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.ExecuteScriptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.ExecuteScriptAtBlockIDRequest, ...grpc.CallOption) *access.ExecuteScriptResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ExecuteScriptResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.ExecuteScriptAtBlockIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteScriptAtLatestBlock provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) ExecuteScriptAtLatestBlock(ctx context.Context, in *access.ExecuteScriptAtLatestBlockRequest, opts ...grpc.CallOption) (*access.ExecuteScriptResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.ExecuteScriptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.ExecuteScriptAtLatestBlockRequest, ...grpc.CallOption) *access.ExecuteScriptResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ExecuteScriptResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.ExecuteScriptAtLatestBlockRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetAccount(ctx context.Context, in *access.GetAccountRequest, opts ...grpc.CallOption) (*access.GetAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.GetAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetAccountRequest, ...grpc.CallOption) *access.GetAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.GetAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountAtBlockHeight provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetAccountAtBlockHeight(ctx context.Context, in *access.GetAccountAtBlockHeightRequest, opts ...grpc.CallOption) (*access.AccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.AccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetAccountAtBlockHeightRequest, ...grpc.CallOption) *access.AccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.AccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetAccountAtBlockHeightRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountAtLatestBlock provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetAccountAtLatestBlock(ctx context.Context, in *access.GetAccountAtLatestBlockRequest, opts ...grpc.CallOption) (*access.AccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.AccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetAccountAtLatestBlockRequest, ...grpc.CallOption) *access.AccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.AccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetAccountAtLatestBlockRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByHeight provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetBlockByHeight(ctx context.Context, in *access.GetBlockByHeightRequest, opts ...grpc.CallOption) (*access.BlockResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.BlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetBlockByHeightRequest, ...grpc.CallOption) *access.BlockResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetBlockByHeightRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByID provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetBlockByID(ctx context.Context, in *access.GetBlockByIDRequest, opts ...grpc.CallOption) (*access.BlockResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.BlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetBlockByIDRequest, ...grpc.CallOption) *access.BlockResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetBlockByIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeaderByHeight provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetBlockHeaderByHeight(ctx context.Context, in *access.GetBlockHeaderByHeightRequest, opts ...grpc.CallOption) (*access.BlockHeaderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.BlockHeaderResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetBlockHeaderByHeightRequest, ...grpc.CallOption) *access.BlockHeaderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockHeaderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetBlockHeaderByHeightRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeaderByID provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetBlockHeaderByID(ctx context.Context, in *access.GetBlockHeaderByIDRequest, opts ...grpc.CallOption) (*access.BlockHeaderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.BlockHeaderResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetBlockHeaderByIDRequest, ...grpc.CallOption) *access.BlockHeaderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockHeaderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetBlockHeaderByIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollectionByID provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetCollectionByID(ctx context.Context, in *access.GetCollectionByIDRequest, opts ...grpc.CallOption) (*access.CollectionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.CollectionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetCollectionByIDRequest, ...grpc.CallOption) *access.CollectionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.CollectionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetCollectionByIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsForBlockIDs provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetEventsForBlockIDs(ctx context.Context, in *access.GetEventsForBlockIDsRequest, opts ...grpc.CallOption) (*access.EventsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.EventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetEventsForBlockIDsRequest, ...grpc.CallOption) *access.EventsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.EventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetEventsForBlockIDsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsForHeightRange provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetEventsForHeightRange(ctx context.Context, in *access.GetEventsForHeightRangeRequest, opts ...grpc.CallOption) (*access.EventsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.EventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetEventsForHeightRangeRequest, ...grpc.CallOption) *access.EventsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.EventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetEventsForHeightRangeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlock provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetLatestBlock(ctx context.Context, in *access.GetLatestBlockRequest, opts ...grpc.CallOption) (*access.BlockResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.BlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetLatestBlockRequest, ...grpc.CallOption) *access.BlockResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetLatestBlockRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlockHeader provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetLatestBlockHeader(ctx context.Context, in *access.GetLatestBlockHeaderRequest, opts ...grpc.CallOption) (*access.BlockHeaderResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.BlockHeaderResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetLatestBlockHeaderRequest, ...grpc.CallOption) *access.BlockHeaderResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockHeaderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetLatestBlockHeaderRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestProtocolStateSnapshot provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetLatestProtocolStateSnapshot(ctx context.Context, in *access.GetLatestProtocolStateSnapshotRequest, opts ...grpc.CallOption) (*access.ProtocolStateSnapshotResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.ProtocolStateSnapshotResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetLatestProtocolStateSnapshotRequest, ...grpc.CallOption) *access.ProtocolStateSnapshotResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ProtocolStateSnapshotResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetLatestProtocolStateSnapshotRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetworkParameters provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetNetworkParameters(ctx context.Context, in *access.GetNetworkParametersRequest, opts ...grpc.CallOption) (*access.GetNetworkParametersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.GetNetworkParametersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetNetworkParametersRequest, ...grpc.CallOption) *access.GetNetworkParametersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.GetNetworkParametersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetNetworkParametersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetTransaction(ctx context.Context, in *access.GetTransactionRequest, opts ...grpc.CallOption) (*access.TransactionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.TransactionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetTransactionRequest, ...grpc.CallOption) *access.TransactionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.TransactionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetTransactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionResult provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) GetTransactionResult(ctx context.Context, in *access.GetTransactionRequest, opts ...grpc.CallOption) (*access.TransactionResultResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.TransactionResultResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetTransactionRequest, ...grpc.CallOption) *access.TransactionResultResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.TransactionResultResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetTransactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) Ping(ctx context.Context, in *access.PingRequest, opts ...grpc.CallOption) (*access.PingResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.PingResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.PingRequest, ...grpc.CallOption) *access.PingResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.PingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.PingRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: ctx, in, opts
func (_m *MockRPCClient) SendTransaction(ctx context.Context, in *access.SendTransactionRequest, opts ...grpc.CallOption) (*access.SendTransactionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *access.SendTransactionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.SendTransactionRequest, ...grpc.CallOption) *access.SendTransactionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.SendTransactionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.SendTransactionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
