package usecase

import (
	"log"
	"order-api/internal/order/domain"
	"order-api/internal/order/repository"
	"order-api/internal/broker"

	"github.com/google/uuid"
)

type OrderUseCase struct {
	orderRepo repository.OrderRepository
	rabbitMQ  *broker.RabbitMQ
}

func NewOrderUseCase(orderRepo *repository.OrderRepository, rabbitMQ *broker.RabbitMQ) *OrderUseCase {
	return &OrderUseCase{*orderRepo, rabbitMQ}
}

func (uc *OrderUseCase) CreateOrder(amount float64) (domain.Order, error) {
	order := domain.Order{
		ID:     uuid.New().String(),
		Amount: amount,
		Status: "pending",
	}

	err := uc.orderRepo.Save(order)
	if err != nil {
		return order, err
	}

	// Publicar evento en RabbitMQ
	err = uc.rabbitMQ.Publish("order.created", order)
	if err != nil {
		log.Println("Error al publicar evento:", err)
	}

	return order, nil
}

func (uc *OrderUseCase) GetOrder(id string) (*domain.Order, error) {
	return uc.orderRepo.GetByID(id)
}