package domain

type MenuRepository interface {
	GetDataMenu(role float64) ([]Menu, error)
	GetSubmenu(idpermission int) ([]Submenu, error)
}

type MenuService interface {
	GetMenu(roleid float64) (int, []ResultMenu, error)
}

type Menu struct {
	ID   int
	Name string
}

type Submenu struct {
	Name string
}

type ResultMenu struct {
	ID      int
	Name    string
	Submenu []Submenu
}
