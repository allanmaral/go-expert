package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/allanmaral/go-expert/14-grpc/internal/database"
	"github.com/allanmaral/go-expert/14-grpc/internal/pb"
	"github.com/allanmaral/go-expert/14-grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open connection to the database: %v", err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryServer := service.NewCategoryService(categoryDB)

	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, categoryServer)
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen to the tpc port 50051: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failes to server gRPC server: %v", err)
	}
}
