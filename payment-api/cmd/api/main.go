package main

import (
	"log"
	"payment-api/internal/broker"
	"payment-api/internal/payment/handler"
	"payment-api/internal/payment/repository"
	"payment-api/internal/payment/usecase"
	"payment-api/db"
	"payment-api/config"

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
	paymentRepo := repository.NewPaymentRepository(database)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepo, rabbitMQ)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)

	// Configurar router
	r := gin.Default()
	r.POST("/payments", paymentHandler.ProcessPayment)
	r.GET("/payments/:id", paymentHandler.GetPayment)

	// Iniciar servidor
	r.Run(":8081")
}