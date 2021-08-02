package main

import (
	"fmt"
	"grpc-tutorial/db"
	"grpc-tutorial/server"
	"grpc-tutorial/services/user"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	fmt.Println("init runs")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading env %v", err)
	}

	db.ConnectMongo()
}

func main() {
	// u := server.UserServiceServer{}
	// grpcServer := grpc.NewServer()
	// user.RegisterUserServiceServer(grpcServer, &u)

	// l, err := net.Listen("tcp", ":9000")
	// if err != nil {
	// 	log.Fatalf("could not listen to :9000 %v", err)
	// }

	// log.Fatal(grpcServer.Serve(l))

	u := server.UserServiceServer{}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &u)

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("could not listen to :9000 %v", err)
	}
	log.Fatal(grpcServer.Serve(l))
}
