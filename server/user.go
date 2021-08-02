package server

import (
	"context"
	"fmt"
	"grpc-tutorial/db"
	"grpc-tutorial/services/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type UserServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (u *UserServiceServer) CreateUser(ctx context.Context, in *user.NewUser) (*user.User, error) {
	res, err := db.GetMongoDB().Users.InsertOne(
		ctx,
		in,
	)
	if err != nil {
		return nil, err
	}
	return &user.User{
		Id:   res.InsertedID.(primitive.ObjectID).Hex(),
		Name: in.GetName(),
		Age:  in.GetAge(),
	}, nil
}

func (u *UserServiceServer) DeleteUser(ctx context.Context, in *user.UserID) (*user.StandardResponse, error) {
	uid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}
	_, err = db.GetMongoDB().Users.DeleteOne(
		ctx,
		bson.M{"_id": uid},
	)
	if err != nil {
		return nil, err
	}
	return &user.StandardResponse{
		Success: true,
		Message: fmt.Sprintf("successfully delete user with id: %s", in.Id),
	}, nil
}
