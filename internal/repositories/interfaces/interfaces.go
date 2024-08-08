package interfaces

import "github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/model"

type MenuRepoInter interface {
	CreateMenu(Menu *model.Menu) error
	FetchByMenuID(id uint) (*model.Menu, error)
	EditMenu(Menu *model.Menu) error
	FetchByName(name string) (*model.Menu, error)
	DeleteMenu(id uint) error
	FetchMenus() ([]*model.Menu, error)
}
