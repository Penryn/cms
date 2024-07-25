package service

import "cms/internal/model"

func GetUserByPhoneNum(phoneNum string) (*model.User,error){
	return d.GetUserByphoneNum(ctx,phoneNum)
}

func CreateUser(username string,sex string,phone string,major string,password string) (error){
	return d.CreateUser(ctx,&model.User{
		Username: username,
		Sex: sex,
		PhoneNum: phone,
		Major: major,
		Password: password,
	})
}