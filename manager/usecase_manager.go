package manager

import "enigma-laundry-clean-code/usecase"

type UseCaseManager interface {
	CustomerCase() usecase.CustomerUseCase
	OrderUseCase() usecase.OrderUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) CustomerCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repo.CustRepo())
}

func (u *useCaseManager) OrderUseCase() usecase.OrderUseCase {
	return usecase.NewOrderUseCase(u.repo.OrderRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
