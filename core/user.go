package core

import (
	"context"
	"fmt"

	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/models"

	drlm "github.com/brainupdaters/drlm-common/pkg/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

// UserLogin logs into DRLM Core
func UserLogin(usr, pwd string) (*drlm.UserLoginResponse, error) {
	req := &drlm.UserLoginRequest{
		Usr: usr,
		Pwd: pwd,
	}

	rsp, err := Client.UserLogin(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"api": API,
	})), req)
	if err != nil {
		log.WithFields(log.Fields{
			"api": API,
			"usr": usr,
			"err": err,
		}).Error("error logging into DRLM Core")
		return &drlm.UserLoginResponse{}, fmt.Errorf("error logging into DRLM Core: %v", err)
	}

	return rsp, nil
}

// UserTokenRenew renews the token
func UserTokenRenew() (*drlm.UserTokenRenewResponse, error) {
	req := &drlm.UserTokenRenewRequest{}

	rsp, err := Client.UserTokenRenew(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"api": API,
		"tkn": cfg.Config.Core.Tkn,
	})), req)
	if err != nil {
		log.WithFields(log.Fields{
			"api":            API,
			"tkn":            cfg.Config.Core.Tkn,
			"tkn_expiration": cfg.Config.Core.TknExpiration,
			"err":            err,
		}).Error("error renewing the user token")

		return &drlm.UserTokenRenewResponse{}, fmt.Errorf("error renewing the user token: %v", err)
	}

	return rsp, nil
}

// UserAdd adds a new User to DRLM Core
func UserAdd(usr *models.User) error {
	req := &drlm.UserAddRequest{
		Usr: usr.Username,
		Pwd: usr.Password,
	}

	_, err := Client.UserAdd(prepareCtx(), req)
	if err != nil {
		log.WithFields(log.Fields{
			"api": API,
			"usr": usr.Username,
			"err": err,
		}).Error("error adding the user to DRLM Core")
		return fmt.Errorf("error adding the user to DRLM Core: %v", err)
	}

	return nil
}

// UserDelete removes an user from DRLM Core
func UserDelete(usr string) error {
	req := &drlm.UserDeleteRequest{
		Usr: usr,
	}

	_, err := Client.UserDelete(prepareCtx(), req)
	if err != nil {
		log.WithFields(log.Fields{
			"api": API,
			"usr": usr,
			"err": err,
		}).Error("error deleting the user from DRLM Core")
		return fmt.Errorf("error deleting the user from DRLM Core: %v", err)
	}

	return nil
}

// UserList lists all the users in DRLM Core
func UserList() (*drlm.UserListResponse, error) {
	req := &drlm.UserListRequest{}

	rsp, err := Client.UserList(prepareCtx(), req)
	if err != nil {
		log.WithFields(log.Fields{
			"api": API,
			"err": err,
		}).Error("error listing the users from DRLM Core")
		return &drlm.UserListResponse{}, fmt.Errorf("error listing the users from DRLM Core: %v", err)
	}

	return rsp, nil
}
