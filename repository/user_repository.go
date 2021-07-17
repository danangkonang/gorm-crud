package repository

import (
	"github.com/danangkonang/gin-rest-api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	FindUsers() []entity.User
	SaveUser(user entity.User)
	UpdateUser(user *entity.User)
	DeleteUser(user *entity.User)
	FindUserById(user entity.User) entity.User
}

type Db struct {
	conn *gorm.DB
}

func NewRepository() Repository {
	dsn := "danang:danang@tcp(localhost:3306)/belajar_gin?charset=utf8&parseTime=True&loc=Local"
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

func (d *Db) SaveUser(users entity.User) {
	d.conn.Create(&users)
}

func (d *Db) DeleteUser(users *entity.User) {
	d.conn.Delete(&users)
}

func (d *Db) UpdateUser(users *entity.User) {
	d.conn.Save(&users)
}

func (d *Db) FindUserById(users entity.User) entity.User {
	d.conn.Find(&users, users.Id)
	return users
}
