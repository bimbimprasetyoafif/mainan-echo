package repository

import (
	"context"
	"database/sql"
	"github.com/coba/model"
)

type mysqlDB struct {
	db *sql.DB
}

func (m *mysqlDB) GetAll() ([]model.Users, error) {
	users := []model.Users{}

	rows, err := m.db.Query("SELECT name, email FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.Users
		if err := rows.Scan(&user.Name, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (m *mysqlDB) SaveUser(ctx context.Context, user model.Users) error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlDB) UpdateUser(ctx context.Context, user model.Users) error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlDB) GetByID(ctx context.Context, id int) (model.UsersCommon, error) {
	//TODO implement me
	panic("implement me")
}

func NewMysql(db *sql.DB) Database {
	return &mysqlDB{
		db: db,
	}
}
