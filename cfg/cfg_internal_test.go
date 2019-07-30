package cfg

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSetDefaults(t *testing.T) {
	assert := assert.New(t)

	t.Run("should work as expected", func(t *testing.T) {
		home, err := homedir.Dir()
		assert.Nil(err)

		v = viper.New()

		SetDefaults()

		assert.Equal("localhost", v.GetString("core.host"))
		assert.Equal(50051, v.GetInt("core.port"))
		assert.Equal(true, v.GetBool("core.tls"))
		assert.Equal("cert/server.crt", v.GetString("core.cert_path"))
		assert.Equal("", v.GetString("core.tkn"))
		assert.Equal(time.Time{}, v.GetTime("core.tkn_expiration"))

		assert.Equal("info", v.GetString("log.level"))
		assert.Equal(filepath.Join(home, ".log/drlm/drlmctl.log"), v.GetString("log.file"))
	})
}

func TestSaveTkn(t *testing.T) {
	assert := assert.New(t)

	t.Run("should save the configuration correctly", func(t *testing.T) {
		fs.FS = afero.NewMemMapFs()

		err := afero.WriteFile(fs.FS, "/etc/drlm/drlmctl.toml", nil, 0644)
		assert.Nil(err)

		Init("")

		now := time.Now()
		assert.Nil(SaveTkn("imaginethisisatoken", now))

		b, err := afero.ReadFile(fs.FS, "/etc/drlm/drlmctl.toml")
		assert.Nil(err)

		home, err := homedir.Dir()
		assert.Nil(err)

		assert.Equal(fmt.Sprintf(`
[core]
  cert_path = "cert/server.crt"
  host = "localhost"
  port = 50051
  tkn = "imaginethisisatoken"
  tkn_expiration = %s
  tls = true

[log]
  file = "%s"
  level = "info"
`, now.Format(time.RFC3339), filepath.Join(home, ".log/drlm/drlmctl.log")), string(b))
	})

	t.Run("should return an error if there's an error saving the configuration", func(t *testing.T) {
		fs.FS = afero.NewMemMapFs()

		err := afero.WriteFile(fs.FS, "/etc/drlm/drlmctl.toml", nil, 0644)
		assert.Nil(err)

		Init("")

		fs.FS = afero.NewReadOnlyFs(fs.FS)
		v.SetFs(fs.FS)

		now := time.Now()
		assert.EqualError(SaveTkn("imaginethisisatoken", now), "operation not permitted")
	})
}
