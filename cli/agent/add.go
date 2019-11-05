package agent

import (
	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/models"
)

// Add connects to the Agent host, creates the drlm user and copies the keys to that user, which has to be admin
func Add(host string, port int, usr, pwd string) {
	a := &models.Agent{
		Host: host,
		Port: int32(port),
		Usr:  usr,
		Pwd:  pwd,
	}

	core.AgentAdd(a)
}
