package service

import (
	"clean-arsitecture/internal/domain"
	"clean-arsitecture/internal/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var CategoryRepository = mocks.CategoryServiceMock{Mock: mock.Mock{}}
var categoriserv = NewCategoryService(&CategoryRepository)

func TestCategoryGetDataByid(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		category := domain.MCategory{
			ID:   3,
			Name: "samsul",
		}
		//ekspentasi return
		CategoryRepository.Mock.On("GetDataById", 3).Return(category)

		result, _ := categoriserv.GetDataById(3)
		assert.Equal(t, category.ID, result.ID)
	})

	t.Run("data not found", func(t *testing.T) {
		categorynil := domain.MCategory{}
		//ekspentasi return
		CategoryRepository.Mock.On("GetDataById", 1).Return(nil)

		result, _ := categoriserv.GetDataById(1)
		assert.Equal(t, categorynil, result)
	})
}
func TestCategoryGet(t *testing.T) {
	t.Run("get data success", func(t *testing.T) {
		category := domain.ResultCategory{
			Total:    100,
			PerPage:  10,
			Page:     1,
			LastPage: 10,
			Data: []domain.MCategory{
				{
					ID:   1,
					Name: "samsul",
				}, {
					ID:   2,
					Name: "samsul2",
				},
			}}
		CategoryRepository.Mock.On("GetData", 1).Return(category.Data)

		result, _ := categoriserv.categoryRepo.GetData(1)

		CategoryRepository.Mock.On("TotalData").Return(100)
		resultcount, _ := categoriserv.categoryRepo.TotalData()

		last_page_counts := float64(resultcount) / float64(10)

		assert.Equal(t, category.LastPage, last_page_counts)

		assert.Equal(t, category.Total, int(resultcount))
		//page harus sama dengan parameter
		assert.Equal(t, category.Page, 1)

		assert.Equal(t, category.Data, result)

		//max size array 10
		assert.LessOrEqual(t, len(category.Data), 10)

	})

}
