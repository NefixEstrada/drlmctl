package core

import (
	"fmt"

	"github.com/brainupdaters/drlmctl/models"

	"github.com/brainupdaters/drlm-common/pkg/os"
	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	log "github.com/sirupsen/logrus"
)

// AgentRequestList lists all the agent requests in DRLM
func AgentRequestList() ([]*models.Agent, error) {
	rsp, err := Client.AgentRequestList(prepareCtx(), &drlm.AgentRequestListRequest{})
	if err != nil {
		log.WithFields(log.Fields{
			"api": API,
			"err": err,
		}).Error("error listing the requested DRLM Agents")
		return []*models.Agent{}, fmt.Errorf("error listing the requested DRLM Agents: %v", err)
	}

	var agents []*models.Agent
	for _, a := range rsp.Agents {
		agents = append(agents, &models.Agent{
			Host:    a.Host,
			OS:      os.OS(a.Os),
			Arch:    os.Arch(a.Arch),
			Version: a.Version,
		})

	}

	return agents, nil
}

// AgentRequestAccept accepts an Agent to DRLM
func AgentRequestAccept(host string) error {
	_, err := Client.AgentAccept(prepareCtx(), &drlm.AgentAcceptRequest{
		Host: host,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"api":  API,
			"err":  err,
			"host": host,
		}).Error("error accepting the DRLM Agent")
		return fmt.Errorf("error accepting the DRLM Agent '%s': %v", host, err)
	}

	return nil
}
