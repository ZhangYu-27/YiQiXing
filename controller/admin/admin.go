package admin

import (
	"awesomeProject/models"
	"awesomeProject/tool"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AdminController struct {
}
type PostMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CompanyListMessage struct {
	PostMessage
	List *[]models.Company `json:"list"`
}
type CompanyInfoMessage struct {
	PostMessage
	CompanyInfo *models.Company `json:"company_info"`
}

func (this AdminController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	adminInfo := models.AdminUser{}.GetAdminInfo("username = ?", username)
	//encrypt, err := tool.Encrypt([]byte(adminInfo.Key), []byte(password))
	//if err != nil {
	//	c.JSONP(http.StatusOK, LoginMessage{500, "系统错误"})
	//	return
	//}
	//fmt.Println(hex.EncodeToString(encrypt))
	//fmt.Println(adminInfo.Key)
	//fmt.Println(adminInfo.Password)
	passwordByet, err := hex.DecodeString(adminInfo.Password)
	if err != nil {
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	decrypt, err := tool.Decrypt([]byte(adminInfo.Key), passwordByet)
	if err != nil {
		fmt.Println(err)
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	fmt.Println(string(decrypt))
	if password == string(decrypt) {
		c.JSONP(http.StatusOK, PostMessage{200, "登录成功"})
	} else {
		c.JSONP(http.StatusOK, PostMessage{220, "密码错误"})
	}
}
func (this AdminController) AddCompany(c *gin.Context) {
	company := &models.Company{}
	company.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	err := c.ShouldBind(company)
	if err != nil {
		c.JSONP(http.StatusOK, PostMessage{250, err.Error()})
		fmt.Println(err)
		return
	}
	if err = company.AddCompany(); err == nil {
		c.JSONP(http.StatusOK, PostMessage{200, "新增成功"})
	} else {
		fmt.Println("-------------------------------------")
		fmt.Println(err)
		fmt.Println("-------------------------------------")
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误，新增失败"})
	}
}

func (this AdminController) GetUserList(c *gin.Context) {
	pageS := c.Param("page")
	numS := c.Param("num")
	if pageS == "" {
		pageS = "1"
	}
	if numS == "" {
		numS = "20"
	}

	list, err := models.Company{}.GetCompanyList(pageS, numS)
	if err != nil {
		fmt.Println(err)
		c.JSONP(http.StatusOK, PostMessage{
			Code:    500,
			Message: "系统错误，查询失败",
		})
		return
	}
	c.JSONP(http.StatusOK, CompanyListMessage{
		PostMessage: PostMessage{200, "请求成功"},
		List:        list,
	})

}
func (this AdminController) GetUserInfo(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSONP(http.StatusOK, PostMessage{
			Code:    220,
			Message: "请携带id请求",
		})
		return
	}
	userInfo, err := models.Company{}.GetCompanyInfo("id = ?", userId)
	if err != nil {
		fmt.Println(err)
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误，查询失败"})
		return
	}
	c.JSONP(http.StatusOK, CompanyInfoMessage{
		PostMessage: PostMessage{200, "请求成功"},
		CompanyInfo: userInfo,
	})
}
