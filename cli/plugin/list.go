// SPDX-License-Identifier: AGPL-3.0-only

package plugin

import (
	"os"

	"github.com/brainupdaters/drlmctl/models"
	"github.com/brainupdaters/drlmctl/plugin"

	"github.com/jedib0t/go-pretty/table"
)

// List lists all the plugins available
func List(repo string) {
	var (
		plugins []*models.Plugin
		err     error
	)
	if repo != "" {
		plugins, err = plugin.List(repo)
	} else {
		plugins, err = plugin.List()
	}

	if err == nil {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Name", "Version", "Description"})
		for _, p := range plugins {
			t.AppendRow([]interface{}{p.Repo + "/" + p.Name, p.Version, p.Description})
		}
		t.SetStyle(table.StyleLight)
		t.Render()
	}
}
