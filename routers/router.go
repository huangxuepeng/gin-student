package routers

import (
	"student/api"
	"student/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	// UserRouter := Router.Group("user").Use(middleware.JWTAuth()) //参与用户操作的, 全部需要需要使用鉴权
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("/", api.Test)                                                //测试
		UserRouter.POST("login", api.PasswordLogin)                                  //登录
		UserRouter.GET("list", api.GetUserList)                                      //获取学生的列表
		UserRouter.POST("add", api.AddUser)                                          //新增
		UserRouter.DELETE("delete/:id", middleware.AuthMiddleware(), api.DeleteUser) //删除指定的学生
		// UserRouter.DELETE("delete/:id", api.DeleteUser)
		UserRouter.POST("update/:id", middleware.AuthMiddleware(), api.UpdateUser) //更新指定的学生
		UserRouter.POST("getuser", api.GetUser)                                    //得到指定学生的信息
	}
}
