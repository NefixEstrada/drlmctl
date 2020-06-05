// SPDX-License-Identifier: AGPL-3.0-only

package plugin

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/models"

	"github.com/blang/semver"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// List returns a list with all the available plugins of a repo. If no repos are provided,
// it will list all the plugins from all the repositories
func List(fs afero.Fs, repos ...string) ([]*models.Plugin, error) {
	d, err := afero.TempDir(fs, "", "drlmctl-plugins-list-")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Errorf("error creating a temporary directory")
		return []*models.Plugin{}, fmt.Errorf("error creating a temporary directory: %v", err)
	}
	defer fs.RemoveAll(d)

	plugins := []*models.Plugin{}
	repoFound := false
	for n, v := range cfg.Config.PluginRepos {
		if len(repos) == 0 || repos[0] == n {
			repoFound = true
			var auth transport.AuthMethod = nil
			if v.Usr != "" && v.Pwd != "" {
				auth = &http.BasicAuth{
					Username: v.Usr,
					Password: v.Pwd,
				}
			}

			log.WithField("name", n).Info("clonning plugin repository")
			r, err := git.PlainClone(d, false, &git.CloneOptions{
				URL:  v.URL,
				Auth: auth,
			})
			if err != nil {
				log.WithFields(log.Fields{
					"err":  err,
					"repo": n,
				}).Errorf("error clonning the plugin repository")
				return []*models.Plugin{}, fmt.Errorf("error clonning the plugin repository %s: %v", n, err)
			}

			dirs, err := afero.ReadDir(fs, d)
			if err != nil {
				log.WithFields(log.Fields{
					"err":  err,
					"repo": n,
				}).Errorf("error listing the plugins of the repository")
				return []*models.Plugin{}, fmt.Errorf("error listing the plugins of repository %s: %v", n, err)
			}

			for _, dir := range dirs {
				if dir.IsDir() && !strings.HasPrefix(dir.Name(), ".") {
					// Get plugins versions from the plugin.json file
					ok, err := afero.Exists(fs, filepath.Join(d, dir.Name(), "plugin.json"))
					if err != nil {
						log.WithFields(log.Fields{
							"err":    err,
							"repo":   n,
							"plugin": dir.Name(),
						}).Error("error checking for the plugin.json file")
						return []*models.Plugin{}, fmt.Errorf("error checking for the plugin.json file of the repository %s and the plugin %s: %v", err, n, dir.Name())
					}

					if !ok {
						continue
					}

					b, err := afero.ReadFile(fs, filepath.Join(d, dir.Name(), "plugin.json"))
					if err != nil {
						log.WithFields(log.Fields{
							"err":    err,
							"repo":   n,
							"plugin": dir.Name(),
						}).Error("error reading the plugin.json file")
						return []*models.Plugin{}, fmt.Errorf("error reading the plugin.json file of the repository %s and the plugin %s: %v", err, n, dir.Name())
					}

					p := &models.Plugin{
						Repo: n,
					}
					if err := json.Unmarshal(b, p); err != nil {
						log.WithFields(log.Fields{
							"err":    err,
							"repo":   n,
							"plugin": dir.Name(),
						}).Error("error parsing the plugin.json file")
						return []*models.Plugin{}, fmt.Errorf("error parsing the plugin.json file of the repository %s and the plugin %s: %v", err, n, dir.Name())
					}

					// Get plugin version from the plugin.json file
					pluginJSONVersion, err := semver.ParseTolerant(p.Version)
					if err != nil {
						fmt.Println(err)
						p.Version = "unknown"
					}
					if p.Version == "" {
						p.Version = "unknown"
					}

					versions := []semver.Version{pluginJSONVersion}

					// Get plugins versions from the tags
					tags, err := r.Tags()
					if err != nil {
						log.WithFields(log.Fields{
							"err":  err,
							"repo": n,
						}).Errorf("error listing the plugin versions of the repository")
						return []*models.Plugin{}, fmt.Errorf("error listing the plugin versions of repository %s: %v", n, err)
					}

					tags.ForEach(func(tag *plumbing.Reference) error {
						if strings.HasPrefix(tag.Name().Short(), dir.Name()) {
							tParts := strings.Split(tag.Name().Short(), "-")
							if len(tParts) == 2 {
								ver, err := semver.ParseTolerant(tParts[1])
								if err != nil {
									fmt.Println(err)
									return nil
								}

								versions = append(versions, ver)
								return nil
							}

							return nil
						}

						return nil
					})

					if len(versions) != 0 {
						semver.Sort(versions)
						p.Version = versions[len(versions)-1].String()
					}

					plugins = append(plugins, p)
				}
			}
		}
	}

	if !repoFound {
		log.WithFields(log.Fields{
			"repo": repos[0],
		}).Errorf("plugin repository not found")
		return []*models.Plugin{}, fmt.Errorf("plugin repository %s not found", repos[0])
	}

	return plugins, nil
}
