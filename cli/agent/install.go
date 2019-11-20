// SPDX-License-Identifier: AGPL-3.0-only

package agent

import (
	"os"

	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/software"

	log "github.com/sirupsen/logrus"
)

// Install compiles / downloads the DRLM Agent binary, installs it in the DRLM Agent server and starts it
func Install(host string, v string) {
	a, err := core.AgentGet(host)
	if err != nil {
		os.Exit(1)
	}

	bin, err := software.SoftwareAgent.Compile(a.OS, a.Arch, v)
	if err != nil {
		log.Fatalf("error installing the DRLM Agent binary: %v", err)
	}

	core.AgentInstall(host, bin)
	if err != nil {
		os.Exit(1)
	}
}
