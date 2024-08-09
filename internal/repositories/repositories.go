package repositories

import (
	"github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/model"
	"gorm.io/gorm"
)

// MenuRepo implements the interfaces.MenuRepoInter interface
type MenuRepo struct {
	DB *gorm.DB
}

// NewMenuRepo creates a new instance of MenuRepo
func NewMenuRepo(db *gorm.DB) *MenuRepo {
	return &MenuRepo{DB: db}
}

// CreateMenu inserts a new menu item into the database
func (u *MenuRepo) CreateMenu(menu *model.Menu) error {
	if err := u.DB.Create(menu).Error; err != nil {
		return err
	}
	return nil
}

// FetchByMenuID retrieves a menu item by its ID
func (u *MenuRepo) FetchByMenuID(id uint) (*model.Menu, error) {
	var menu model.Menu
	err := u.DB.First(&menu, id).Error
	return &menu, err
}

// FetchByName retrieves a menu item by its name
func (u *MenuRepo) FetchByName(name string) (*model.Menu, error) {
	var menu model.Menu
	err := u.DB.Where("name = ?", name).First(&menu).Error
	return &menu, err
}

// EditMenu updates an existing menu item
func (u *MenuRepo) EditMenu(menu *model.Menu) error {
	err := u.DB.Save(menu).Error
	return err
}

// DeleteMenu deletes a menu item by its ID
func (u *MenuRepo) DeleteMenu(id uint) error {
	err := u.DB.Delete(&model.Menu{}, id).Error
	return err
}

// FetchMenus retrieves all menu items
func (u *MenuRepo) FetchMenus() ([]*model.Menu, error) {
	var menus []*model.Menu
	err := u.DB.Find(&menus).Error
	return menus, err
}
