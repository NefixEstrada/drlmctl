package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/brainupdaters/drlmctl/cfg"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

// GenerateCfg creates a configuration file with the default values
func GenerateCfg(t *testing.T) {
	assert := assert.New(t)
	fs.FS = afero.NewMemMapFs()

	err := afero.WriteFile(fs.FS, "/etc/drlm/drlmctl.toml", []byte(fmt.Sprintf(`[core]
tkn = "thisisatoken"
tkn_expiration = %s`, time.Now().Format(time.RFC3339Nano))), 0644)
	assert.Nil(err)

	cfg.Init("")
}
