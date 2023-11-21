package usecase

import (
	"enigma-laundry-clean-code/model/dto"
	"enigma-laundry-clean-code/model/entity"
	"enigma-laundry-clean-code/repository"
	"time"
)

type CustomerUseCase interface {
	FindByID(id string) (entity.Customer, error)
	CreatedCustomer(payload dto.CustomersRequestDto) (entity.Customer, error)
	GetAll() ([]entity.Customer, error)
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (c *customerUseCase) FindByID(id string) (entity.Customer, error) {
	customer, err := c.repo.GetByID(id)
	if err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func (c *customerUseCase) CreatedCustomer(payload dto.CustomersRequestDto) (entity.Customer, error) {
	newCustomer := entity.Customer{
		Id:          payload.Id,
		Name:        payload.Name,
		PhoneNumber: payload.PhoneNumber,
		Address:     payload.Address,
		Email:       payload.Email,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	customer, err := c.repo.Create(newCustomer)
	if err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func (c *customerUseCase) GetAll() ([]entity.Customer, error) {
	customers, err := c.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
