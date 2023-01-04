package repository

import (
	"clean-arsitecture/internal/domain"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (m *CategoryRepository) GetData(offset int) ([]domain.MCategory, error) {

	var category []domain.MCategory
	err := m.db.Limit(10).Offset(offset).Find(&category).Error
	return category, err
}

func (m *CategoryRepository) GetDataById(id int) (domain.MCategory, error) {
	var category domain.MCategory
	err := m.db.Where("id = ?", id).First(&category).Error
	return category, err
}

func (m *CategoryRepository) TotalData() (int64, error) {
	var count int64
	var category []domain.MCategory
	err := m.db.Model(&category).Count(&count).Error

	return count, err
}
func (m *CategoryRepository) CreateData(category *domain.MCategory) (*domain.MCategory, error) {
	err := m.db.Create(&category).Error
	return category, err
}
func (m *CategoryRepository) DeleteData(id int) error {
	var category domain.MCategory
	err := m.db.Where("id = ?", id).Delete(&category).Error
	return err
}
func (m *CategoryRepository) UpdateData(id int, category *domain.UpdateCategory) error {
	err := m.db.Table("m_category").Where("id = ?", id).Updates(category).Error
	return err
}
