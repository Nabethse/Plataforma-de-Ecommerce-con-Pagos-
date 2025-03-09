package main

import (
	"log"
	"order-api/internal/broker"
	"order-api/internal/order/handler"
	"order-api/internal/order/repository"
	"order-api/internal/order/usecase"
	"order-api/db"
	"order-api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración
	config.Load()

	// Conectar a la base de datos
	database := db.ConnectDB()

	// Inicializar RabbitMQ
	rabbitMQ, err := broker.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Iniciar repositorio, caso de uso y handler
	orderRepo := repository.NewOrderRepository(database)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, rabbitMQ)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	// Configurar router
	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)

	// Iniciar servidor
	r.Run(":8080")
}
