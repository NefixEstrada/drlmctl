// SPDX-License-Identifier: AGPL-3.0-only

package software

import (
	"fmt"
	"os/exec"
	"path/filepath"

	stdOS "os"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/os"
	"github.com/spf13/afero"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// software is a DRLM component that can be either downloaded or compiled
type software struct {
	GitService gitService
	User       string
	Repo       string
	Out        string
}

var (
	// SoftwareCore is DRLM Core
	SoftwareCore = &software{
		GitService: gitServiceGitHub,
		User:       "brainupdaters",
		Repo:       "drlm-core",
		Out:        "drlm-core",
	}
	// SoftwareAgent is DRLM Agent
	SoftwareAgent = &software{
		GitService: gitServiceGitHub,
		User:       "brainupdaters",
		Repo:       "drlm-agent",
		Out:        "drlm-agent",
	}
)

// Compile compiles a software piece
func (s *software) Compile(o os.OS, arch os.Arch, v string) (string, error) {
	d, err := afero.TempDir(fs.FS, "", "drlmctl-compile-")
	if err != nil {
		return "", fmt.Errorf("error creating a temporary directory: %v", err)
	}

	r, err := git.PlainClone(d, false, &git.CloneOptions{
		URL:      s.GitService.CloneURL(s.User, s.Repo),
		Progress: stdOS.Stdout,
	})
	if err != nil {
		return "", fmt.Errorf("error clonning the repository: %v", err)
	}

	if v == "" {
		v, err = s.GitService.GetRelease(s.User, s.Repo, "latest")
		if err != nil {
			return "", fmt.Errorf("error getting the latest version: %v", err)
		}
	}

	w, err := r.Worktree()
	if err != nil {
		return "", fmt.Errorf("error getting the repository worktree: %v", err)
	}

	if err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewTagReferenceName(v),
	}); err != nil {
		return "", fmt.Errorf("error checking out to the version %s: %v", v, err)
	}

	if err := stdOS.Chdir(d); err != nil {
		return "", fmt.Errorf("error changing the working directory to the repository: %v", err)
	}

	cmd := exec.Command("make", "build")
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("error compiling the binary: %v: %s", err, out)
	}

	return filepath.Join(d, s.Out), nil
}
