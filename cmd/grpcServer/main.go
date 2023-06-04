package main

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vyctor/fc-grpc/internal/database"
	"github.com/vyctor/fc-grpc/internal/pb"
	"github.com/vyctor/fc-grpc/internal/service"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("sqlite3", ".db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)
	grpcServer := grpc.NewServer()

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	list, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	grpcServer.Serve(list)
}
