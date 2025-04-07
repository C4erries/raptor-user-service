package server

import (
	"context"

	userpb "github.com/c4erries/raptor-proto/userpb"
	"google.golang.org/grpc"
)

type apiServer struct {
	userpb.UnimplementedUserServiceServer
	server Server
}

type Server interface {
	GetUser(
		ctx context.Context,
		userid string,
	) (Userid string, name string, email string)

	CreateUser(
		ctx context.Context,
		name string,
		email string,
	) (userid string)

	UpdateUser(
		ctx context.Context,
		name string,
	) (ok bool)

	DeleteUser(
		ctx context.Context,
		userid string,
	) (ok bool)
}

// реализовать
func Start() {

}

func Register(grpcServer *grpc.Server, server Server) {
	userpb.RegisterUserServiceServer(grpcServer, &apiServer{server: server})
}

// реализовать
func (s *apiServer) GetUser(ctx context.Context, in *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{Userid: in.Userid, Name: "Say My Name", Email: "e@e.e"}, nil
}

// реализовать
func (s *apiServer) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return &userpb.CreateUserResponse{Userid: "example-id-dfsjfdsd"}, nil
}

// реализовать
func (s *apiServer) UpdateUser(ctx context.Context, in *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	return &userpb.UpdateUserResponse{Success: true}, nil
}

// реализовать
func (s *apiServer) DeleteUser(ctx context.Context, in *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	return &userpb.DeleteUserResponse{Success: true}, nil
}
