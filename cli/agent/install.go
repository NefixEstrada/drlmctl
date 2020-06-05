// SPDX-License-Identifier: AGPL-3.0-only

package agent

import (
	"os"

	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/software"
	"github.com/spf13/afero"

	log "github.com/sirupsen/logrus"
)

// Install compiles / downloads the DRLM Agent binary, installs it in the DRLM Agent server and starts it
func Install(fs afero.Fs, host string, port int, usr, pwd, v string) {
	a, err := core.AgentGet(host)
	if err != nil {
		os.Exit(1)
	}

	a.SSHPort = int32(port)
	a.SSHUsr = usr
	a.SSHPwd = pwd

	bin, err := software.SoftwareAgent.Compile(a.OS, a.Arch, v)
	if err != nil {
		log.Fatalf("error installing the DRLM Agent binary: %v", err)
	}

	core.AgentInstall(fs, a, bin)
	if err != nil {
		os.Exit(1)
	}
}
