// SPDX-License-Identifier: AGPL-3.0-only

package plugin

import (
	"os"
	"strings"

	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/models"
	"github.com/brainupdaters/drlmctl/software"
)

// Add adds a plugin to a DRLM Agent
func Add(host, plugin, v string) {
	parts := strings.Split(plugin, "/")
	if len(parts) < 2 {
		panic("AAaaaaa")
	}

	p := &models.Plugin{
		Repo:    parts[0],
		Name:    parts[1],
		Version: v,
	}

	a, err := core.AgentGet(host)
	if err != nil {
		os.Exit(1)
	}

	bin, err := software.CompilePlugin(p, a.OS, a.Arch, p.Version)
	if err != nil {
		panic(err)
	}

	core.PluginAdd(host, p, bin)
}
