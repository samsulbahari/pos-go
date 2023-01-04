package mocks

import (
	"clean-arsitecture/internal/domain"
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type CategoryServiceMock struct {
	Mock mock.Mock
}

func (_m *CategoryServiceMock) GetDataById(id int) (domain.MCategory, error) {
	ret := _m.Mock.Called(id)
	if ret.Get(0) != nil {
		r := ret.Get(0).(domain.MCategory)
		return domain.MCategory{ID: id, Name: r.Name}, nil
	} else {
		return domain.MCategory{}, errors.New("Data not found")

	}

}

func (_m *CategoryServiceMock) GetData(page int) ([]domain.MCategory, error) {
	ret := _m.Mock.Called(page)
	if ret.Get(0) != nil {

		r := ret.Get(0).([]domain.MCategory)
		return r, nil
	} else {
		fmt.Println("a")
		return nil, domain.ErrFailedGetData

	}

}

func (_m *CategoryServiceMock) TotalData() (int64, error) {
	ret := _m.Mock.Called()
	if ret.Get(0) != nil {
		return 100, nil
	} else {
		return 0, errors.New("error count data")

	}
}
