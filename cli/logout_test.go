// SPDX-License-Identifier: AGPL-3.0-only

package cli_test

import (
	"testing"
	"time"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/cli"
	"github.com/brainupdaters/drlmctl/utils/tests"

	"github.com/stretchr/testify/assert"
)

func TestLogout(t *testing.T) {
	tests.GenerateCfg(t)
	assert := assert.New(t)

	assert.Equal("thisisatoken", cfg.Config.Core.Tkn)
	assert.NotNil(cfg.Config.Core.TknExpiration)

	cli.Logout()

	assert.Equal("", cfg.Config.Core.Tkn)
	assert.Equal(time.Unix(0, 0), cfg.Config.Core.TknExpiration)
}
