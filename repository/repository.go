package repository

import (
	"context"
	"github.com/coba/model"
)

type Database interface {
	GetAll() ([]model.Users, error)
	SaveUser(ctx context.Context, user model.Users) error
	UpdateUser(ctx context.Context, user model.Users) error
	GetByID(ctx context.Context, id int) (model.UsersCommon, error)
}
