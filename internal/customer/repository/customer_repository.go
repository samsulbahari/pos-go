package repository

import (
	"clean-arsitecture/internal/domain"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db}
}

func (m *CustomerRepository) GetData(offset int) ([]domain.MCustomer, error) {

	var customer []domain.MCustomer
	err := m.db.Limit(10).Offset(offset).Find(&customer).Error
	return customer, err
}

func (m *CustomerRepository) GetDataById(id int) (domain.MCustomer, error) {
	var customer domain.MCustomer
	err := m.db.Where("id = ?", id).First(&customer).Error
	return customer, err
}

func (m *CustomerRepository) GetDataByEmail(email string) (domain.MCustomer, error) {
	var customer domain.MCustomer
	err := m.db.Where("email = ?", email).First(&customer).Error
	return customer, err
}

func (m *CustomerRepository) TotalData() (int64, error) {
	var count int64
	var customer []domain.MCustomer
	err := m.db.Model(&customer).Count(&count).Error

	return count, err
}
func (m *CustomerRepository) CreateData(customer *domain.MCustomer) (*domain.MCustomer, error) {
	err := m.db.Create(&customer).Error
	return customer, err
}
func (m *CustomerRepository) DeleteData(id int) error {
	var customer domain.MCustomer
	err := m.db.Where("id = ?", id).Delete(&customer).Error
	return err
}

func (m *CustomerRepository) UpdateData(id int, customer *domain.UpdateCustomer) error {
	err := m.db.Table("m_customer").Where("id = ?", id).Updates(customer).Error
	return err
}
