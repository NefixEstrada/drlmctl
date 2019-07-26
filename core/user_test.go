package core_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/models"
	"github.com/brainupdaters/drlmctl/utils/tests"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestUserLogin(t *testing.T) {
	assert := assert.New(t)

	t.Run("should work as expected", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TLS = false

		expiration := time.Now().Add(5 * time.Minute).Unix()

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserLogin", context.Background(), &drlm.UserLoginRequest{Usr: "nefix", Pwd: "f0cKt3Rf$"}, []grpc.CallOption(nil)).Return(
			&drlm.UserLoginResponse{
				Tkn:           "thisisatoken",
				TknExpiration: &timestamp.Timestamp{Seconds: expiration},
			}, nil,
		)
		core.Client = theCoreClientMock

		rsp, err := core.UserLogin("nefix", "f0cKt3Rf$")
		assert.Nil(err)
		assert.Equal(&drlm.UserLoginResponse{
			Tkn:           "thisisatoken",
			TknExpiration: &timestamp.Timestamp{Seconds: expiration},
		}, rsp)
	})

	t.Run("should return an error if there's an error logging in", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TLS = false

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserLogin", context.Background(), &drlm.UserLoginRequest{Usr: "nefix", Pwd: "f0cKt3Rf$"}, []grpc.CallOption(nil)).Return(
			&drlm.UserLoginResponse{}, errors.New("testing error"),
		)
		core.Client = theCoreClientMock

		rsp, err := core.UserLogin("nefix", "f0cKt3Rf$")
		assert.EqualError(err, "error logging into DRLM Core: testing error")
		assert.Equal(&drlm.UserLoginResponse{}, rsp)
	})
}

func TestUserTokenRenew(t *testing.T) {
	assert := assert.New(t)

	t.Run("should work as expected", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TLS = false

		expiration := time.Now().Add(5 * time.Minute).Unix()

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserTokenRenew", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API, "tkn", "thisisatoken")), &drlm.UserTokenRenewRequest{}, []grpc.CallOption(nil)).Return(
			&drlm.UserTokenRenewResponse{
				Tkn:           "thisisanewtoken",
				TknExpiration: &timestamp.Timestamp{Seconds: expiration},
			}, nil,
		)
		core.Client = theCoreClientMock

		rsp, err := core.UserTokenRenew()
		assert.Nil(err)
		assert.Equal(&drlm.UserTokenRenewResponse{
			Tkn:           "thisisanewtoken",
			TknExpiration: &timestamp.Timestamp{Seconds: expiration},
		}, rsp)
	})

	t.Run("should return an error if there's an error renewing the token", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TLS = false

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserTokenRenew", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API, "tkn", "thisisatoken")), &drlm.UserTokenRenewRequest{}, []grpc.CallOption(nil)).Return(
			&drlm.UserTokenRenewResponse{}, errors.New("testing error"),
		)
		core.Client = theCoreClientMock

		rsp, err := core.UserTokenRenew()
		assert.EqualError(err, "error renewing the user token: testing error")
		assert.Equal(&drlm.UserTokenRenewResponse{}, rsp)
	})

	t.Run("should not panic if there's no token in the configuration", func(t *testing.T) {
		fs.FS = afero.NewMemMapFs()

		err := afero.WriteFile(fs.FS, "/etc/drlm/drlmctl.toml", nil, 0644)
		assert.Nil(err)

		cfg.Init("")
		cfg.Config.Core.TLS = false

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserTokenRenew", metadata.AppendToOutgoingContext(context.Background(), "api", core.API, "tkn", ""), &drlm.UserTokenRenewRequest{}, []grpc.CallOption(nil)).Return(
			&drlm.UserTokenRenewResponse{}, errors.New("testing error"),
		)
		core.Client = theCoreClientMock

		rsp, err := core.UserTokenRenew()
		assert.EqualError(err, "error renewing the user token: testing error")
		assert.Equal(&drlm.UserTokenRenewResponse{}, rsp)
	})
}

func TestUserAdd(t *testing.T) {
	assert := assert.New(t)

	t.Run("should work as expected", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TLS = false

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserAdd", metadata.AppendToOutgoingContext(context.Background(), "api", core.API, "tkn", "thisisatoken"), &drlm.UserAddRequest{Usr: "nefix", Pwd: "f0cKt3Rf$"}, []grpc.CallOption(nil)).Return(
			&drlm.UserAddResponse{}, nil,
		)
		core.Client = theCoreClientMock

		usr := &models.User{
			Username: "nefix",
			Password: "f0cKt3Rf$",
		}

		err := core.UserAdd(usr)
		assert.Nil(err)
	})

	t.Run("should return if there's an error adding the user", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TLS = false

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserAdd", metadata.AppendToOutgoingContext(context.Background(), "api", core.API, "tkn", "thisisatoken"), &drlm.UserAddRequest{Usr: "nefix", Pwd: "f0cKt3Rf$"}, []grpc.CallOption(nil)).Return(
			&drlm.UserAddResponse{}, errors.New("testing error"),
		)
		core.Client = theCoreClientMock

		usr := &models.User{
			Username: "nefix",
			Password: "f0cKt3Rf$",
		}

		err := core.UserAdd(usr)
		assert.EqualError(err, "error adding the user to DRLM Core: testing error")
	})
}
