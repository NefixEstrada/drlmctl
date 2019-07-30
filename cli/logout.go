package cli

import (
	"time"

	"github.com/brainupdaters/drlmctl/cfg"
)

// Logout removes the token and the token expiration time, forcing the user to login again
func Logout() {
	cfg.SaveTkn("", time.Unix(0, 0))
}
