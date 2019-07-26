package core

import (
	"context"
	"crypto/x509"
	"path/filepath"
	"testing"
	"time"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/utils/tests"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	cmnTests "github.com/brainupdaters/drlm-common/pkg/tests"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestReadCert(t *testing.T) {
	assert := assert.New(t)

	t.Run("should read the certificate correctly", func(t *testing.T) {
		tests.GenerateCfg(t)
		cmnTests.GenerateCert(t, fs.FS, "server", filepath.Dir(cfg.Config.Core.CertPath))

		cp, err := readCert()
		assert.Nil(err)
		assert.NotEqual(&x509.CertPool{}, cp)
	})

	t.Run("should return an error if there's an error reading the certificate file", func(t *testing.T) {
		tests.GenerateCfg(t)

		cp, err := readCert()
		assert.EqualError(err, "error reading the certificate file: open cert/server.crt: file does not exist")
		assert.Equal(&x509.CertPool{}, cp)
	})

	t.Run("should return an error if there's an error parsing the certificate", func(t *testing.T) {
		tests.GenerateCfg(t)
		cmnTests.GenerateCert(t, fs.FS, "server", filepath.Dir(cfg.Config.Core.CertPath))

		afero.WriteFile(fs.FS, cfg.Config.Core.CertPath, []byte("This isn't a cert!"), 0644)

		cp, err := readCert()
		assert.EqualError(err, "error parsing the certificate: invalid certificate")
		assert.Equal(&x509.CertPool{}, cp)
	})
}

func TestPrepareCtx(t *testing.T) {
	assert := assert.New(t)

	t.Run("should prepare the context correctly", func(t *testing.T) {
		tests.GenerateCfg(t)

		ctx := prepareCtx()

		assert.Equal(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"api": API,
			"tkn": "thisisatoken",
		})), ctx)
	})

	t.Run("should renew the token if it's required", func(t *testing.T) {
		tests.GenerateCfg(t)
		cfg.Config.Core.TknExpiration = time.Now()

		newExpiration := time.Now().Add(5 * time.Minute)

		theCoreClientMock := &tests.CoreClientMock{}
		theCoreClientMock.On("UserTokenRenew", metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api", API, "tkn", "thisisatoken")), &drlm.UserTokenRenewRequest{}, []grpc.CallOption(nil)).Return(
			&drlm.UserTokenRenewResponse{
				Tkn:           "thisisanewtoken",
				TknExpiration: &timestamp.Timestamp{Seconds: newExpiration.Unix()},
			}, nil,
		)
		Client = theCoreClientMock

		ctx := prepareCtx()

		assert.Equal(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"api": API,
			"tkn": "thisisanewtoken",
		})), ctx)
		assert.Equal("thisisanewtoken", cfg.Config.Core.Tkn)
	})

	t.Run("should call the Login command if the token renew fails", func(t *testing.T) {

	})
}
