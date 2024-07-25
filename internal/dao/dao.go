package dao

import (
	"cms/internal/model"
	"context"

	"gorm.io/gorm"
)

type Dao struct {
	orm *gorm.DB
}

func New(orm *gorm.DB) *Dao {
	return &Dao{orm: orm}
}

type Daos interface{
	// User
	GetUserByUsername(ctx context.Context,username string) (*model.User,error)
	CreateUser(ctx context.Context,user *model.User) error
	// Contact
	CreateContact(ctx context.Context,contact *model.Contact) error
	UpdateContact(ctx context.Context,id uint,sid string,name string,sex string,phone string,major string,blacklist bool) error
	DeleteContact(ctx context.Context,id uint) error
	GetContactByOwnerID(ctx context.Context,uid uint) ([]model.Contact,error)

}