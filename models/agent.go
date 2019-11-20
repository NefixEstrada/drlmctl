// SPDX-License-Identifier: AGPL-3.0-only

package models

import "github.com/brainupdaters/drlm-common/pkg/os"

// Agent is the model for a DRLM Agent
type Agent struct {
	Host string
	Port int32
	Usr  string
	Pwd  string
	OS   os.OS
	Arch os.Arch
}
