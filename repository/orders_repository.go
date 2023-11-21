package repository

import (
	"database/sql"
	"enigma-laundry-clean-code/model/entity"
)

type OrderRepository interface {
	GetByID(id string) (entity.Orders, error)
}

type orderRepository struct {
	db *sql.DB
}

func (o *orderRepository) GetByID(id string) (entity.Orders, error) {
	var order entity.Orders
	err := o.db.QueryRow(`
	SELECT order_id, cust_id, cust_name, service, unit, outlet_name, order_date, status FROM orders WHERE order_id = $1;`, id).Scan(
		&order.OrderId,
		&order.CustomerId,
		&order.CustomerName,
		&order.Service,
		&order.Unit,
		&order.OutletName,
		&order.OrderDate,
		&order.Status,
	)
	if err != nil {
		return entity.Orders{}, err
	}
	return order, nil
}

func NewOrderReposiotry(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}
