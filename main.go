package main

import (
	"awesomeProject/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("kkkkkkk111")) //创建一个基于cookie的存储引擎 传入数据为密钥
	r.Use(sessions.Sessions("myssession", store))
	routers.AdminRoutersInit(r)
	//adminRouters.GET("/get-conpanys/:page", admin.AdminController{}.GetUserList)
	//r.GET("", func(c *gin.Context) {
	//	name := c.Param("page")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})
	r.Run("0.0.0.0:8800")
}
