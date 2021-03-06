// SPDX-License-Identifier: AGPL-3.0-only

package cfg_test

import (
	"path/filepath"
	"testing"

	"github.com/brainupdaters/drlmctl/cfg"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/test"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func assertCfg(t *testing.T) {
	assert := assert.New(t)

	home, err := homedir.Dir()
	assert.Nil(err)

	assert.Equal("localhost", cfg.Config.Core.Host)
	assert.Equal(50051, cfg.Config.Core.Port)
	assert.Equal(true, cfg.Config.Core.TLS)
	assert.Equal("cert/server.crt", cfg.Config.Core.CertPath)
	assert.Equal("", cfg.Config.Core.Tkn)
	assert.NotNil(cfg.Config.Core.TknExpiration)

	assert.Equal("info", cfg.Config.Log.Level)
	assert.Equal(filepath.Join(home, ".log/drlm/drlmctl.log"), cfg.Config.Log.File)
}

type TestCfgSuite struct {
	test.Test
}

func TestCfg(t *testing.T) {
	suite.Run(t, new(TestCfgSuite))
}

func (s *TestCfgSuite) TestInit() {
	s.Run("should work as expected", func() {
		fs.FS = afero.NewMemMapFs()

		err := afero.WriteFile(fs.FS, "/etc/drlm/drlmctl.toml", nil, 0644)
		s.Nil(err)

		cfg.Init("")

		assertCfg(s.T())
	})

	s.Run("should work as expected with a specified configuration file", func() {
		fs.FS = afero.NewMemMapFs()

		err := afero.WriteFile(fs.FS, "/drlmctl.toml", nil, 0644)
		s.Nil(err)

		cfg.Init("/drlmctl.toml")

		assertCfg(s.T())
	})

	s.Run("should fail and exit if there's an error reading the configuration", func() {
		fs.FS = afero.NewMemMapFs()

		s.Exits(func() { cfg.Init("") })
	})

	s.Run("should fail and exit if there's an error unmarshaling the configuration", func() {
		fs.FS = afero.NewMemMapFs()

		err := afero.WriteFile(fs.FS, "/etc/drlm/drlmctl.json", []byte("invalid config"), 0644)
		s.Nil(err)

		s.Exits(func() { cfg.Init("") })
	})
}
