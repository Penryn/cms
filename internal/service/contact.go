package service

import "cms/internal/model"

func CreateContact(uid uint, sid string, name string, sex string, phone string, major string, blacklist bool) error {
	return d.CreateContact(ctx, &model.Contact{
		OwnerID:   uid,
		StudentID: sid,
		Name:      name,
		Sex:       sex,
		PhoneNum:  phone,
		Major:     major,
		Blacklist: blacklist,
	})
}

func UpdateContact(id uint, sid string, name string,sex string, phone string, major string, blacklist bool) error {
	return d.UpdateContact(ctx, id, sid, name,sex, phone, major, blacklist)
}

func DeleteContact(id uint) error {
	return d.DeleteContact(ctx, id)
}

func GetContactList(uid uint) ([]model.Contact, error) {
	return d.GetContactByOwnerID(ctx, uid)
}