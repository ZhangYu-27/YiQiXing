package routers

import (
	"awesomeProject/controller/admin"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/", middleware.InitMiddleware)
	{
		adminRouters.POST("/login", admin.AdminController{}.Login)
		adminRouters.POST("/add-company", admin.AdminController{}.AddCompany)
		adminRouters.GET("/get-companys", admin.AdminController{}.GetUserList)
		adminRouters.GET("/get-companys/page/:page/num/:num", admin.AdminController{}.GetUserList)

	}

}
