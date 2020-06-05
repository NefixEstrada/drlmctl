package request

import (
	"os"

	"github.com/brainupdaters/drlmctl/core"

	"github.com/jedib0t/go-pretty/table"
)

// List lists all the agent requests in DRLM
func List() {
	agents, err := core.AgentRequestList()
	if err == nil {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Host", "Version", "OS", "Arch", "Accepted"})
		for _, a := range agents {
			t.AppendRow([]interface{}{a.Host, a.Version, a.OS, a.Arch, false})
		}
		t.SetStyle(table.StyleLight)
		t.Render()
	}
}
