package main

import (
	"context"
	"fmt"
	"grpc-tutorial/services/user"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading env %v", err)
	}
}

func main() {
	conn, err := grpc.Dial(os.Getenv("PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't connect to grpc server: %v", err)
	}
	defer conn.Close()

	userService := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var users = make(map[string]int32)
	var ids = make(map[string]string)
	users["Alice"] = 43
	users["Bob"] = 30


	for name, age := range users {
		r, err := userService.CreateUser(ctx, &user.NewUser{
			Name: name,
			Age:  age,
		})
		if err != nil {
			log.Fatalf("Could not create new user :%v", err)
		}
		log.Printf(`User Details:
		Name: %s,
		AGE: %d,
		ID: %s
		`, r.GetName(), r.GetAge(), r.GetId())
		ids[name] = r.GetId()
	}

	resp, err := userService.DeleteUser(ctx, &user.UserID{Id: ids["Alice"]})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp, "<- resp")
}
