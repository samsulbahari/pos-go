package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MCategory struct {
	ID        int
	Name      string `json:"name" form:"name" binding:"required"`
	Image     string
	Status    string `form:"status" binding:"required,boolean"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}

type UpdateCategory struct {
	ID        int
	Name      string `json:"name" form:"name"`
	Image     string
	Status    string `form:"status" binding:"boolean"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}

type CategoryRepository interface {
	GetData(page int) ([]MCategory, error)
	GetDataById(id int) (MCategory, error)
	TotalData() (int64, error)
	CreateData(category *MCategory) (*MCategory, error)
	DeleteData(id int) error
	UpdateData(id int, category *UpdateCategory) error
}

type CategoryService interface {
	GetData(page int) (ResultCategory, error)
	GetDataById(id int) (MCategory, error)
	CreateData(ctx *gin.Context, category *MCategory) (int, error)
	DeleteData(ctx *gin.Context) (int, error)
	UpdateData(ctx *gin.Context, category *UpdateCategory) (int, error)
}

type ResultCategory struct {
	Total    int
	PerPage  int
	Page     int
	LastPage float64
	Data     []MCategory
}
