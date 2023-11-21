package usecase

import (
	"enigma-laundry-clean-code/model/entity"
	"enigma-laundry-clean-code/repository"
)

type OrderUseCase interface {
	FindByID(id string) (entity.Orders, error)
}

type orderUseCase struct {
	repo repository.OrderRepository
}

func (o *orderUseCase) FindByID(id string) (entity.Orders, error) {
	order, err := o.repo.GetByID(id)
	if err != nil {
		return entity.Orders{}, err
	}

	return order, nil
}

func NewOrderUseCase(repo repository.OrderRepository) OrderUseCase {
	return &orderUseCase{repo: repo}
}
