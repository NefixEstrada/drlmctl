// SPDX-License-Identifier: AGPL-3.0-only

package agent

import (
	"os"

	"github.com/brainupdaters/drlmctl/core"

	"github.com/jedib0t/go-pretty/table"
)

// List lists all the agents in DRLM
func List() {
	agents, err := core.AgentList()
	if err == nil {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Host", "Port", "User", "OS", "Arch"})
		for _, a := range agents {
			t.AppendRow([]interface{}{a.Host, a.Port, a.Usr, a.OS, a.Arch})
		}
		t.SetStyle(table.StyleLight)
		t.Render()
	}
}
