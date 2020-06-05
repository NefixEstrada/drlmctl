// SPDX-License-Identifier: AGPL-3.0-only

package core

import (
	"os/user"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/software"

	"github.com/brainupdaters/drlm-common/pkg/os"
	"github.com/brainupdaters/drlm-common/pkg/os/client"
	"github.com/brainupdaters/drlm-common/pkg/ssh"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// Install compiles / downloads the DRLM Core binary, installs it in the DRLM Core server and starts it
func Install(fs afero.Fs, v string, port int, usr, pwd string) {
	// Compile DRLM Core
	bin, err := software.SoftwareCore.Compile(cfg.Config.Core.OS, cfg.Config.Core.Arch, v)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	ctlCli := &client.Local{}
	ctlOS, err := os.DetectOS(ctlCli)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	// Configure Core
	// cfg.Config.Core.SSHPort = port
	// cfg.Config.Core.SSHKeys, err = ctlOS.CmdSSHGetHostKeys(ctlCli, cfg.Config.Core.Host, cfg.Config.Core.SSHPort)
	// if err != nil {
	// 	log.Fatalf("error adding the core server: %v", err)
	// }

	// s, err := ssh.NewSessionWithPassword(cfg.Config.Core.Host, cfg.Config.Core.SSHPort, usr, pwd, cfg.Config.Core.SSHKeys)
	// if err != nil {
	// 	log.Fatalf("error adding the core server: %v", err)
	// }
	// defer s.Close()

	// coreCli := &client.SSH{
	// 	Session: s,
	// }

	// cfg.Config.Core.OS, err = os.DetectOS(coreCli)
	// if err != nil {
	// 	log.Fatalf("error adding the core server: %v", err)
	// }

	// cfg.Config.Core.Arch, err = os.DetectArch(coreCli)
	// if err != nil {
	// 	log.Fatalf("error adding the core server: %v", err)
	// }

	// if err = cfg.Config.Save(); err != nil {
	// 	log.Fatalf("error adding the core server: %v", err)
	// }

	// // Create the DRLM user
	// if err = cfg.Config.Core.OS.CmdUserCreate(coreCli, "drlm", "changeme"); err != nil {
	// 	log.Fatalf("error adding the core server: %v", err)
	// }

	// // Copy SSH keys
	u, err := user.Current()
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	keysPath, err := ctlOS.CmdSSHGetKeysPath(ctlCli, u.Username)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	s, err := ssh.NewSessionWithKey(fs, cfg.Config.Core.Host, cfg.Config.Core.SSHPort, "drlm", keysPath, cfg.Config.Core.SSHKeys)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}
	defer s.Close()

	coreCli := &client.SSH{
		Session: s,
	}

	b, err := afero.ReadFile(fs, bin)
	if err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}

	if err := cfg.Config.Core.OS.CmdPkgInstallBinary(coreCli, "drlm", "drlm-core", b); err != nil {
		log.Fatalf("error installing the DRLM Core binary: %v", err)
	}
}
