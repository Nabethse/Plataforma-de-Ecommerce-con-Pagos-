package handler

import (
	"net/http"
	"order-api/internal/order/usecase"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase *usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{orderUseCase}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var request struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderUseCase.CreateOrder(request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la orden"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.orderUseCase.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden no encontrada"})
		return
	}

	c.JSON(http.StatusOK, order)
}