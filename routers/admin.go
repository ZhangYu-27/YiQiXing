package routers

import (
	"awesomeProject/controller/admin"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/", middleware.InitMiddleware)
	{
		adminRouters.POST("login", admin.AdminController{}.Login)
		adminRouters.POST("add-company", admin.AdminController{}.AddCompany)
		adminRouters.GET("get-companys", admin.AdminController{}.GetUserList)
		adminRouters.GET("get-companys/page/:page/num/:num", admin.AdminController{}.GetUserList)
		adminRouters.GET("get-company/id/:id", admin.AdminController{}.GetUserInfo)
		adminRouters.POST("add-follow", admin.FollowController{}.AddFllow)
		adminRouters.GET("get-follow-list/company_id/:company_id", admin.FollowController{}.GetFollowList)
		adminRouters.POST("add-tradmark", admin.TradmarkController{}.AddTradmark)
		adminRouters.GET("get-tradmarks/company_id/:company_id/company_name/:company_name/tradmark/:tradmark", admin.TradmarkController{}.GetTradmarkList)
		adminRouters.PUT("update-tradmark", admin.TradmarkController{}.UpdateTradmark)
	}

}
