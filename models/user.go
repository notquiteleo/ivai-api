package models

import (
	"errors"

	"gorm.io/gorm"
)

type Users struct {
	ID                int    `json:"id"`
	UID               string `json:"uid"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Mobile            string `json:"mobile"`
	Wechat            string `json:"wechat"`
	Authentication_id int    `json:"authentication_id"`
}

func GetUserByMobile(mobile string) (*Users, error) {
	var user *Users
	err := DB.Where("mobile = ?", mobile).First(&user).Error
	return user, err
}

func FindOrCreateUser(item *Users) (*Users, error) {
	var user *Users
	err := DB.Where("mobile = ?", item.Mobile).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = DB.Create(&item).Error
		if err != nil {
			return nil, err
		}
		return item, nil
	}
	return user, nil
}

func GetUserByConditions(conditions map[string]interface{}) (*Users, error) {
	var user *Users
	err := DB.Where(conditions).First(&user).Error
	return user, err
}