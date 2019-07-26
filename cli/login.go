package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/core"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

// Login logs in into the DRLM Core server and saves the token and the expiration time in the config
func Login(usr, pwd string) {
	if usr == "" || pwd == "" {
		r := bufio.NewReader(os.Stdin)

		fmt.Print("Enter username: ")

		var err error
		usr, err = r.ReadString('\n')
		if err != nil {
			log.Fatalf("error logging in: error reading the username: %v", err)
		}
		usr = strings.TrimSpace(usr)

		fmt.Print("Enter password: ")
		bPwd, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatalf("error logging in: error reading the password: %v", err)
		}
		pwd = strings.TrimSpace(string(bPwd))

		fmt.Print("\n")
	}

	rsp, err := core.UserLogin(usr, pwd)
	if err != nil {
		os.Exit(1)
	}

	if err := cfg.SaveTkn(rsp.Tkn, time.Unix(rsp.TknExpiration.Seconds, 0)); err != nil {
		log.Fatalf("error saving the new token to the configuration: %v", err)
	}
}
