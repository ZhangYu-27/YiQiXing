package admin

import (
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ArticleController struct {
}

type ArticleInfoMessage struct {
	PostMessage
	Info *models.Article
}

func (this ArticleController) AddArticle(c *gin.Context) {
	article := &models.Article{}
	err := c.ShouldBind(article)
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{220, err.Error()})
		return
	}
	article.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	err = article.AddArticle()
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	c.JSONP(http.StatusOK, PostMessage{200, "新增成功"})
}
func (this ArticleController) UpdateArticle(c *gin.Context) {
	article := &models.UpdateArticle{}
	err := c.ShouldBind(article)
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{220, err.Error()})
		return
	}
	article.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	err = article.UpdateArticle()
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	c.JSONP(http.StatusOK, PostMessage{200, "修改成功"})
}
func (this ArticleController) GetArticle(c *gin.Context) {
	articleId := c.Param("article_id")
	articleInfo, err := models.Article{}.GetArticle("id = ?", articleId)
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{500, err.Error()})
		return
	}
	c.JSONP(http.StatusOK, ArticleInfoMessage{
		PostMessage: PostMessage{200, "请求成功"},
		Info:        articleInfo,
	})

}
