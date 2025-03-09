package handler

import (
	"net/http"
	"payment-api/internal/payment/usecase"
	"payment-api/internal/payment/usecase"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUseCase *usecase.PaymentUseCase
}

func NewPaymentHandler(paymentUseCase *usecase.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{paymentUseCase}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var request struct {
		OrderID string  `json:"order_id"`
		Amount  float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.paymentUseCase.ProcessPayment(request.OrderID, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el pago"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	payment, err := h.paymentUseCase.GetPayment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pago no encontrado"})
		return
	}

	c.JSON(http.StatusOK, payment)
}