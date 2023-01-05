package service

import (
	"clean-arsitecture/internal/domain"
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
