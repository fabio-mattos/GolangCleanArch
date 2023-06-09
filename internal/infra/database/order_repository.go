package database

import (
	"database/sql"

	"github.com/fabio-mattos/GolangCleanArch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetAll() ([]entity.Order, error) {
	rows, err := r.Db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return []entity.Order{}, err
	}
	defer rows.Close()

	var ordersOutput []entity.Order
	for rows.Next() {
		var order entity.Order
		err = rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return []entity.Order{}, err
		}
		ordersOutput = append(ordersOutput, order)
	}
	return ordersOutput, nil
}
