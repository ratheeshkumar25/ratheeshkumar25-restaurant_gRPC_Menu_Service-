package di

import (
	"github.com/ratheeshkumar/restaurant_menu_serviceV1/config"
	"github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/db"
	"github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/handlers"
	"github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/repositories"
	"github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/server"
	"github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/services"
)

func Init() {
	config.LoadConfig()
	db := db.ConnectDB()

	menuRepo := repositories.NewMenuRepo(db)
	menuSvc := services.NewMenuService(menuRepo)
	menuHandle := handlers.NewMenuHandler(menuSvc)
	server.NewGrpcServer(menuHandle)
}
