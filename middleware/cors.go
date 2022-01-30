package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func Cors(context *gin.Context) {
// 	method := context.Request.Method
// 	// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
// 	context.Header("Access-Control-Allow-Origin", "http://localhost:8080, http://localhost:8081/#/login")
// 	// context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
// 	// fmt.Println(context.GetHeader("Origin"))
// 	// 必须，设置服务器支持的所有跨域请求的方法
// 	context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
// 	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
// 	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token, Authorization, X-CSRF-Token, session")
// 	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
// 	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token, Content-Length, Access-Control-Allow-Origin")
// 	context.Header("Access-Control-Max-Age", "0")
// 	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
// 	context.Header("Access-Control-Allow-Credentials", "true")
// 	// 放行所有OPTIONS方法
// 	if method == "OPTIONS" {
// 		context.AbortWithStatus(http.StatusNoContent)
// 		return
// 	}
// 	context.Next()
// }

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:8081"}
	config.AllowCredentials = true
	return cors.New(config)
}
