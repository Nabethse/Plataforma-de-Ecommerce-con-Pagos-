package repository

import (
	"database/sql"
	"log"
	"order-api/internal/order/domain"

	_ "github.com/lib/pq"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Save(order domain.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (id, amount, status) VALUES ($1, $2, $3)", order.ID, order.Amount, order.Status)
	if err != nil {
		log.Println("Error al guardar la orden:", err)
	}
	return err
}

func (r *OrderRepository) GetByID(id string) (*domain.Order, error) {
	order := &domain.Order{}
	err := r.db.QueryRow("SELECT id, amount, status FROM orders WHERE id = $1", id).
		Scan(&order.ID, &order.Amount, &order.Status)
	if err != nil {
		return nil, err
	}
	return order, nil
}