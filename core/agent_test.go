// SPDX-License-Identifier: AGPL-3.0-only

package core_test

import (
	"context"
	"errors"
	"testing"

	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/models"
	"github.com/brainupdaters/drlmctl/utils/tests"

	"github.com/brainupdaters/drlm-common/pkg/os"
	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type TestAgentSuite struct {
	suite.Suite
	cli *drlm.DRLMClient
}

func TestAgent(t *testing.T) {
	suite.Run(t, new(TestAgentSuite))
}

func (s *TestAgentSuite) TestList() {
	s.Run("should return the list of agents correctly", func() {
		tests.GenerateCfg(s.T())

		mock := &tests.CoreClientMock{}
		mock.On("AgentList", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API, "tkn", "thisisatoken")), &drlm.AgentListRequest{}, []grpc.CallOption(nil)).Return(&drlm.AgentListResponse{
			Agents: []*drlm.AgentListResponse_Agent{
				{
					Host:          "192.168.122.100",
					Port:          22,
					User:          "drlm",
					Version:       "v1.0.0",
					Arch:          drlm.Arch_ARCH_AMD64,
					Os:            drlm.OS_OS_LINUX,
					OsVersion:     "5.3.5",
					Distro:        "NixOS",
					DistroVersion: "20.03",
				},
				{
					Host:          "192.168.122.101",
					Port:          22,
					User:          "drlm",
					Version:       "v1.0.0",
					Arch:          drlm.Arch_ARCH_AMD64,
					Os:            drlm.OS_OS_LINUX,
					OsVersion:     "4.19.2",
					Distro:        "Debian",
					DistroVersion: "9",
				},
			},
		}, nil)
		core.Client = mock

		rsp, err := core.AgentList()

		s.Nil(err)
		s.Equal([]*models.Agent{
			{
				Host: "192.168.122.100",
				Port: 22,
				Usr:  "drlm",
				OS:   os.Linux,
				Arch: os.ArchAmd64,
			},
			{
				Host: "192.168.122.101",
				Port: 22,
				Usr:  "drlm",
				OS:   os.Linux,
				Arch: os.ArchAmd64,
			},
		}, rsp)
	})

	s.Run("should return an error if there's an error getting the agents list", func() {
		tests.GenerateCfg(s.T())

		mock := &tests.CoreClientMock{}
		mock.On("AgentList", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", core.API, "tkn", "thisisatoken")), &drlm.AgentListRequest{}, []grpc.CallOption(nil)).Return(&drlm.AgentListResponse{}, errors.New("testing error"))
		core.Client = mock

		rsp, err := core.AgentList()

		s.EqualError(err, "error listing the DRLM Agents: testing error")
		s.Equal([]*models.Agent{}, rsp)
	})
}

func (s *TestAgentSuite) TestGet() {
	s.Run("should return the agent correctly", func() {

	})
}
