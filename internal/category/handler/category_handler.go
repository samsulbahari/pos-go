package handler

import (
	"clean-arsitecture/internal/domain"
	"clean-arsitecture/internal/libraries"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CategoryHandler struct {
	categoryservice domain.CategoryService
}

func NewCategoryHandler(r *gin.RouterGroup, c domain.CategoryService) {
	handler := &CategoryHandler{
		categoryservice: c,
	}
	r.GET("/category", handler.GetData)
	r.GET("/category/:id", handler.GetDataById)
	r.POST("/category", handler.CreateData)
	r.DELETE("/category", handler.DeleteData)
	r.PUT("/category", handler.UpdateData)

}

func (c *CategoryHandler) GetData(ctx *gin.Context) {

	pageParam := ctx.Query("page")
	page, err := strconv.Atoi(pageParam)

	if err != nil {
		ctx.JSON(422, gin.H{
			"message": "Invalid input ID",
		})
		return
	}

	data, err := c.categoryservice.GetData(page)
	if err != nil {
		ctx.JSON(500, ResponseError{Message: err.Error()})
		return

	}
	ctx.JSON(200, gin.H{
		"message": "success get data",
		"data":    data,
	})
}
func (c *CategoryHandler) GetDataById(ctx *gin.Context) {
	idstring := ctx.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		ctx.JSON(422, ResponseError{Message: "id musk integer"})
		return
	}

	data, err := c.categoryservice.GetDataById(id)
	if err != nil {
		ctx.JSON(404, ResponseError{Message: err.Error()})
		return

	}
	ctx.JSON(200, gin.H{
		"message": "success get data",
		"data":    data,
	})
}
func (c *CategoryHandler) CreateData(ctx *gin.Context) {
	var category domain.MCategory

	err := ctx.ShouldBind(&category)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(422, gin.H{
			"message": validation_response,
		})
		return
	}
	resp, err := c.categoryservice.CreateData(ctx, &category)
	if err != nil {
		ctx.JSON(resp, ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(resp, gin.H{
		"message": "success insert data",
	})

}

func (c *CategoryHandler) DeleteData(ctx *gin.Context) {
	resp, err := c.categoryservice.DeleteData(ctx)
	if err != nil {
		ctx.JSON(resp, ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(resp, gin.H{
		"message": "success delete data",
	})
}

func (c *CategoryHandler) UpdateData(ctx *gin.Context) {
	var category domain.UpdateCategory

	err := ctx.ShouldBind(&category)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(422, gin.H{
			"message": validation_response,
		})
		return
	}

	resp, err := c.categoryservice.UpdateData(ctx, &category)
	if err != nil {
		ctx.JSON(resp, ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(resp, gin.H{
		"message": "success update data",
	})
}
