package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ryonryon/Go-Practice/src/interfaces/handler"
	"time"
)

// gorm.Modelの定義
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	gorm.Model
	Name string `json:"name"`
}

//func InitializeUser(db *gorm.DB) {
//	db.AutoMigrate(User{})
//}

func (u *User) IndexUser() (*[]User, error) {
	db := handler.CreateConnection()
	user := &[]User{}
	db.Find(user)
	if cap(*user) == 0 {
		return nil, fmt.Errorf("cannot find")
	}
	return user, nil
}

func (u *User) IndexOneUser(id string) (*[]User, error) {
	db := handler.CreateConnection()
	user := &[]User{}
	db.Where("id = ?", id).Find(user)
	if id == "" {
		return nil, fmt.Errorf("parameter cannot find")
	}
	if cap(*user) == 0 {
		return nil, fmt.Errorf("id: %s is not exist", id)
	}
	return user, nil
}

func (u *User) CreateUser(c *gin.Context) error {
	db := handler.CreateConnection()
	user := new(User)
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	db.Create(user)
	return nil
}

func (u *User) UpdateUser(c *gin.Context, id string) error {
	db := handler.CreateConnection()
	user := new(User)
	db.Where("id = ?", id).Find(user)
	if user.ID == 0 {
		return fmt.Errorf("id: %s is not exist", id)
	}
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	db.Save(&user)
	return nil
}

func (u *User) DestroyUser(id string) error {
	db := handler.CreateConnection()
	user := &[]User{}
	if id == "" {
		return fmt.Errorf("parameter cannot find")
	}
	db.Where("id = ?", id).Find(user)
	if cap(*user) == 0 {
		return fmt.Errorf("id: %s is not exist", id)
	}
	db.Where("id = ?", id).Delete(user)
	return nil
}