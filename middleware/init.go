package middleware

import (
	"awesomeProject/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//全局的中间件
func InitMiddleware(c *gin.Context) {
	url := c.Request.URL.String()
	if strings.Contains(url, "admin") {
		urlMap := map[string]bool{
			"/admin/login":    true,
			"/admin/login_in": true,
		}
		fmt.Println(url)
		if urlMap[url] != true {
			adminId := tool.GetSession("adminId", c)
			if adminId == nil {
				c.Redirect(http.StatusFound, "/admin/login")
			}
		}
	}
}
