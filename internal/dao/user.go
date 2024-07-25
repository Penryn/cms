package dao

import (
	"cms/internal/model"
	"context"
)

func (d *Dao) GetUserByphoneNum(ctx context.Context,phone string) (*model.User,error){
	var user model.User
	err := d.orm.WithContext(ctx).Where("phone_num = ?",phone).First(&user).Error
	if err != nil {
		return nil,err
	}
	return &user,nil
}

func (d *Dao) CreateUser(ctx context.Context,user *model.User) (error){
	err := d.orm.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}