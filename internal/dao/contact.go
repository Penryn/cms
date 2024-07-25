package dao

import (
	"cms/internal/model"
	"context"
)

func (d *Dao) CreateContact(ctx context.Context, contact *model.Contact) error {
	err := d.orm.WithContext(ctx).Create(&contact).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao)UpdateContact(ctx context.Context, id uint, sid string, name string,sex string, phone string, major string, blacklist bool) error {
	err := d.orm.WithContext(ctx).Model(&model.Contact{}).Where("id = ?", id).Updates(map[string]interface{}{
		"student_id": sid,
		"name": name,
		"sex": sex,
		"phone_num": phone,
		"major": major,
		"blacklist": blacklist,
	}).Error
	return err
}

func (d *Dao) DeleteContact(ctx context.Context, id uint) error {
	err := d.orm.WithContext(ctx).Where("id = ?", id).Delete(&model.Contact{}).Error
	return err
}

func (d *Dao) GetContactByOwnerID(ctx context.Context, uid uint) ([]model.Contact, error) {
	var contacts []model.Contact
	err := d.orm.WithContext(ctx).Where("owner_id = ?", uid).Find(&contacts).Error
	return contacts, err
}