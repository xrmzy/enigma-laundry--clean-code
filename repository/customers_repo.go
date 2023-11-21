package repository

import (
	"database/sql"
	"enigma-laundry-clean-code/model/entity"
	"time"
)

type CustomerRepository interface {
	GetByID(id string) (entity.Customer, error)
	Create(payload entity.Customer) (entity.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) GetByID(id string) (entity.Customer, error) {
	var customer entity.Customer
	err := c.db.QueryRow(`SELECT cust_id, cust_name, phone_number, address, email FROM customers WHERE cust_id = $1`, id).
		Scan(
			&customer.Id,
			&customer.Name,
			&customer.PhoneNumber,
			&customer.Address,
			&customer.Email,
		)
	if err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) Create(payload entity.Customer) (entity.Customer, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return entity.Customer{}, err
	}

	var customer entity.Customer
	err = tx.QueryRow(
		`
		INSERT INTO customers (cust_id, cust_name, phone_number, address, email)
    VALUES ($1, $2, $3, $4, $5)
		RETURNING cust_id, status, created_at, updated_at
		`, payload.Id,
		time.Now(),
	).Scan(
		&customer.Id,
		&customer.Name,
		&customer.PhoneNumber,
		&customer.Address,
		&customer.Email,
	)
	if err != nil {
		tx.Rollback()
		return entity.Customer{}, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return entity.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) GetAll(entity.Customer, error) {
	rows, err := c.db.Query(`SELECT cust_id, cust_name, phone_number, address, email, created_at, updated_at FROM customers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		var customer entity.Customer
		if err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.PhoneNumber,
			&customer.Address,
			&customer.Email,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
