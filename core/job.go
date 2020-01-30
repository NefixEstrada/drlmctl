// SPDX-License-Identifier: AGPL-3.0-only

package core

import (
	"fmt"

	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	log "github.com/sirupsen/logrus"
)

// JobSchedule schedules a new job
func JobSchedule(host, job, config string) error {
	_, err := Client.JobSchedule(prepareCtx(), &drlm.JobScheduleRequest{
		AgentHost: host,
		Name:      job,
		Config:    config,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"api": API,
			"err": err,
		}).Error("error scheduling the job")
		return fmt.Errorf("error scheduling the job: %v", err)
	}

	return nil
}
