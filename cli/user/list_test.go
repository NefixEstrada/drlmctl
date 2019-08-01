package user_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/brainupdaters/drlmctl/cli/user"
	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/utils/tests"

	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	assert := assert.New(t)

	t.Run("should show the contents correctly", func(t *testing.T) {
		tests.GenerateCfg(t)

		now := time.Now()

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserList", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API, "tkn", "thisisatoken")), &drlm.UserListRequest{}, []grpc.CallOption(nil)).Return(
			&drlm.UserListResponse{
				Users: []*drlm.UserListResponse_User{
					&drlm.UserListResponse_User{
						Usr:       "nefix",
						AuthType:  drlm.AuthType_LOCAL,
						CreatedAt: &timestamp.Timestamp{Seconds: now.Unix()},
						UpdatedAt: &timestamp.Timestamp{Seconds: now.Unix()},
					},
					&drlm.UserListResponse_User{
						Usr:       "admin",
						AuthType:  drlm.AuthType_LOCAL,
						CreatedAt: &timestamp.Timestamp{Seconds: now.Unix()},
						UpdatedAt: &timestamp.Timestamp{Seconds: now.Unix()},
					},
					&drlm.UserListResponse_User{
						Usr:       "notnefix",
						AuthType:  drlm.AuthType_LOCAL,
						CreatedAt: &timestamp.Timestamp{Seconds: now.Unix()},
						UpdatedAt: &timestamp.Timestamp{Seconds: now.Unix()},
					},
				},
			}, nil,
		)
		core.Client = theCoreClientMock

		var b bytes.Buffer
		user.List(&b)

		assert.Equal(fmt.Sprintf(`┌──────────┬───────────┬─────────────────────┐
│ USERNAME │ AUTH TYPE │ CREATED AT          │
├──────────┼───────────┼─────────────────────┤
│ nefix    │ Local     │ %s │
│ admin    │ Local     │ %s │
│ notnefix │ Local     │ %s │
└──────────┴───────────┴─────────────────────┘
`, now.Format("2006/01/02 15:04:05"), now.Format("2006/01/02 15:04:05"), now.Format("2006/01/02 15:04:05")), b.String())
	})

	t.Run("shouldn't print anything if there's an error", func(t *testing.T) {
		tests.GenerateCfg(t)

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserList", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API, "tkn", "thisisatoken")), &drlm.UserListRequest{}, []grpc.CallOption(nil)).Return(&drlm.UserListResponse{}, errors.New("testing error"))
		core.Client = theCoreClientMock

		var b bytes.Buffer
		user.List(&b)

		assert.Equal("", b.String())
	})
}
