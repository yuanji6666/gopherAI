package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	

	enterRouter := r.Group("/api/v1")
	{
		//注册用户注册，登陆，验证码路由
		RegisterUserRouter(enterRouter.Group("/user"))
	}

	
	return r
}