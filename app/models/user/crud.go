package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

func (user *User) Create() (err error) {
	if err := model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
	}

	return nil
}

func Get(idstr string) (User, error) {
	id := types.StringToUint64(idstr)

	var user User

	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// All 获取所有用户数据
func All() ([]User, error) {
	var users []User
	if err := model.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
