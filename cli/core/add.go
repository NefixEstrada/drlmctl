package core

import (
	"os/user"
	"path/filepath"

	"github.com/brainupdaters/drlmctl/cfg"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/os"
	"github.com/brainupdaters/drlm-common/pkg/os/client"
	"github.com/brainupdaters/drlm-common/pkg/ssh"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// Add connects to the Core host, creates the drlm user and copies the keys to that user, which has to be admin
func Add(host string, port int, usr, pwd string) {
	ctlCli := &client.Local{}
	ctlOS, err := os.DetectOS(ctlCli)
	if err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	cfg.Config.Core.Host = host
	cfg.Config.Core.SSHPort = port
	cfg.Config.Core.SSHKeys, err = ctlOS.CmdSSHGetHostKeys(ctlCli, cfg.Config.Core.Host, cfg.Config.Core.SSHPort)
	if err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	s, err := ssh.NewSessionWithPassword(cfg.Config.Core.Host, cfg.Config.Core.SSHPort, usr, pwd, cfg.Config.Core.SSHKeys)
	if err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}
	defer s.Close()

	coreCli := &client.SSH{
		Session: s,
	}

	cfg.Config.Core.OS, err = os.DetectOS(coreCli)
	if err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	cfg.Config.Core.Arch, err = os.DetectArch(coreCli)
	if err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	if err = cfg.Config.Save(); err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	// Create the DRLM user
	if err = cfg.Config.Core.OS.CmdUserCreate(coreCli, "drlm", "changeme"); err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	u, err := user.Current()
	if err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	keysPath, err := ctlOS.CmdSSHGetKeysPath(ctlCli, u.Username)
	if err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	pubKey, err := afero.ReadFile(fs.FS, filepath.Join(keysPath, "id_rsa.pub"))
	if err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	// Copy the SSH public key
	if err = cfg.Config.Core.OS.CmdSSHCopyID(coreCli, "drlm", pubKey); err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	// Disable the DRLM agent user
	if err = cfg.Config.Core.OS.CmdUserDisable(coreCli, "drlm"); err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	// Make the user administrator
	if err = cfg.Config.Core.OS.CmdUserMakeAdmin(coreCli, "drlm"); err != nil {
		log.Fatalf("error adding the core server: %v", err)
	}

	coreKeysPath, err := cfg.Config.Core.OS.CmdSSHGetKeysPath(coreCli, "drlm")
	if err != nil {
		log.Fatalf("error getting the DRLM Core SSH keys path: %v", err)
	}

	if err = cfg.Config.Core.OS.CmdSSHGenerateKeyPair(coreCli, coreKeysPath); err != nil {
		log.Fatalf("error generating the SSH key pair: %v", err)
	}
}
