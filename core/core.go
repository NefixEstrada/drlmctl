package core

import (
	"bufio"
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/brainupdaters/drlmctl/cfg"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"golang.org/x/crypto/ssh/terminal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// API is the API version of the client
const API string = "v1.0.0"

// Client is the DRLM Core client
var Client drlm.DRLMClient

// Conn is the actual connection with the DRLM Core server
var Conn *grpc.ClientConn

// Init initializes the DRLM Core client
func Init() {
	var grpcDialOptions = []grpc.DialOption{}

	if cfg.Config.Core.TLS {
		cp, err := readCert()
		if err != nil {
			log.WithFields(log.Fields{
				"cert_path": cfg.Config.Core.CertPath,
			}).Fatalf("error loading the TLS certificate: %v", err)
		}

		cred := credentials.NewClientTLSFromCert(cp, "")

		grpcDialOptions = append(grpcDialOptions, grpc.WithTransportCredentials(cred))
	} else {
		grpcDialOptions = append(grpcDialOptions, grpc.WithInsecure())
	}

	Conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.Config.Core.Host, cfg.Config.Core.Port), grpcDialOptions...)
	if err != nil {
		log.WithFields(log.Fields{
			"host": cfg.Config.Core.Host,
			"port": cfg.Config.Core.Port,
		}).Fatalf("error creating the client for DRLM Core: %v", err)
	}

	Client = drlm.NewDRLMClient(Conn)
}

// readCert reads the DRLM Core certificate from the configuration path using a FS sent as parameter
func readCert() (*x509.CertPool, error) {
	b, err := afero.ReadFile(fs.FS, cfg.Config.Core.CertPath)
	if err != nil {
		return &x509.CertPool{}, fmt.Errorf("error reading the certificate file: %v", err)
	}

	p := x509.NewCertPool()
	if ok := p.AppendCertsFromPEM(b); !ok {
		return &x509.CertPool{}, errors.New("error parsing the certificate: invalid certificate")
	}

	return p, nil
}

func prepareCtx() context.Context {
	if cfg.Config.Core.TknExpiration.Before(time.Now().Add(30 * time.Second)) {
		rsp, err := UserTokenRenew()
		if err != nil {
			cfg.SaveTkn(rsp.Tkn, time.Unix(0, 0))
		} else {
			cfg.SaveTkn(rsp.Tkn, time.Unix(rsp.TknExpiration.Seconds, 0))
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
			log.Fatalf("error saving the new token to the configuration: %v", err)
		}
	}

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"api": API,
		"tkn": cfg.Config.Core.Tkn,
	}))
	return ctx
}
