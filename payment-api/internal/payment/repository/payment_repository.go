package repository

import (
	"database/sql"
	"log"
	"payment-api/internal/payment/domain"

	_ "github.com/lib/pq"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (r *PaymentRepository) Save(payment domain.Payment) error {
	_, err := r.db.Exec("INSERT INTO payments (id, order_id, amount, status) VALUES ($1, $2, $3, $4)",
		payment.ID, payment.OrderID, payment.Amount, payment.Status)
	if err != nil {
		log.Println("Error al guardar el pago:", err)
	}
	return err
}

func (r *PaymentRepository) GetByID(id string) (*domain.Payment, error) {
	payment := &domain.Payment{}
	err := r.db.QueryRow("SELECT id, order_id, amount, status FROM payments WHERE id = $1", id).
		Scan(&payment.ID, &payment.OrderID, &payment.Amount, &payment.Status)
	if err != nil {
		return nil, err
	}
	return payment, nil
}