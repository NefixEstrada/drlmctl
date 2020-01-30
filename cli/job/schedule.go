// SPDX-License-Identifier: AGPL-3.0-only

package job

import (
	"os"

	"github.com/brainupdaters/drlmctl/core"
)

// Schedule schedules a new job in an agent
func Schedule(host, job, config string) {
	if err := core.JobSchedule(host, job, config); err != nil {
		os.Exit(1)
	}
}
