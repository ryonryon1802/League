package implements

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ryonryon/Go-Practice/src/interfaces/handler"
	"github.com/ryonryon/Go-Practice/src/models/db"
)

var (
	UserImplements IUserImplements
)

func init() {
	UserImplements = NewUserImplements(handler.DB)
}

type userImplements struct {
	*gorm.DB
}

type IUserImplements interface {
	SelectAllUser() ([]*db.User, error)
	SelectSingleUser(ctx *gin.Context) (*db.User, error)
	CreateUser(ctx *gin.Context) error
	UpdateUser(ctx *gin.Context) error
	DeleteUser(ctx *gin.Context) error
}

func NewUserImplements(db *gorm.DB) IUserImplements {
	return &userImplements{db}
}

func (impl *userImplements) SelectAllUser() ([]*db.User, error) {
	var users []*db.User
	err := impl.DB.Find(users).Error
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("cannot find")
	}
	return users, nil
}

func (impl *userImplements) SelectSingleUser(ctx *gin.Context) (*db.User, error) {
	var user *db.User
	id := ctx.Param("id")
	err := impl.DB.Where("id = ?", id).Find(user).Error
	if err != nil {
		return nil, err
	}
	if id == "" {
		return nil, fmt.Errorf("parameter cannot find")
	}
	if user == nil {
		return nil, fmt.Errorf("id: %s is not exist", id)
	}
	return user, nil
}

func (impl *userImplements) CreateUser(ctx *gin.Context) error {
	user := new(db.User)
	err := ctx.Bind(&user)
	if err != nil {
		return err
	}
	err = impl.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (impl *userImplements) UpdateUser(ctx *gin.Context) error {
	user := new(db.User)
	id := ctx.Param("id")
	if id == "" {
		return fmt.Errorf("parameter cannot find")
	}
	err := impl.DB.Where("id = ?", id).Find(user).Error
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return fmt.Errorf("id: %s is not exist", id)
	}
	err = ctx.Bind(&user)
	if err != nil {
		return err
	}
	err = impl.DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (impl *userImplements) DeleteUser(ctx *gin.Context) error {
	var users  []*db.User
	id := ctx.Param("id")
	if id == "" {
		return fmt.Errorf("parameter cannot find")
	}
	err := impl.DB.Where("id = ?", id).Find(users).Error
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return fmt.Errorf("id: %s is not exist", id)
	}
	err = impl.DB.Where("id = ?", id).Delete(users).Error
	if err != nil {
		return err
	}
	return nil
}