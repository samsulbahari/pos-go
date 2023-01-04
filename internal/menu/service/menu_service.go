package service

import (
	"clean-arsitecture/internal/domain"
)

type MenuService struct {
	menuRepo domain.MenuRepository
}

func NewMenuService(mr domain.MenuRepository) *MenuService {
	return &MenuService{menuRepo: mr}
}

func (ms *MenuService) GetMenu(roleid float64) (int, []domain.ResultMenu, error) {
	resp, err := ms.menuRepo.GetDataMenu(roleid)
	if err != nil {
		return 500, nil, domain.ErrFailedGetData
	}

	result := make([]domain.ResultMenu, 0)

	for _, getdata := range resp {
		submenu, _ := ms.menuRepo.GetSubmenu(getdata.ID)
		if submenu == nil {
			result = append(result, domain.ResultMenu{
				ID:      getdata.ID,
				Name:    getdata.Name,
				Submenu: []domain.Submenu{},
			})
		} else {
			result = append(result, domain.ResultMenu{
				ID:      getdata.ID,
				Name:    getdata.Name,
				Submenu: submenu,
			})
		}

	}
	return 200, result, nil
}
