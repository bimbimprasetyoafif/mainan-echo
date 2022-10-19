package users

import (
	"github.com/coba/dto"
	"github.com/coba/repository"
)

type UserService interface {
	GetAllUser() ([]dto.DTOUser, error)
}

type user struct {
	uRepo repository.Database
}

func (u *user) GetAllUser() ([]dto.DTOUser, error) {
	res, err := u.uRepo.GetAll()
	if err != nil {
		return nil, err
	}

	dtoS := []dto.DTOUser{}

	for _, v := range res {
		dtoS = append(dtoS, dto.DTOUser{
			Name:  v.Name,
			Email: v.Email,
		})
	}

	return dtoS, nil
}

func NewUser(repo repository.Database) UserService {
	return &user{
		repo,
	}
}
