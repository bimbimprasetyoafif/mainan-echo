package repository

import (
	"context"
	"github.com/coba/model"
)

type coacroach struct {
}

func (c *coacroach) GetAll() ([]model.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (c *coacroach) SaveUser(ctx context.Context, user model.Users) error {
	//TODO implement me
	panic("implement me")
}

func (c *coacroach) UpdateUser(ctx context.Context, user model.Users) error {
	//TODO implement me
	panic("implement me")
}

func (c *coacroach) GetByID(ctx context.Context, id int) (model.UsersCommon, error) {
	//TODO implement me
	panic("implement me")
}

func NewCochroach() Database {
	return &coacroach{}
}
