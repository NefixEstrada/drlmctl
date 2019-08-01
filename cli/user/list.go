package user

import (
	"io"
	"strings"
	"time"

	"github.com/brainupdaters/drlmctl/core"

	"github.com/jedib0t/go-pretty/table"
)

// List lists all the users in DRLM Core
func List(out io.Writer) {
	rsp, err := core.UserList()
	if err == nil {
		t := table.NewWriter()
		t.SetOutputMirror(out)
		t.AppendHeader(table.Row{"Username", "Auth Type", "Created At"})
		for _, u := range rsp.Users {
			t.AppendRow([]interface{}{u.Usr, strings.Title(strings.ToLower(u.AuthType.String())), time.Unix(u.CreatedAt.Seconds, 0).Format("2006/01/02 15:04:05")})
		}
		t.SetStyle(table.StyleLight)
		t.Render()
	}
}
