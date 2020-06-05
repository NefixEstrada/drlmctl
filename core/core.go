// SPDX-License-Identifier: AGPL-3.0-only

package core

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/spf13/afero"

	cmnCore "github.com/brainupdaters/drlm-common/pkg/core"
	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"google.golang.org/grpc/metadata"
)

// API is the API version of the client
const API string = "v1.0.0"

// Client is the DRLM Core client
var Client drlm.DRLMClient

// Init initializes the DRLM Core client
func Init(fs afero.Fs) {
	Client, _ = cmnCore.NewClient(
		fs,
		cfg.Config.Core.TLS,
		cfg.Config.Core.CertPath,
		cfg.Config.Core.Host,
		cfg.Config.Core.Port,
	)
}

func prepareCtx() context.Context {
	if cfg.Config.Core.TknExpiration.Before(time.Now().Add(30 * time.Second)) {
		rsp, err := UserTokenRenew()
		if err != nil {
			if err := cfg.SaveTkn(rsp.Tkn, time.Unix(0, 0)); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := cfg.SaveTkn(rsp.Tkn, time.Unix(rsp.TknExpiration.Seconds, 0)); err != nil {
				log.Fatal(err)
			}
		}
	}

	if cfg.Config.Core.Tkn == "" {
		r := bufio.NewReader(os.Stdin)

		fmt.Print("Enter username: ")

		var err error
		usr, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("error logging in: error reading the username: %v", err)
		}
		usr = strings.TrimSpace(usr)

		fmt.Print("Enter password: ")
		bPwd, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatalf("error logging in: error reading the password: %v", err)
		}
		pwd := strings.TrimSpace(string(bPwd))

		fmt.Print("\n")

		rsp, err := UserLogin(usr, pwd)
		if err != nil {
			os.Exit(1)
		}

		if err := cfg.SaveTkn(rsp.Tkn, time.Unix(rsp.TknExpiration.Seconds, 0)); err != nil {
			log.Fatal(err)
		}
	}

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"api": API,
		"tkn": cfg.Config.Core.Tkn,
	}))
	return ctx
}
