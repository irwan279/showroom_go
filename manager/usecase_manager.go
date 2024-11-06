package manager

import (
	"mncbank/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUseCase
	GetLoginUsecase() usecase.LoginUseCase
	GetCustomerUsecase() usecase.CustomerUseCase
}

type usecaseManager struct {
	repoManager RepoManager

	usrUsecase  usecase.UserUseCase
	lgUsecase   usecase.LoginUseCase
	custUsecase usecase.CustomerUseCase
}

var onceLoadVehicleUsecase sync.Once
var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadCustomerUsecase sync.Once
var onceLoadCashUsacase sync.Once
var onceLoadMemberUsacase sync.Once
var onceLoadRentUsacase sync.Once
var onceLoadCreditUsacase sync.Once
var onceLoadReportUsacase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.UserUseCase {
	onceLoadUserUsecase.Do(func() {
		um.usrUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.usrUsecase
}

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUseCase {
	onceLoadLoginUsecase.Do(func() {
		um.lgUsecase = usecase.NewLoginUseCase(um.repoManager.GetLoginRepo(), um.repoManager.GetCustomerRepo())
	})
	return um.lgUsecase
}

func (um *usecaseManager) GetCustomerUsecase() usecase.CustomerUseCase {
	onceLoadCustomerUsecase.Do(func() {
		um.custUsecase = usecase.NewCustomerUseCase(um.repoManager.GetCustomerRepo(), um.repoManager.GetUserRepo())
	})
	return um.custUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
