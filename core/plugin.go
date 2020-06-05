// SPDX-License-Identifier: AGPL-3.0-only

package core

import (
	"fmt"
	"io"

	"github.com/brainupdaters/drlmctl/models"
	"github.com/spf13/afero"

	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	log "github.com/sirupsen/logrus"
)

// PluginAdd adds a plugin to a DRLM Agent
func PluginAdd(fs afero.Fs, host string, p *models.Plugin, binPath string) error {
	f, err := fs.Open(binPath)
	if err != nil {
		log.WithFields(log.Fields{
			"plugin": p.Name,
			"path":   binPath,
			"err":    err,
		}).Error("error opening the plugin binary")
		return fmt.Errorf("error opening the plugin %s binary at %s: %v", p.Name, binPath, err)
	}

	stream, err := Client.AgentPluginAdd(prepareCtx())
	if err != nil {
		panic("AAAAAA")
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

			panic("AAA")
		}

		err = stream.Send(&drlm.AgentPluginAddRequest{
			Host:    host,
			Repo:    p.Repo,
			Plugin:  p.Name,
			Version: p.Version,
			Bin:     buf[:n],
		})
		if err != nil {
			panic("AAAAAA")
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	return nil
}
