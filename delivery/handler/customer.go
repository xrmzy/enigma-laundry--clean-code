package handler

import (
	"enigma-laundry-clean-code/model/dto"
	"enigma-laundry-clean-code/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	uc usecase.CustomerUseCase
	rg *gin.RouterGroup
}

func (cs *CustomerHandler) createdHandler(c *gin.Context) {
	var payload dto.CustomersRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	payloadResponse, err := cs.uc.CreatedCustomer(payload)
	if err != nil {
		dto.SendSingleResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	dto.SendSingleResponse(c, http.StatusCreated, "OK", payloadResponse)
}

func (cs *CustomerHandler) Route() {
	cs.rg.POST("/register", cs.createdHandler)
}

func NewCustomerHandler(uc usecase.CustomerUseCase, rg *gin.RouterGroup) *CustomerHandler {
	return &CustomerHandler{uc: uc, rg: rg}
}

func (cs *CustomerHandler) getByID(c *gin.Context) {
	id := c.Param("id")
	customer, err := cs.uc.FindByID(id)
	if err != nil {
		dto.SendSingleResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	dto.SendSingleResponse(c, http.StatusOK, "OK", customer)
}

func (cs *CustomerHandler) RouteGetByID() {
	cs.rg.GET("/customer/:id", cs.getByID)
}
