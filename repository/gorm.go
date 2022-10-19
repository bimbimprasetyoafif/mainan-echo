package repository

import (
	"context"
	"github.com/coba/model"
	"gorm.io/gorm"
)

type gormSql struct {
	db *gorm.DB
}

func (m *gormSql) GetAll() ([]model.Users, error) {
	u := []model.Users{}

	err := m.db.Model(&model.Users{}).Find(&u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (m *gormSql) SaveUser(ctx context.Context, user model.Users) error {
	return m.db.Create(&user).Error
}

func (m *gormSql) UpdateUser(ctx context.Context, user model.Users) error {
	return m.db.Model(&model.Users{}).
		Where("id = ?", user.ID).
		Updates(&model.Users{
			Name:  user.Name,
			Email: user.Email}).Error
}

func (m *gormSql) GetByID(ctx context.Context, id int) (model.UsersCommon, error) {
	panic("")
}

func NewGorm(db *gorm.DB) Database {
	return &gormSql{
		db: db,
	}
}
