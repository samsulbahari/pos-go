package service

import (
	"clean-arsitecture/internal/domain"
	"clean-arsitecture/internal/libraries"
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type CategoryService struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryService(cr domain.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: cr,
	}
}

func (cs *CategoryService) GetData(page int) (domain.ResultCategory, error) {
	var Result domain.ResultCategory
	offset := (page - 1) * 10
	data, err := cs.categoryRepo.GetData(offset)
	if err != nil {
		return Result, domain.ErrFailedGetData
	}
	count, err := cs.categoryRepo.TotalData()

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
func (cs *CategoryService) GetDataById(id int) (domain.MCategory, error) {
	res, err := cs.categoryRepo.GetDataById(id)
	if err != nil {
		return domain.MCategory{}, domain.ErrNotFound
	}
	return res, nil
}
func (cs *CategoryService) CreateData(ctx *gin.Context, category *domain.MCategory) (int, error) {

	file, multipartheader, _ := ctx.Request.FormFile("image")

	if file == nil {
		return 422, errors.New("image cant null")
	}

	imageExtension := strings.Split(multipartheader.Filename, ".")

	if imageExtension[1] != "jpg" && imageExtension[1] != "jpeg" && imageExtension[1] != "png" {

		return 422, errors.New("image musk jpg,jpeg,png")
	}
	fileName := libraries.RandomString()
	cld, _ := libraries.Setupcloudinary()
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: string(fileName),
		Folder:   "pos"})

	if err != nil {
		return 500, domain.ErrInternalServerError
	}
	category.Image = resp.SecureURL
	_, err = cs.categoryRepo.CreateData(category)
	if err != nil {
		return 500, domain.ErrInternalServerError
	}

	return 200, nil
}
func (cs *CategoryService) DeleteData(ctx *gin.Context) (int, error) {
	pageParam := ctx.Query("id")
	id, err := strconv.Atoi(pageParam)
	if err != nil {
		return 422, domain.ErrFailedInputId
	}

	_, err = cs.categoryRepo.GetDataById(id)
	if err != nil {
		return 404, domain.ErrNotFound
	}

	err = cs.categoryRepo.DeleteData(id)
	if err != nil {
		return 500, domain.ErrInternalServerError
	}
	return 200, nil
}
func (cs *CategoryService) UpdateData(ctx *gin.Context, category *domain.UpdateCategory) (int, error) {
	pageParam := ctx.Query("id")
	id, err := strconv.Atoi(pageParam)
	if err != nil {
		return 422, domain.ErrFailedInputId
	}
	data, err := cs.categoryRepo.GetDataById(id)
	if err != nil {
		return 404, domain.ErrNotFound
	}

	file, multipartheader, _ := ctx.Request.FormFile("image")

	if file == nil {
		//update with no replace image
		err := cs.categoryRepo.UpdateData(id, category)
		if err != nil {
			return 500, domain.ErrInternalServerError
		}
		return 200, nil
	} else {
		imageExtension := strings.Split(multipartheader.Filename, ".")
		if imageExtension[1] != "jpg" && imageExtension[1] != "jpeg" && imageExtension[1] != "png" {

			return 422, errors.New("image musk jpg,jpeg,png")
		}
		if data.Image == "" {
			cld, _ := libraries.Setupcloudinary()
			randstring := libraries.RandomString()

			//upload new imagae
			resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
				PublicID: string(randstring),
				Folder:   "pos"})

			if err != nil {
				return 500, domain.ErrInternalServerError
			}

			category.Image = resp.SecureURL

			err = cs.categoryRepo.UpdateData(id, category)
			if err != nil {
				return 500, domain.ErrInternalServerError
			}
			return 200, nil
		}
		getpublicID := strings.Split(data.Image, "/")
		removeExtension := strings.Split(getpublicID[8], ".")

		cld, _ := libraries.Setupcloudinary()
		_, err := cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
			PublicIDs: []string{"pos/" + removeExtension[0]},
		})

		randstring := libraries.RandomString()

		//upload new imagae
		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
			PublicID: string(randstring),
			Folder:   "pos"})

		if err != nil {
			return 500, domain.ErrInternalServerError
		}

		category.Image = resp.SecureURL

		err = cs.categoryRepo.UpdateData(id, category)
		if err != nil {
			return 500, domain.ErrInternalServerError
		}
		return 200, nil

	}

}
