package handlers

import (
	"context"

	pb "github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/pb"
	service_interfaces "github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/services/interfaces"
)

type MenuHandlers struct {
	pb.MenuServiceServer
	svc service_interfaces.MenuServiceInter
}

func NewMenuHandler(svc service_interfaces.MenuServiceInter) *MenuHandlers {
	return &MenuHandlers{
		svc: svc,
	}
}

func (m *MenuHandlers) CreateMenu(ctx context.Context, p *pb.MenuItem) (*pb.MenuResponse, error) {
	result, err := m.svc.CreateBookService(p)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (m *MenuHandlers) FetchByMenuID(ctx context.Context, p *pb.MenuID) (*pb.MenuItem, error) {
	result, err := m.svc.FetchByMenuID(p)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (m *MenuHandlers) FetchByName(ctx context.Context, p *pb.FoodByName) (*pb.MenuItem, error) {
	result, err := m.svc.FetchByName(p)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (m *MenuHandlers) FetchByMenus(ctx context.Context, p *pb.NoParam) (*pb.MenuList, error) {
	result, err := m.svc.FetchMenus(p)
	if err != nil {
		return nil, err
	}
	return result, err
}
