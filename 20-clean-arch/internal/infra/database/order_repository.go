package database

import (
	"database/sql"

	"github.com/allanmaral/go-expert/20-clean-arch/internal/entity"
)

type OrderRepositorySQL struct {
	db *sql.DB
}

var _ entity.OrderRepository = (*OrderRepositorySQL)(nil)

func NewOrderRepository(db *sql.DB) *OrderRepositorySQL {
	return &OrderRepositorySQL{db}
}

func (r *OrderRepositorySQL) Save(order *entity.Order) error {
	stmt, err := r.db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepositorySQL) GetTotal() (int, error) {
	var total int
	if err := r.db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
