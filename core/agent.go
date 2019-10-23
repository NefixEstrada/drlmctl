package core

import (
	"fmt"
	"io"

	"github.com/brainupdaters/drlmctl/models"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/os"
	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	log "github.com/sirupsen/logrus"
)

// AgentList returns all the agents in DRLM
func AgentList() ([]*models.Agent, error) {
	rsp, err := Client.AgentList(prepareCtx(), &drlm.AgentListRequest{})
	if err != nil {
		log.WithFields(log.Fields{
			"api": API,
			"err": err,
		}).Error("error listing the DRLM Agents")
		return []*models.Agent{}, fmt.Errorf("error listing the DRLM Agents: %v", err)
	}

	var agents []*models.Agent
	for _, a := range rsp.Agents {
		agents = append(agents, &models.Agent{
			Host: a.Host,
			Port: a.Port,
			Usr:  a.User,
			OS:   os.OS(a.Os),
			Arch: os.Arch(a.Arch),
		})
	}

	return agents, nil
}

// AgentGet returns the agent by it's host
func AgentGet(host string) (*models.Agent, error) {
	req := &drlm.AgentGetRequest{
		Host: host,
	}

	rsp, err := Client.AgentGet(prepareCtx(), req)
	if err != nil {
		log.WithFields(log.Fields{
			"api":  API,
			"host": host,
			"err":  err,
		}).Error("error getting the DRLM Agent")
		return &models.Agent{}, fmt.Errorf("error getting the DRLM Agent: %v", err)
	}

	return &models.Agent{
		Host: host,
		OS:   os.OS(rsp.Os),
		Arch: os.Arch(rsp.Arch),
	}, nil
}

// AgentAdd adds a new Agent to DRLM Core
func AgentAdd(a *models.Agent) error {
	req := &drlm.AgentAddRequest{
		Host:     a.Host,
		Port:     a.Port,
		User:     a.Usr,
		Password: a.Pwd,
	}

	_, err := Client.AgentAdd(prepareCtx(), req)
	if err != nil {
		log.WithFields(log.Fields{
			"api":  API,
			"host": a.Host,
			"port": a.Port,
			"usr":  a.Usr,
			"err":  err,
		}).Error("error adding the DRLM Agent")
		return fmt.Errorf("error adding the DRLM Agent: %v", err)
	}

	return nil
}

// AgentInstall installs the Agent binary to the server
func AgentInstall(host string, binPath string) error {
	f, err := fs.FS.Open(binPath)
	if err != nil {
		log.WithFields(log.Fields{
			"path": binPath,
			"err":  err,
		}).Error("error opening the DRLM Agent binary")
		return fmt.Errorf("error opening the DRLM Agent binary: %v", err)
	}

	stream, err := Client.AgentInstall(prepareCtx())
	if err != nil {
		log.WithField("err", err).Error("error creating the DRLM Agent binary upload stream")
		return fmt.Errorf("error creating the DRLM Agent binary upload stream: %v", err)
	}
	defer stream.CloseSend()

	buf := make([]byte, 4096)

	writting := true
	for writting {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				writting = false
				err = nil
				continue
			}

			log.WithField("err", err).Error("error copying the DRLM Agent binary to the GRPC upload buffer")
			return fmt.Errorf("error copying the DRLM Agent binary to the GRPC upload buffer: %v", err)
		}

		err = stream.Send(&drlm.AgentInstallRequest{
			Host: host,
			Bin:  buf[:n],
		})
		if err != nil {
			log.WithField("err", err).Error("error sending the DRLM Agent binary chunk to the server")
			return fmt.Errorf("error sending the DRLM Agent binary chunk to the server: %v", err)
		}
	}

	rsp, err := stream.CloseAndRecv()
	if err != nil {
		log.WithField("err", err).Error("error closing the DRLM Agent binary upload stream")
		return fmt.Errorf("error closing the DRLM Agent binary upload stream: %v", err)
	}

	if rsp.Code != drlm.AgentInstallResponse_OK {
		log.WithFields(log.Fields{
			"code":    rsp.Code,
			"message": rsp.Message,
		}).Error("error uploading the DRLM Agent binary to the server")
		return fmt.Errorf("error uploading the DRLM Agent binary to the server: %v", err)
	}

	return nil
}
