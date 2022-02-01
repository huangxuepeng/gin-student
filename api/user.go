package api

import (
	"fmt"
	"net/http"
	"student/dao"
	"student/middleware"
	"student/model"
	"student/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//实现优雅的分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

//获取所有的信息(完成)
func GetUserList(ctx *gin.Context) {
	var students []model.Student
	cont := ctx.Query("query")
	if cont == "" {
		result := dao.DB.Find(&students)
		if result.Error != nil {
			zap.S().Panic("GetUserList 数据库查找出现错误")
		}
	} else {
		// var student model.Student
		dao.DB.Where("stu_num LIKE ? ", fmt.Sprint("%"+cont+"%")).Find(&students)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": gin.H{
			"studentList": students,
		},
	})
}

//test
func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "你好好",
	})
}

//登录(完成)
func PasswordLogin(ctx *gin.Context) {
	var user model.Student
	//拿到输入的账号和密码
	mobiles := ctx.PostForm("mobile")
	passwords := ctx.PostForm("password")
	fmt.Println(mobiles, passwords)
	if mobiles == "" {
		fmt.Println("手机号码不能为空")
		return
	}
	if !response.IsMobile(mobiles) {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 402,
			"msg":  "手机号码错误",
		})
		return
	}
	dao.DB.Where("mobile=?", mobiles).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"msg": "该手机号码未完成注册, 请完成注册",
		})
		return
	}
	if passwords != user.Password {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码错误",
		})
		return
	}
	dd, err := middleware.ReleaseToken(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"user":  user,
		"token": dd,
	})

}

//删除指定的学生
func DeleteUser(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, user)
}

//新增学生 (完成)
func AddUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "新增学生",
	})
	//从前端获取数据
	mobile := ctx.PostForm("Mobile")
	password := ctx.PostForm("Password")
	name := ctx.PostForm("Name ")
	stuNum := ctx.PostForm("StuNum ")
	gender := ctx.PostForm("Gender")
	zap.S().Infof("前端获取数据成功: "+
		mobile, password, name, stuNum, gender,
	)

	//判断手机号码是否已经被注册
	if response.IsMobile(mobile) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "手机号码已被注册",
		})
		return
	}
	//判断学号是不是已经被注册
	if response.IsStuNum(stuNum) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "该学号已被注册",
		})
		return
	}
	//判断手机号码的合法性
	if !response.ValidateMobile(mobile) {
		zap.S().Infof("手机号码不合法!")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg":   "手机号码不合法",
			"error": mobile,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "完成注册",
	})
}

//更新指定的学生
func UpdateUser(ctx *gin.Context) {
	var user model.Student
	//在主页更新
	stu_num := ctx.PostForm("StuNum")
	mobile := ctx.PostForm("Mobile")
	name := ctx.PostForm("Name")
	gender := ctx.PostForm("Gender")
	fmt.Println(stu_num, mobile, name, gender)
	dao.DB.Where("stu_num=?", stu_num).First(&user)

	result := dao.DB.Model(&user).Updates(map[string]interface{}{"Mobile": mobile, "Name": name, "Gender": gender})
	if result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

//得到指定学生的信息
func GetUser(ctx *gin.Context) {
	arr := []model.Student{}
	var user model.Student
	stuNum := ctx.PostForm("stu_num")
	dao.DB.Where("stu_num=?", stuNum).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "没有符合要求的人",
			"data": gin.H{
				"user": arr,
			},
		})
		return
	}
	arr = append(arr, user)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"user": arr,
		},
	})

}
