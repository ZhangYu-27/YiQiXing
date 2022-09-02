package routers

import (
	"awesomeProject/controller/admin"
	"awesomeProject/middleware"
	"awesomeProject/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/", middleware.InitMiddleware)
	{
		adminRouters.POST("login", admin.AdminController{}.Login)
		adminRouters.POST("add-company", admin.AdminController{}.AddCompany)
		adminRouters.GET("get-companys", admin.AdminController{}.GetUserList)
		adminRouters.GET("get-companys/page/:page/num/:num/is_vip/:is_vip", admin.AdminController{}.GetUserList)
		adminRouters.GET("get-company/id/:id", admin.AdminController{}.GetUserInfo)
		adminRouters.POST("add-follow", admin.FollowController{}.AddFllow)
		adminRouters.GET("get-follow-list/company_id/:company_id", admin.FollowController{}.GetFollowList)
		adminRouters.POST("add-tradmark", admin.TradmarkController{}.AddTradmark)
		adminRouters.GET("get-tradmarks/company_id/:company_id/company_name/:company_name/tradmark/:tradmark", admin.TradmarkController{}.GetTradmarkList)
		adminRouters.PUT("update-tradmark", admin.TradmarkController{}.UpdateTradmark)
		adminRouters.PUT("update-company", admin.AdminController{}.UpdateCompany)
		adminRouters.GET("get-tardmark/tardmark_id/:tradmark_id", admin.TradmarkController{}.GetTradmark)

		adminRouters.POST("add-article", admin.ArticleController{}.AddArticle)
		adminRouters.PUT("update-article", admin.ArticleController{}.UpdateArticle)
		adminRouters.GET("get-article/article_id/:article_id", admin.ArticleController{}.GetArticle)

		adminRouters.GET("get-file-list", admin.FileController{}.GetFileList)
		adminRouters.POST("upload-file", admin.FileController{}.UploadFile)
		adminRouters.GET("get-file/file_name/:file_name", admin.FileController{}.GetFile)
	}
}

func adminInit(c *gin.Context) {
	url := c.Request.URL.String()
	urlMap := map[string]bool{
		"/login": true,
	}
	fmt.Println(url)
	if urlMap[url] != true {
		adminId := tool.GetSession("adminId", c)
		if adminId == nil {
			c.Redirect(http.StatusFound, "/admin/login")
		}
	}
}
