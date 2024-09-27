package main

import (
	"database/sql"
	"fmt"
	"music_service/service"
	"net"
	"os"
	"github.com/gremislaw/music_service/api"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("DB connection error")
	}
	defer db.Close()

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
	if err != nil {
		fmt.Printf("listen error")
        panic("listen error")
    }
	api.RegisterMusicServiceServer(grpcServer, service.NewService(db))
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}