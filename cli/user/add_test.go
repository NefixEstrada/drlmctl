// SPDX-License-Identifier: AGPL-3.0-only

package user_test

import (
	"context"
	"testing"

	"github.com/brainupdaters/drlmctl/cli/user"
	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/models"
	"github.com/brainupdaters/drlmctl/utils/tests"

	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestUserAdd(t *testing.T) {
	t.Run("should work as expected", func(t *testing.T) {
		tests.GenerateCfg(t)

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserAdd", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API, "tkn", "thisisatoken")), &drlm.UserAddRequest{Usr: "nefix", Pwd: "f0cKt3Rf$"}, []grpc.CallOption(nil)).Return(
			&drlm.UserAddResponse{}, nil,
		)
		core.Client = theCoreClientMock

		usr := &models.User{
			Username: "nefix",
			Password: "f0cKt3Rf$",
		}

		user.Add(usr)
	})

}
