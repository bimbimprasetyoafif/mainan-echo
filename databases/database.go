package databases

import (
	"database/sql"
	"fmt"
	"github.com/coba/config"
	"github.com/coba/model"
	m "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	cfg := config.Cfg

	alamat := cfg.DB_ADDRESS
	username := cfg.DB_USERNAME
	password := cfg.DB_PASSWORD
	dbName := cfg.DB_NAME
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, alamat, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: true,
	})
	if err != nil {
		panic(err)
	}

	DB = db

	DB.AutoMigrate(&model.Users{}, &model.ProfilePicture{})

}

func InitDatabaseSql() *sql.DB {
	cfg := m.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "192.168.1.71:3306",
		DBName: "orm_aja",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	return db

}
