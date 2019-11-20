// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"github.com/brainupdaters/drlmctl/core"
	"github.com/brainupdaters/drlmctl/models"
)

// Add adds a new user to DRLM Core
func Add(usr *models.User) {
	core.UserAdd(usr)
}
