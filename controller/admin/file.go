package admin

import (
	"awesomeProject/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type FileController struct {
}

type FileMessage struct {
	PostMessage
	List *[]string `json:"list"`
}

func checkDir() error {
	_, err := os.Stat("/www/wwwroot/file/")
	if err != nil {
		err = os.Mkdir("/www/wwwroot/file/", 0777)
		if err != nil {
			return err
		}
	}
	_, err = os.Stat("/www/wwwroot/file/")
	if err != nil {
		return err
	}
	return nil
}

func (this FileController) UploadFile(c *gin.Context) {
	err := checkDir()
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{
			Code:    200,
			Message: "初始化失败",
		})
		return
	}
	err = tool.UploadImg(c)
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{500, "上传失败"})
		return
	}
	c.JSONP(http.StatusOK, PostMessage{200, "上传成功"})
}

func (this FileController) GetFileList(c *gin.Context) {

	dir, err := ioutil.ReadDir("/www/wwwroot/file/")
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	files := make([]string, 0)
	for _, file := range dir {
		files = append(files, file.Name())
	}
	c.JSONP(http.StatusOK, FileMessage{
		PostMessage: PostMessage{200, "请求成功"},
		List:        &files,
	})
}
func (this FileController) GetFile(c *gin.Context) {
	fileName := c.Param("file_name")
	_, err := os.Stat("/www/wwwroot/file/" + fileName)
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{500, "文件不存在"})
		return
	}
	file, err := os.OpenFile("/www/wwwroot/file/"+fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("", err)
		c.JSONP(http.StatusOK, PostMessage{500, "系统错误"})
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name()))
	c.File(file.Name())
}
