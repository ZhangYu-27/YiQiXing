package admin

import (
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FollowController struct {
}
type FollowListMessage struct {
	PostMessage
	List *[]models.FollowUp
}

func (this FollowController) AddFllow(c *gin.Context) {
	follow := models.NewFollow()
	err := c.ShouldBind(follow)
	if err != nil {
		fmt.Println(err)
		c.JSONP(http.StatusOK, PostMessage{220, err.Error()})
		return
	}
	company, err := models.Company{}.GetCompanyInfo("id = ?", strconv.Itoa(follow.CompanyID))
	if company.ID == 0 {
		fmt.Println(company)
		c.JSONP(http.StatusOK, PostMessage{220, "目标公司不存在"})
		return
	}
	err = follow.AddFollow()
	if err != nil {
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	c.JSONP(http.StatusOK, PostMessage{200, "新增成功"})
}

func (this FollowController) GetFollowList(c *gin.Context) {
	companyId := c.Param("company_id")
	if companyId == "" {
		c.JSONP(http.StatusOK, PostMessage{220, "未接受到数据"})
		return
	}
	followList, err := models.FollowUp{}.GetFollowList("company_id", companyId)
	if err != nil {
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	c.JSONP(http.StatusOK, FollowListMessage{
		PostMessage: PostMessage{
			Code:    200,
			Message: "请求成功",
		},
		List: followList,
	})
}
