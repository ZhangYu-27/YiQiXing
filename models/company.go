package models

import (
	"awesomeProject/tool"
	"fmt"
	"strconv"
)

type Company struct {
	ID             int    `json:"id" gorm:"column:id"`                                                                              // 企业id
	Name           string `json:"name" gorm:"column:name" form:"name" binding:"required"`                                           // 公司名称
	CallName       string `json:"call_name" gorm:"column:call_name" form:"call_name"`                                               // 联系人姓名
	Phone          string `json:"phone" gorm:"column:phone" form:"phone"`                                                           // 联系电话
	Intention      string `json:"intention" gorm:"column:intention" form:"intention"`                                               // 意向度
	Industry       string `json:"industry" gorm:"column:industry" form:"industry"`                                                  // 行业
	Trademark      string `json:"trademark" gorm:"column:trademark" form:"trademark"`                                               // 商标用逗号分割
	CreateAt       string `json:"create_at" gorm:"column:create_at"`                                                                // 创建时间
	NextFollowTime string `json:"next_follow_time" gorm:"column:next_follow_time" form:"next_follow_time" time_format:"2006-01-02"` // 下次跟进时间
}

func NewCompany() *Company {
	return &Company{}
}

func (m *Company) TableName() string {
	return "company"
}

func (this Company) GetCompanyInfo(where string, value string) (Company *Company, err error) {
	fmt.Println(where, value)
	res := tool.DB.Where(where, value).Find(&Company)
	fmt.Println(Company)
	if res.Error != nil {
		err = res.Error
		return
	}
	err = nil
	return
}

func (this *Company) AddCompany() error {
	res := tool.DB.Create(this)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}
func (this Company) GetCompanyList(pages string, pageSizes string) (list *[]Company, err error) {
	page, err := strconv.Atoi(pages)
	if err != nil {
		return
	}
	if page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizes)
	if err != nil {
		return
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 50
	}
	res := tool.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&list)
	return list, res.Error
}
