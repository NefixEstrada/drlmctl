package core

import (
	"os/user"
	"path/filepath"

	"github.com/brainupdaters/drlmctl/cfg"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/os"
	"github.com/brainupdaters/drlm-common/pkg/ssh"
	"github.com/spf13/afero"
)

// Add connects to the Core host, creates the drlm user and copies the keys to that user, which has to be admin
func Add(host string, port int, usr, pwd string, isAdmin bool) {
	ctlCli := &os.ClientLocal{}
	ctlOS, err := os.DetectOS(ctlCli)
	if err != nil {
		// PANIC
		panic(err)
	}

	cfg.Config.Core.Host = host
	cfg.Config.Core.SSHPort = port
	cfg.Config.Core.SSHKeys, err = ctlOS.CmdSSHGetHostKeys(ctlCli, cfg.Config.Core.Host, cfg.Config.Core.SSHPort)
	if err != nil {
		// PANIC
		panic(err)
	}

	s, err := ssh.NewSessionWithPassword(cfg.Config.Core.Host, cfg.Config.Core.SSHPort, usr, pwd, cfg.Config.Core.SSHKeys)
	if err != nil {
		// PANIC
		panic(err)
	}
	defer s.Close()

	coreCli := &os.ClientSSH{
		Session: s,
		IsAdmin: isAdmin,
	}

	cfg.Config.Core.OS, err = os.DetectOS(coreCli)
	if err != nil {
		// PANIC
		panic(err)
	}

	cfg.Config.Core.Arch, err = os.DetectArch(coreCli)
	if err != nil {
		// PANIC
		panic(err)
	}

	if err = cfg.Config.Save(); err != nil {
		// PANIC
		panic(err)
	}

	// Create the DRLM user
	if err = cfg.Config.Core.OS.CmdUserCreate(coreCli, "drlm", "changeme"); err != nil {
		// PANIC
		panic(err)
	}

	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	keysPath, err := ctlOS.CmdSSHGetKeysPath(ctlCli, u.Username)
	if err != nil {
		panic(err)
	}

	pubKey, err := afero.ReadFile(fs.FS, filepath.Join(keysPath, "id_rsa.pub"))
	if err != nil {
		// PANIC
		panic(err)
	}

	// Copy the SSH public key
	if err = cfg.Config.Core.OS.CmdSSHCopyID(coreCli, "drlm", pubKey); err != nil {
		// PANIC
		panic(err)
	}

	// Disable the DRLM agent user
	if err = cfg.Config.Core.OS.CmdUserDisable(coreCli, "drlm"); err != nil {
		// PANIC
		panic(err)
	}

	// Make the user administrator
	if err = cfg.Config.Core.OS.CmdUserMakeAdmin(coreCli, "drlm"); err != nil {
		// PANIC
		panic(err)
	}
}
