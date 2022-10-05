package databases

import (
	"github.com/coba/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:root@tcp(192.168.1.71:3306)/orm_aja?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: true,
	})
	if err != nil {
		panic(err)
	}

	DB = db

	DB.AutoMigrate(&model.Users{}, &model.ProfilePicture{})

}
