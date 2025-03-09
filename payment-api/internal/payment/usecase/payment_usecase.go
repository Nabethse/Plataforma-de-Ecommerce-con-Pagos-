package usecase

import (
	"log"
	"payment-api/internal/payment/domain"
	"payment-api/internal/payment/repository"
	"payment-api/internal/broker"

	"github.com/google/uuid"
)

type PaymentUseCase struct {
	paymentRepo repository.PaymentRepository
	rabbitMQ    *broker.RabbitMQ
}

func NewPaymentUseCase(paymentRepo *repository.PaymentRepository, rabbitMQ *broker.RabbitMQ) *PaymentUseCase {
	return &PaymentUseCase{*paymentRepo, rabbitMQ}
}

func (uc *PaymentUseCase) ProcessPayment(orderID string, amount float64) (domain.Payment, error) {
	payment := domain.Payment{
		ID:      uuid.New().String(),
		OrderID: orderID,
		Amount:  amount,
		Status:  "completed",
	}

	err := uc.paymentRepo.Save(payment)
	if err != nil {
		return payment, err
	}

	// Publicar evento en RabbitMQ
	err = uc.rabbitMQ.Publish("payment.processed", payment)
	if err != nil {
		log.Println("Error al publicar evento:", err)
	}

	return payment, nil
}

func (uc *PaymentUseCase) GetPayment(id string) (*domain.Payment, error) {
	return uc.paymentRepo.GetByID(id)
}