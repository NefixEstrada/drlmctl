// SPDX-License-Identifier: AGPL-3.0-only

package agent

import (
	"github.com/brainupdaters/drlmctl/core"
)

// Add connects to the Agent host, creates the drlm user and copies the keys to that user, which has to be admin
func Add(host string) {
	core.AgentAdd(host)
}
