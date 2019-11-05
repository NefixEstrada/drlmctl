package tests

import (
	"context"
	"errors"

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

// AgentAdd mocks the AgentAdd  gRPC method
func (c *CoreClientMock) AgentAdd(ctx context.Context, req *drlm.AgentAddRequest, opts ...grpc.CallOption) (*drlm.AgentAddResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.AgentAddResponse), args.Error(1)
}

// AgentInstall mocks the AgentInstall  gRPC method
// AgentInstall(ctx context.Context, opts ...grpc.CallOption) (DRLM_AgentInstallClient, error)
func (c *CoreClientMock) AgentInstall(ctx context.Context, opts ...grpc.CallOption) (drlm.DRLM_AgentInstallClient, error) {
	return nil, errors.New("mock not implemented")
}

// AgentDelete mocks the AgentDelete  gRPC method
func (c *CoreClientMock) AgentDelete(ctx context.Context, req *drlm.AgentDeleteRequest, opts ...grpc.CallOption) (*drlm.AgentDeleteResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.AgentDeleteResponse), args.Error(1)
}

// AgentList mocks the AgentList  gRPC method
func (c *CoreClientMock) AgentList(ctx context.Context, req *drlm.AgentListRequest, opts ...grpc.CallOption) (*drlm.AgentListResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.AgentListResponse), args.Error(1)
}

// AgentGet mocks the AgentGet  gRPC method
func (c *CoreClientMock) AgentGet(ctx context.Context, req *drlm.AgentGetRequest, opts ...grpc.CallOption) (*drlm.AgentGetResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.AgentGetResponse), args.Error(1)
}

// JobSchedule mocks the JobSchedule  gRPC method
func (c *CoreClientMock) JobSchedule(ctx context.Context, req *drlm.JobScheduleRequest, opts ...grpc.CallOption) (*drlm.JobScheduleResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.JobScheduleResponse), args.Error(1)
}

// JobCancel mocks the JobCancel  gRPC method
func (c *CoreClientMock) JobCancel(ctx context.Context, req *drlm.JobCancelRequest, opts ...grpc.CallOption) (*drlm.JobCancelResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.JobCancelResponse), args.Error(1)
}

// JobList mocks the JobList  gRPC method
func (c *CoreClientMock) JobList(ctx context.Context, req *drlm.JobListRequest, opts ...grpc.CallOption) (*drlm.JobListResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.JobListResponse), args.Error(1)
}

// JobNotify mocks the JobNotify  gRPC method
func (c *CoreClientMock) JobNotify(ctx context.Context, req *drlm.JobNotifyRequest, opts ...grpc.CallOption) (*drlm.JobNotifyResponse, error) {
	args := c.Called(ctx, req, opts)
	return args.Get(0).(*drlm.JobNotifyResponse), args.Error(1)
}
