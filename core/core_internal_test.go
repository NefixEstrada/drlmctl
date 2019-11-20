// SPDX-License-Identifier: AGPL-3.0-only

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
	"github.com/brainupdaters/drlm-common/pkg/test"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type TestCoreInternalSuite struct {
	test.Test
}

func TestCoreInternal(t *testing.T) {
	suite.Run(t, new(TestCoreInternalSuite))
}

func (s *TestCoreInternalSuite) TestReadCert() {

	s.Run("should read the certificate correctly", func() {
		s.GenerateCfg(test.CfgPathCtl, func() { cfg.Init("") })
		s.GenerateCert("server", filepath.Dir(cfg.Config.Core.CertPath))

		cp, err := readCert()
		s.Nil(err)
		s.NotEqual(&x509.CertPool{}, cp)
	})

	s.Run("should return an error if there's an error reading the certificate file", func() {
		s.GenerateCfg(test.CfgPathCtl, func() { cfg.Init("") })

		cp, err := readCert()
		s.EqualError(err, "error reading the certificate file: open cert/server.crt: file does not exist")
		s.Equal(&x509.CertPool{}, cp)
	})

	s.Run("should return an error if there's an error parsing the certificate", func() {
		s.GenerateCfg(test.CfgPathCtl, func() { cfg.Init("") })
		// cmnTests.GenerateCert(t, fs.FS, "server", filepath.Dir(cfg.Config.Core.CertPath))

		afero.WriteFile(fs.FS, cfg.Config.Core.CertPath, []byte("This isn't a cert!"), 0644)

		cp, err := readCert()
		s.EqualError(err, "error parsing the certificate: invalid certificate")
		s.Equal(&x509.CertPool{}, cp)
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
}
