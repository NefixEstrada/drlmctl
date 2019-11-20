// SPDX-License-Identifier: AGPL-3.0-only

package core

import (
	"os/user"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/software"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/os"
	"github.com/brainupdaters/drlm-common/pkg/os/client"
	"github.com/brainupdaters/drlm-common/pkg/ssh"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// Install compiles / downloads the DRLM Core binary, installs it in the DRLM Core server and starts it
func Install(v string) {
	bin, err := software.SoftwareCore.Compile(cfg.Config.Core.OS, cfg.Config.Core.Arch, v)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	ctlCli := &client.Local{}
	ctlOS, err := os.DetectOS(ctlCli)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	u, err := user.Current()
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	keysPath, err := ctlOS.CmdSSHGetKeysPath(ctlCli, u.Username)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	s, err := ssh.NewSessionWithKey(cfg.Config.Core.Host, cfg.Config.Core.SSHPort, "drlm", keysPath, cfg.Config.Core.SSHKeys)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}
	defer s.Close()

	coreCli := &client.SSH{
		Session: s,
	}

	b, err := afero.ReadFile(fs.FS, bin)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	if err := cfg.Config.Core.OS.CmdPkgInstallBinary(coreCli, "drlm", "drlm-core", b); err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}
}
