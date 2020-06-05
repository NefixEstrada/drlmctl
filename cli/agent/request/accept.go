package request

import "github.com/brainupdaters/drlmctl/core"

// Accept accepts an Agent to DRLM
func Accept(host string) {
	core.AgentRequestAccept(host)
}
