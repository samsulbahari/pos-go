package handler

import (
	"clean-arsitecture/internal/domain"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService domain.CustomerService
}

func NewCustomerHandler(r *gin.RouterGroup, cs domain.CustomerService) {
	handler := &CustomerHandler{customerService: cs}
	r.GET("customer", handler.GetData)
	r.GET("customer/:id", handler.GetDataById)
}

func (ch *CustomerHandler) GetData(ctx *gin.Context) {

	data, err := ch.customerService.GetData(ctx)
	if err != nil {
		ctx.JSON(500, domain.ResponseError{Message: err.Error()})
		return

	}
	ctx.JSON(200, gin.H{
		"message": "success get data",
		"data":    data,
	})
}
func (c *CustomerHandler) GetDataById(ctx *gin.Context) {
	code, data, err := c.customerService.GetDataById(ctx)
	if err != nil {
		ctx.JSON(code, domain.ResponseError{Message: err.Error()})
		return

	}
	ctx.JSON(code, gin.H{
		"message": "success get data",
		"data":    data,
	})
}
