package tool

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
	"time"
)

func UnixToDate(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")

}
func SetSession(key string, value string, c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		MaxAge: 3600 * 6,
	})
	session.Set(key, value)
	session.Save()
}
func GetSession(key string, c *gin.Context) interface{} {
	session := sessions.Default(c).Get(key)
	return session
}
func UploadImg(c *gin.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	} else {
		_ = strings.ToLower(path.Ext(file.Filename)) //获取文件后缀

		filePath := "./file/" + file.Filename
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			return err
		}
	}
	return nil
}
