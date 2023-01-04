package handler

import (
	"clean-arsitecture/internal/domain"

	"github.com/gin-gonic/gin"
)

type MenuService struct {
	menuService domain.MenuService
}

func NewMenuHandler(r *gin.RouterGroup, ms domain.MenuService) {
	handler := &MenuService{
		menuService: ms,
	}
	r.GET("menu", handler.Getdata)
}

func (ms *MenuService) Getdata(ctx *gin.Context) {
	id, _ := ctx.Get("role")
	resp, data, err := ms.menuService.GetMenu(id.(float64))
	if err != nil {
		ctx.JSON(resp, domain.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(resp, gin.H{
		"data": data,
	})
}
