// SPDX-License-Identifier: AGPL-3.0-only

package core

import (
	"github.com/brainupdaters/drlmctl/cfg"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// Add connects to the Core host, creates the drlm user and copies the keys to that user, which has to be admin
func Add(fs afero.Fs, host string) {
	cfg.Config.Core.Host = host
	if err := cfg.Config.Save(); err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}
}
