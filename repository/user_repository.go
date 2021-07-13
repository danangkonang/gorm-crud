package repository

import (
	"github.com/danangkonang/gin-rest-api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	FindUsers() []entity.User
}

type Db struct {
	conn *gorm.DB
}

func NewRepository() Repository {
	// fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	"",
	// 	"",
	// 	"",
	// 	"",
	// 	"",
	// )
	dsn := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error")
	}
	db.AutoMigrate(&entity.User{})
	return &Db{
		conn: db,
	}
}

func (d *Db) FindUsers() []entity.User {
	var users []entity.User
	d.conn.Find(&users)
	return users
}
