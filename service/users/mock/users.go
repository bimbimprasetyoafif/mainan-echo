package mock

import (
	"github.com/coba/dto"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (u *UserMock) GetAllUser() ([]dto.DTOUser, error) {
	args := u.Called()

	return args.Get(0).([]dto.DTOUser), args.Error(1)
}
