package manager

import (
	"mncbank/repository"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repository.UserRepo
	GetLoginRepo() repository.LoginRepo
	GetCustomerRepo() repository.CustomerRepo
	GetMemberRepo() repository.MemberRepo
}

type repoManager struct {
	infraManager InfraManager
	usrRepo      repository.UserRepo
	lgRepo       repository.LoginRepo
	custRepo     repository.CustomerRepo
	memberRepo   repository.MemberRepo
}

var onceLoadVehicleRepo sync.Once
var onceLoadUserRepo sync.Once
var onceLoadLoginRepo sync.Once
var onceLoadCustomerRepo sync.Once
var onceLoadCashRepo sync.Once
var onceLoadMemberRepo sync.Once
var onceLoadRentRepo sync.Once
var onceLoadCreditRepo sync.Once
var onceLoadReportRepo sync.Once

func (rm *repoManager) GetUserRepo() repository.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repository.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func (rm *repoManager) GetLoginRepo() repository.LoginRepo {
	onceLoadLoginRepo.Do(func() {
		rm.lgRepo = repository.NewLoginRepo(rm.infraManager.GetDB())
	})
	return rm.lgRepo
}

func (rm *repoManager) GetCustomerRepo() repository.CustomerRepo {
	onceLoadCustomerRepo.Do(func() {
		rm.custRepo = repository.NewCustomerRepo(rm.infraManager.GetDB())
	})
	return rm.custRepo
}

func (rm *repoManager) GetMemberRepo() repository.MemberRepo {
	onceLoadMemberRepo.Do(func() {
		rm.memberRepo = repository.NewMemberDbRepository(rm.infraManager.GetDB())
	})
	return rm.memberRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
