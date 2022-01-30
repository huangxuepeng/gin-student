package service

import (
	"fmt"
	"student/dao"
	"student/model"
)

type UserLogin struct {
	StuNum   string `form:"stu_num" json:"stu_num" binding:"require"`
	Password string `form:"password" json:"password" binding:"require"`
}

func Login() model.Student {
	var service UserLogin
	var user model.Student
	if err := dao.DB.Where("stu_num = ?", service.StuNum).First(&user).Error; err != nil {
		fmt.Println(user)
	}
	return user
}
