// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "akila/x/feemarket/types"
)

// FeeMarketQueryClient is an autogenerated mock type for the QueryClient type
type FeeMarketQueryClient struct {
	mock.Mock
}

// BaseFee provides a mock function with given fields: ctx, in, opts
func (_m *FeeMarketQueryClient) BaseFee(ctx context.Context, in *types.QueryBaseFeeRequest, opts ...grpc.CallOption) (*types.QueryBaseFeeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryBaseFeeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryBaseFeeRequest, ...grpc.CallOption) *types.QueryBaseFeeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryBaseFeeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryBaseFeeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockGas provides a mock function with given fields: ctx, in, opts
func (_m *FeeMarketQueryClient) BlockGas(ctx context.Context, in *types.QueryBlockGasRequest, opts ...grpc.CallOption) (*types.QueryBlockGasResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryBlockGasResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryBlockGasRequest, ...grpc.CallOption) *types.QueryBlockGasResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryBlockGasResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryBlockGasRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *FeeMarketQueryClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryParamsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) *types.QueryParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryParamsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewQueryClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueryClient creates a new instance of QueryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeeMarketQueryClient(t mockConstructorTestingTNewQueryClient) *FeeMarketQueryClient {
	mock := &FeeMarketQueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
