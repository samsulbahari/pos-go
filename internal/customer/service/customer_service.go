package service

import (
	"clean-arsitecture/internal/domain"
	"errors"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerService struct {
	custemerRepo domain.CustomerRepository
}

func NewCustomerSerivce(cr domain.CustomerRepository) *CustomerService {
	return &CustomerService{custemerRepo: cr}
}

func (cs *CustomerService) GetData(ctx *gin.Context) (domain.ResultCustomer, error) {
	pageParam := ctx.Query("page")
	page, err := strconv.Atoi(pageParam)

	if err != nil {
		return domain.ResultCustomer{}, domain.ErrBadParamInput
	}

	var Result domain.ResultCustomer
	offset := (page - 1) * 10
	data, err := cs.custemerRepo.GetData(offset)
	if err != nil {
		return Result, domain.ErrFailedGetData
	}
	count, err := cs.custemerRepo.TotalData()

	if err != nil {
		return Result, domain.ErrFailedGetData
	}

	last_page_counts := float64(count) / float64(10)
	last_page := math.Ceil(last_page_counts)
	if last_page == 0 {
		last_page = 1
	}

	Result.Data = data
	Result.Page = page
	Result.PerPage = 10
	Result.Total = int(count)
	Result.LastPage = last_page

	return Result, nil
}
func (cs *CustomerService) GetDataById(ctx *gin.Context) (int, domain.MCustomer, error) {
	idstring := ctx.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {

		return 422, domain.MCustomer{}, domain.ErrFailedInputId
	}

	res, err := cs.custemerRepo.GetDataById(id)
	if err != nil {
		return 404, domain.MCustomer{}, domain.ErrNotFound
	}
	return 200, res, nil
}

func (cs *CustomerService) CreateData(ctx *gin.Context, custemerRepo *domain.MCustomer) (int, error) {
	_, err := cs.custemerRepo.GetDataByEmail(custemerRepo.Email)

	if err == nil {
		return 422, errors.New("email musk unique")
	}

	_, err = cs.custemerRepo.CreateData(custemerRepo)
	if err != nil {
		return 500, domain.ErrInternalServerError
	}

	return 200, nil
}

func (cs *CustomerService) DeleteData(ctx *gin.Context) (int, error) {
	pageParam := ctx.Query("id")
	id, err := strconv.Atoi(pageParam)
	if err != nil {
		return 422, domain.ErrFailedInputId
	}

	_, err = cs.custemerRepo.GetDataById(id)
	if err != nil {
		return 404, domain.ErrNotFound
	}

	err = cs.custemerRepo.DeleteData(id)
	if err != nil {
		return 500, domain.ErrInternalServerError
	}
	return 200, nil
}
func (cs *CustomerService) UpdateData(ctx *gin.Context, customer *domain.UpdateCustomer) (int, error) {
	pageParam := ctx.Query("id")
	id, err := strconv.Atoi(pageParam)
	if err != nil {
		return 422, domain.ErrFailedInputId
	}
	data, err := cs.custemerRepo.GetDataById(id)
	if err != nil {
		return 404, domain.ErrNotFound
	}
	if data.Email == customer.Email {
		err = cs.custemerRepo.UpdateData(id, customer)
		if err != nil {
			return 500, domain.ErrInternalServerError
		}
		return 200, nil
	}
	_, err = cs.custemerRepo.GetDataByEmail(data.Email)
	if err == nil {
		return 422, errors.New("email musk unique")
	}

	return 200, nil

}
