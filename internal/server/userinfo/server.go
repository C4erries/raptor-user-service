package server

import (
	"context"

	userpb "github.com/c4erries/raptor-proto/userpb"
	"github.com/c4erries/raptor-user-service/internal/model"
	"google.golang.org/grpc"
)

type serverAPI struct {
	userpb.UnimplementedUserServiceServer
	service Service
}

type Service interface {
	GetUser(
		ctx context.Context,
		userid int64,
	) (*model.User, error)

	CreateUser(
		ctx context.Context,
		name string,
		email string,
	) (userid int64, err error)

	UpdateUser(
		ctx context.Context,
		userid int64,
		name string,
		email string,
	) error

	DeleteUser(
		ctx context.Context,
		userid int64,
	) error
}

// реализовать
func Start() {

}

func Register(grpcServer *grpc.Server, service Service) {
	userpb.RegisterUserServiceServer(grpcServer, &serverAPI{service: service})
}

/* реализовать
func (s *serverAPI) GetUser(ctx context.Context, in *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	db, err := storage()
	if err != nil {
		return nil, err
	}
	var name, email string
	var query string = "SELECT name, email FROM users WHERE userid=$1"
	row := db.QueryRow(query, in.Userid)
	err = row.Scan(name, email)
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserResponse{Userid: in.Userid, Name: name, Email: email}, nil
}
*/
// реализовать
func (s *serverAPI) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	panic("not implemented")
	//return &userpb.CreateUserResponse{Userid: "example-id-dfsjfdsd"}, nil
}

// реализовать
func (s *serverAPI) UpdateUser(ctx context.Context, in *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	panic("not implemented")
	//return &userpb.UpdateUserResponse{Success: true}, nil
}

// реализовать
func (s *serverAPI) DeleteUser(ctx context.Context, in *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	panic("not implemented")
	//return &userpb.DeleteUserResponse{Success: true}, nil
}
