// SPDX-License-Identifier: AGPL-3.0-only

package software

import (
	"fmt"
	stdOS "os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/models"

	"github.com/blang/semver"
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

/*

 TODO: OS AND ARCH!

*/

// CompilePlugin compiles a plugin
func CompilePlugin(fs afero.Fs, p *models.Plugin, os os.OS, arch os.Arch, v string) (string, error) {
	repo, ok := cfg.Config.PluginRepos[p.Repo]
	if !ok {
		return "", fmt.Errorf("plugin repository not found")
	}

	d, err := afero.TempDir(fs, "", "drlmctl-compile-")
	if err != nil {
		return "", fmt.Errorf("error creating a temporary directory: %v", err)
	}

	r, err := git.PlainClone(d, false, &git.CloneOptions{
		URL:      repo.URL,
		Progress: stdOS.Stdout,
	})
	if err != nil {
		return "", fmt.Errorf("error clonning the repository: %v", err)
	}

	if v != "" {
		if p.Version != v {
			tags, err := r.Tags()
			if err != nil {
				panic(err)
			}

			var tag *plumbing.Reference
			if err := tags.ForEach(func(t *plumbing.Reference) error {
				if strings.HasPrefix(t.Name().Short(), p.Name) {
					tParts := strings.Split(tag.Name().Short(), "-")
					if len(tParts) == 2 {
						ver, err := semver.ParseTolerant(tParts[1])
						if err != nil {
							return nil
						}

						pVer, err := semver.ParseTolerant(p.Version)
						if err != nil {
							return fmt.Errorf("invalid plugin version")
						}

						if ver.Equals(pVer) {
							tag = t
						}

						return nil
					}

					return nil
				}

				return nil
			}); err != nil {
				return "", err
			}

			if tag != nil {
				w, err := r.Worktree()
				if err != nil {
					return "", fmt.Errorf("error getting the repository worktree: %v", err)
				}

				if err := w.Checkout(&git.CheckoutOptions{
					Branch: tag.Name(),
				}); err != nil {
					return "", fmt.Errorf("error checking out to the version %s: %v", v, err)
				}
			}
		}
	} else {
		head, err := r.Head()
		if err != nil {
			return "", fmt.Errorf("error getting the head of the repository: %v", err)
		}

		p.Version = head.Hash().String()
	}

	if err := stdOS.Chdir(filepath.Join(d, p.Name)); err != nil {
		return "", fmt.Errorf("error changing the working directory to the repository: %v", err)
	}

	cmd := exec.Command("make", "build")
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("error compiling the binary: %v: %s", err, out)
	}

	return filepath.Join(d, p.Name, p.Name), nil
}

// Compile compiles a software piece
func (s *software) Compile(o os.OS, arch os.Arch, v string) (string, error) {
	d := filepath.Join("/src", s.Repo)
	// d, err := afero.TempDir(fs.FS, "", "drlmctl-compile-")
	// if err != nil {
	// 	return "", fmt.Errorf("error creating a temporary directory: %v", err)
	// }
	// defer fs.FS.RemoveAll(d)

	// r, err := git.PlainClone(d, false, &git.CloneOptions{
	// 	URL:      s.GitService.CloneURL(s.User, s.Repo),
	// 	Progress: stdOS.Stdout,
	// })
	// if err != nil {
	// 	return "", fmt.Errorf("error clonning the repository: %v", err)
	// }

	// if v == "" {
	// 	v, err = s.GitService.GetRelease(s.User, s.Repo, "latest")
	// 	if err != nil {
	// 		return "", fmt.Errorf("error getting the latest version: %v", err)
	// 	}
	// }

	// w, err := r.Worktree()
	// if err != nil {
	// 	return "", fmt.Errorf("error getting the repository worktree: %v", err)
	// }

	// if err = w.Checkout(&git.CheckoutOptions{
	// 	Branch: plumbing.NewTagReferenceName(v),
	// }); err != nil {
	// 	return "", fmt.Errorf("error checking out to the version %s: %v", v, err)
	// }

	if err := stdOS.Chdir(d); err != nil {
		return "", fmt.Errorf("error changing the working directory to the repository: %v", err)
	}

	cmd := exec.Command("make", "build")
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("error compiling the binary: %v: %s", err, out)
	}

	return filepath.Join(d, s.Out), nil
}
