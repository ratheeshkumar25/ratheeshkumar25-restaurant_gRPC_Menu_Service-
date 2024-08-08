package services

import (
	"github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/model"
	pb "github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/pb"
	"github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/repositories/interfaces"
	service_interfaces "github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/services/interfaces"
)

type MenuService struct {
	repo interfaces.MenuRepoInter
}

// CreateBookService implements service_interfaces.MenuServiceInter.
func (m *MenuService) CreateBookService(p *pb.MenuItem) (*pb.MenuResponse, error) {
	menu := &model.Menu{
		Category:  p.Category,
		Name:      p.Name,
		Price:     p.Price,
		FoodImage: p.Foodimage,
		Duration:  p.Duration,
	}

	err := m.repo.CreateMenu(menu)
	if err != nil {
		return nil, err
	}

	return &pb.MenuResponse{
		Staus:   "Success",
		Error:   "",
		Message: "Item added successfully",
	}, nil
}

// FetchByMenuID implements service_interfaces.MenuServiceInter.
func (m *MenuService) FetchByMenuID(p *pb.MenuID) (*pb.MenuItem, error) {
	menu, err := m.repo.FetchByMenuID(uint(p.Id))
	if err != nil {
		return nil, err
	}
	return &pb.MenuItem{
		Id:        uint32(menu.ID),
		Category:  menu.Category,
		Name:      menu.Name,
		Price:     menu.Price,
		Foodimage: menu.FoodImage,
		Duration:  menu.Duration,
	}, nil
}

// FetchByName implements service_interfaces.MenuServiceInter.
func (m *MenuService) FetchByName(p *pb.FoodByName) (*pb.MenuItem, error) {
	menu, err := m.repo.FetchByName(p.Name)
	if err != nil {
		return nil, err
	}
	return &pb.MenuItem{
		Id:        uint32(menu.ID),
		Category:  menu.Category,
		Name:      menu.Name,
		Price:     menu.Price,
		Foodimage: menu.FoodImage,
		Duration:  menu.Duration,
	}, nil
}

// FetchMenus implements service_interfaces.MenuServiceInter.
func (m *MenuService) FetchMenus(p *pb.NoParam) (*pb.MenuList, error) {
	result, err := m.repo.FetchMenus()
	if err != nil {
		return nil, err
	}

	var menus []*pb.MenuItem
	for _, menu := range result {
		menus = append(menus, &pb.MenuItem{
			Id:        uint32(menu.ID),
			Category:  menu.Name,
			Name:      menu.Name,
			Price:     menu.Price,
			Foodimage: menu.FoodImage,
			Duration:  menu.Duration,
		})
	}
	return &pb.MenuList{
		Item: menus,
	}, nil
}

func NewMenuService(repo interfaces.MenuRepoInter) service_interfaces.MenuServiceInter {
	return &MenuService{
		repo: repo,
	}
}
