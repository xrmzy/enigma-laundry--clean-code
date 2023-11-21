package manager

import "enigma-laundry-clean-code/repository"

type RepoManager interface {
	CustRepo() repository.CustomerRepository
	OrderRepo() repository.OrderRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) CustRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Connection())
}

func (r *repoManager) OrderRepo() repository.OrderRepository {
	return repository.NewOrderReposiotry(r.infra.Connection())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
