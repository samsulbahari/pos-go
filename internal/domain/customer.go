package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MCustomer struct {
	ID        int
	Name      string `json:"name" binding:"required,min=2"`
	Address   string `json:"address" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,max=13"`
	Status    string `json:"status"`
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}

type ResultCustomer struct {
	Total    int
	PerPage  int
	Page     int
	LastPage float64
	Data     []MCustomer
}

type CustomerRepository interface {
	GetDataByEmail(email string) (MCustomer, error)
	GetData(page int) ([]MCustomer, error)
	GetDataById(id int) (MCustomer, error)
	TotalData() (int64, error)
	CreateData(category *MCustomer) (*MCustomer, error)
	DeleteData(id int) error
	//UpdateData(id int, category *UpdateCategory) error
}

type CustomerService interface {
	GetData(ctx *gin.Context) (ResultCustomer, error)
	GetDataById(ctx *gin.Context) (int, MCustomer, error)
	CreateData(ctx *gin.Context, custemerRepo *MCustomer) (int, error)
	DeleteData(ctx *gin.Context) (int, error)
	// UpdateData(ctx *gin.Context, category *UpdateCategory) (int, error)
}
