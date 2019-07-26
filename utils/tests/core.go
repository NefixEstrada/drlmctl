package tests

import (
	"context"

	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// CoreClientMock is a mock for the DRLM core gRPC client
type CoreClientMock struct {
	mock.Mock
}

// UserLogin mocks the UserLogin gRPC method
func (c *CoreClientMock) UserLogin(ctx context.Context, req *drlm.UserLoginRequest, opts ...grpc.CallOption) (*drlm.UserLoginResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.UserLoginResponse), args.Error(1)
}

// UserTokenRenew mocks the UserTokenRenew gRPC method
func (c *CoreClientMock) UserTokenRenew(ctx context.Context, req *drlm.UserTokenRenewRequest, opts ...grpc.CallOption) (*drlm.UserTokenRenewResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.UserTokenRenewResponse), args.Error(1)
}

// UserAdd mocks the UserAdd gRPC method
func (c *CoreClientMock) UserAdd(ctx context.Context, req *drlm.UserAddRequest, opts ...grpc.CallOption) (*drlm.UserAddResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.UserAddResponse), args.Error(1)
}

// UserDelete mocks the UserDelete gRPC method
func (c *CoreClientMock) UserDelete(ctx context.Context, req *drlm.UserDeleteRequest, opts ...grpc.CallOption) (*drlm.UserDeleteResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.UserDeleteResponse), args.Error(1)
}

// UserList mocks the UserList gRPC method
func (c *CoreClientMock) UserList(ctx context.Context, req *drlm.UserListRequest, opts ...grpc.CallOption) (*drlm.UserListResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.UserListResponse), args.Error(1)
}
