package userinfo

import (
	"context"
	"log/slog"

	"github.com/c4erries/raptor-user-service/internal/database/postgres"
	"github.com/c4erries/raptor-user-service/internal/model"
)

type Userinfo struct {
	log     *slog.Logger
	storage *postgres.Postgres
}

func New(storage *postgres.Postgres) *Userinfo {
	return &Userinfo{
		log:     slog.Default(),
		storage: storage,
	}
}

func (s *Userinfo) GetUser(ctx context.Context, userid int64) (*model.User, error) {
	db := s.storage.DB
	var name, email string
	var query string = "SELECT name, email FROM users WHERE userid=$1"
	row := db.QueryRow(query, userid)

	if err := row.Scan(name, email); err != nil {
		return nil, err
	}

	return &model.User{Id: userid, Name: name, Email: email}, nil
}

// реализовать
func (s *Userinfo) CreateUser(ctx context.Context, name string, email string) (userid int64, err error) {
	//panic("not implemented")
	userid = int64(13213)
	return userid, nil
}

// реализовать
func (s *Userinfo) UpdateUser(ctx context.Context, userid int64, name string, email string) error {
	return nil
}

// реализовать
func (s *Userinfo) DeleteUser(ctx context.Context, userid int64) error {
	return nil
}
