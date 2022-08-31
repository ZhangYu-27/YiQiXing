package models

import (
	"awesomeProject/tool"
	"fmt"
)

type AdminUser struct {
	ID       int    `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"` // 账户名
	Password string `json:"password" gorm:"column:password"` // 密码
	LoginIp  string `json:"login_ip" gorm:"column:login_ip"` // 登录ip
	LoginAt  string `json:"login_at" gorm:"column:login_at"` // 登录时间
	Key      string `gorm:"column:key"`
}

func (m *AdminUser) TableName() string {
	return "admin"
}
func (this AdminUser) GetAdminInfo(where string, value string) *AdminUser {
	AdminInfo := new(AdminUser)
	fmt.Println(where, value)
	tool.DB.Where(where, value).Find(AdminInfo)
	return AdminInfo
}
