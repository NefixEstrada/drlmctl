package cli_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/cli"
	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/utils/tests"

	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	cmnTests "github.com/brainupdaters/drlm-common/pkg/tests"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestLogin(t *testing.T) {
	assert := assert.New(t)

	t.Run("should work with the username and password passed as parameters", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TLS = false

		now := time.Now()

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserLogin", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API)), &drlm.UserLoginRequest{Usr: "nefix", Pwd: "f0cKt3Rf$"}, []grpc.CallOption(nil)).Return(
			&drlm.UserLoginResponse{
				Tkn:           "thisisatoken",
				TknExpiration: &timestamp.Timestamp{Seconds: now.Unix()},
			}, nil,
		)
		core.Client = theCoreClientMock

		cli.Login("nefix", "f0cKt3Rf$")

		assert.Equal("thisisatoken", cfg.Config.Core.Tkn)
		assert.Equal(time.Unix(now.Unix(), 0), cfg.Config.Core.TknExpiration)
	})

	t.Run("should exit if there's an error during the authentication", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TLS = false
		cfg.Config.Core.Tkn = ""
		cfg.Config.Core.TknExpiration = time.Unix(0, 0)

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserLogin", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API)), &drlm.UserLoginRequest{Usr: "nefix", Pwd: "f0cKt3Rf$"}, []grpc.CallOption(nil)).Return(
			&drlm.UserLoginResponse{}, errors.New("testing error"),
		)
		core.Client = theCoreClientMock

		cmnTests.AssertExits(t, func() { cli.Login("nefix", "f0cKt3Rf$") })

		assert.Equal("", cfg.Config.Core.Tkn)
		assert.NotNil(cfg.Config.Core.TknExpiration)

	})

	t.Run("should exit if there's an error saving the token to the configuration file", func(t *testing.T) {
		cfg.Config = nil

		now := time.Now()

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserLogin", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API)), &drlm.UserLoginRequest{Usr: "nefix", Pwd: "f0cKt3Rf$"}, []grpc.CallOption(nil)).Return(
			&drlm.UserLoginResponse{
				Tkn:           "thisisatoken",
				TknExpiration: &timestamp.Timestamp{Seconds: now.Unix()},
			}, nil,
		)
		core.Client = theCoreClientMock

		cmnTests.AssertExits(t, func() { cli.Login("nefix", "f0cKt3Rf$") })
	})
}
