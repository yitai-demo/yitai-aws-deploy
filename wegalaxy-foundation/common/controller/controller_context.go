package controller

import (
	gp_domain "github.com/degalaxy/gp-common/domain"
	"github.com/degalaxy/gp-common/lockmanager"
	gp_repo "github.com/degalaxy/gp-common/repository"
	"github.com/degalaxy/helper/dahelper"
)

type ControllerContext struct {
	VerifyTokenService gp_domain.VerifyTokenService
	DaHelper           dahelper.DAHelper
	DbConnection       *gp_repo.DBConnection
	LockManager        lockmanager.LockManager
}

func NewControllerContext(
	verifyTokenService gp_domain.VerifyTokenService,
	daHelper dahelper.DAHelper,
	dbConnection *gp_repo.DBConnection,
	lockManager lockmanager.LockManager,
) (*ControllerContext, error) {
	return &ControllerContext{
		VerifyTokenService: verifyTokenService,
		DaHelper:           daHelper,
		DbConnection:       dbConnection,
		LockManager:        lockManager,
	}, nil
}
