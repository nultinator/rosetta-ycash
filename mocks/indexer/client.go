// Code generated by mockery v1.0.0. DO NOT EDIT.

package indexer

import (
	context "context"

	bitcoin "github.com/nultinator/rosetta-ycash/ycash"

	mock "github.com/stretchr/testify/mock"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetRawBlock provides a mock function with given fields: _a0, _a1
func (_m *Client) GetRawBlock(_a0 context.Context, _a1 *types.PartialBlockIdentifier) (*bitcoin.Block, []string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *bitcoin.Block
	if rf, ok := ret.Get(0).(func(context.Context, *types.PartialBlockIdentifier) *bitcoin.Block); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bitcoin.Block)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(context.Context, *types.PartialBlockIdentifier) []string); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.PartialBlockIdentifier) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NetworkStatus provides a mock function with given fields: _a0
func (_m *Client) NetworkStatus(_a0 context.Context) (*types.NetworkStatusResponse, error) {
	ret := _m.Called(_a0)

	var r0 *types.NetworkStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context) *types.NetworkStatusResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseBlock provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) ParseBlock(_a0 context.Context, _a1 *bitcoin.Block, _a2 map[string]*types.AccountCoin) (*types.Block, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *bitcoin.Block, map[string]*types.AccountCoin) *types.Block); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *bitcoin.Block, map[string]*types.AccountCoin) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}