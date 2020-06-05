// SPDX-License-Identifier: AGPL-3.0-only

package plugin

import (
	"os"

	"github.com/brainupdaters/drlmctl/models"
	"github.com/brainupdaters/drlmctl/plugin"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/afero"
)

// List lists all the plugins available
func List(fs afero.Fs, repo string) {
	var (
		plugins []*models.Plugin
		err     error
	)
	if repo != "" {
		plugins, err = plugin.List(fs, repo)
	} else {
		plugins, err = plugin.List(fs)
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
