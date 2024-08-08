package server

import (
	"log"
	"net"

	"github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/handlers"
	pb "github.com/ratheeshkumar/restaurant_menu_serviceV1/internal/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(handlr *handlers.MenuHandlers) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("error creating listener on port 8083")
	}

	grp := grpc.NewServer()
	pb.RegisterMenuServiceServer(grp, handlr)

	log.Printf("listening on gRPC server 8083")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")
	}
}
