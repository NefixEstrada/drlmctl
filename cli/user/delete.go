package user

import "github.com/brainupdaters/drlmctl/core"

// Delete deletes an user from DRLM Core
func Delete(usr string) {
	core.UserDelete(usr)
}
