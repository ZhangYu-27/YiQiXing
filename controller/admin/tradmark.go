package admin

import (
	"awesomeProject/models"
	"awesomeProject/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TradmarkController struct {
}
type TradmarkListMessage struct {
	PostMessage
	List *[]models.Tradmark
}
type TradmarkInfo struct {
	PostMessage
	Info *models.Tradmark
}

func (this TradmarkController) AddTradmark(c *gin.Context) {
	tradmark := &models.Tradmark{}
	err := c.ShouldBind(tradmark)
	if err != nil {
		fmt.Println(err)
		c.JSONP(http.StatusOK, PostMessage{220, err.Error()})
		return
	}
	tradmark.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	err = tradmark.AddTradmark()
	if err != nil {
		fmt.Println(err)
		c.JSONP(http.StatusOK, PostMessage{500, err.Error()})
		return
	}
	c.JSONP(http.StatusOK, PostMessage{200, "添加成功"})
}
func (this TradmarkController) GetTradmarkList(c *gin.Context) {
	tradmark := models.NewTradmark()
	db := tool.DB
	if c.Param("company_name") != "0" {
		companyName := c.Param("company_name")
		company, err := models.Company{}.GetCompanyInfo("name LIKE ?", "%%"+companyName+"%%")
		if err != nil || company.ID == 0 {
			fmt.Println(company)
			c.JSONP(http.StatusOK, PostMessage{220, "用户不存在"})
			fmt.Println("", err)
			return
		}
		db.Where("company_id", company.ID)
	}
	if c.Param("company_id") != "0" {
		db.Where("company_id", c.Param("company_id"))
	}
	fmt.Println(c.Param("tradmark"))
	if c.Param("tradmark") != "0" {
		db.Where("tradmark LIKE", "%%"+c.Param("tradmark"))
	}
	tradmarkList, err := tradmark.GetTradmarkList(db)
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	c.JSONP(http.StatusOK, TradmarkListMessage{
		PostMessage: PostMessage{200, "请求成功"},
		List:        tradmarkList,
	})
}
func (this TradmarkController) UpdateTradmark(c *gin.Context) {
	updateTradmark := &models.UpdateTradmark{}
	err := c.ShouldBind(updateTradmark)
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{220, err.Error()})
		return
	}
	err = updateTradmark.UpdateTradmark()
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{500, "更新失败"})
		return
	}
	c.JSONP(http.StatusOK, PostMessage{http.StatusOK, "更新成功"})
}
func (this TradmarkController) GetTradmark(c *gin.Context) {
	tradmarkId := c.Param("tradmark_id")
	tradmarkInfo, err := models.Tradmark{}.GetTradmarkInfo("id = ?", tradmarkId)
	if err != nil {
		fmt.Println(err)
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	c.JSONP(http.StatusOK, TradmarkInfo{
		PostMessage: PostMessage{200, "请求成功"},
		Info:        tradmarkInfo,
	})

}
