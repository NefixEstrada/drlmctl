package core_test

import (
	"path/filepath"
	"testing"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/utils/tests"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	cmnTests "github.com/brainupdaters/drlm-common/pkg/tests"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	t.Run("should work as expected with TLS", func(t *testing.T) {
		tests.GenerateCfg(t)
		cmnTests.GenerateCert(t, fs.FS, "server", filepath.Dir(cfg.Config.Core.CertPath))

		core.Init()
		assert.NotEqual(&grpc.ClientConn{}, core.Conn)
	})

	t.Run("should work as expected without TLS", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TLS = false

		core.Init()
		assert.NotEqual(&grpc.ClientConn{}, core.Conn)
	})

	t.Run("should return an error if there's an error loading the TLS certificate", func(t *testing.T) {
		tests.GenerateCfg(t)
		cmnTests.AssertExits(t, core.Init)
	})
}
