package models

import "awesomeProject/tool"

type FollowUp struct {
	ID        int    `json:"id" gorm:"column:id"`                                                      // 跟进意向id
	CreateAt  string `json:"create_at" gorm:"column:create_at" form:"follow_time" binding:"required"`  // 跟进时间
	Service   int    `json:"service" gorm:"column:service" form:"service" binding:"required,max=1"`    // 1，商标注册2案件3变转续
	Remarks   string `json:"remarks" gorm:"column:remarks" form:"remarks" binding:"required"`          // 备注
	CompanyID int    `json:"company_id" gorm:"column:company_id" form:"company_id" binding:"required"` // 公司id
}

func (m *FollowUp) TableName() string {
	return "follow_up"
}
func NewFollow() *FollowUp {
	return &FollowUp{}
}
func (this *FollowUp) AddFollow() (err error) {
	res := tool.DB.Create(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (this FollowUp) GetFollowList(where string, value string) (*[]FollowUp, error) {
	followList := make([]FollowUp, 0)
	res := tool.DB.Where(where, value).Order("id desc").Find(&followList)
	if res.Error != nil {
		return &followList, res.Error
	}
	return &followList, nil
}
