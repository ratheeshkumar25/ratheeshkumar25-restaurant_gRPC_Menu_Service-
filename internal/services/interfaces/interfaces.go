package service_interfaces

import (
	pb "github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/pb"
)

type MenuServiceInter interface {
	CreateBookService(p *pb.MenuItem) (*pb.MenuResponse, error)
	FetchByMenuID(p *pb.MenuID) (*pb.MenuItem, error)
	FetchByName(p *pb.FoodByName) (*pb.MenuItem, error)
	FetchMenus(p *pb.NoParam) (*pb.MenuList, error)
}
