package response

import (
	"regexp"
	"student/dao"
	"student/model"

	"go.uber.org/zap"
)

//正则表达式判断手机号码是否合法
func ValidateMobile(mobile string) bool {
	//使用正则表达式判断是否合法
	ok, err := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if err != nil {
		zap.S().Panic("手机号码验证失败!")
	}
	if !ok {
		return false
	}
	return true
}

//判断手机号码是否已经被使用
func IsMobile(mobile string) bool {
	var user model.Student
	dao.DB.Where("mobile=?", mobile).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

//判断学号是否被使用
func IsStuNum(stunum string) bool {
	var user model.Student
	dao.DB.Where("stunum=?", stunum).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}
