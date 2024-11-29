package repo

import (
	"context"
	"gorm.io/gorm"
)

type UserI interface {
	CreateUser(context.Context, *User) error
	GetUserByID(context.Context, int64) (User, error)
	UpdateUser(context.Context, *User, uint) error
	DeleteUser(context.Context, uint) error
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	IsActive bool   `gorm:"not null"`
}

//type UserService struct {
//	Store UserStore
//}
//
//func NewService(store UserStore) *UserService {
//	return &UserService{
//		Store: store,
//	}
//}
