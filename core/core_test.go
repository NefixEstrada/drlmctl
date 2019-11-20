// SPDX-License-Identifier: AGPL-3.0-only

package core_test

import (
	"path/filepath"
	"testing"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/core"

	"github.com/brainupdaters/drlm-common/pkg/test"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type TestCoreSuite struct {
	test.Test
}

func TestCore(t *testing.T) {
	suite.Run(t, new(TestCoreSuite))
}

func (s *TestCoreSuite) TestInit() {
	s.Run("should work as expected with TLS", func() {
		s.GenerateCfg(test.CfgPathCtl, func() { cfg.Init("") })
		s.GenerateCert("server", filepath.Dir(cfg.Config.Core.CertPath))

		core.Init()
		s.NotEqual(&grpc.ClientConn{}, core.Conn)
	})

	s.Run("should work as expected without TLS", func() {
		s.GenerateCfg(test.CfgPathCtl, func() { cfg.Init("") })
		cfg.Config.Core.TLS = false

		core.Init()
		s.NotEqual(&grpc.ClientConn{}, core.Conn)
	})

	s.Run("should exit if there's an error loading the TLS certificate", func() {
		s.GenerateCfg(test.CfgPathCtl, func() { cfg.Init("") })

		s.Exits(core.Init)
	})
}
