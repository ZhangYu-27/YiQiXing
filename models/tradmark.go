package models

import (
	"awesomeProject/tool"
	"gorm.io/gorm"
)

type Tradmark struct {
	ID           int    `json:"id" gorm:"column:id"`
	Trademark    string `json:"trademark" gorm:"column:trademark" form:"trademark" binding:"required"`    // 商标名称
	State        int8   `json:"state" gorm:"column:state" form:"state" binding:"required"`                // 1已递交2回执3受理4驳回5初审6注册证书
	CreateAt     string `json:"create_at" gorm:"column:create_at"`                                        // 创建时间
	CompanyID    int    `json:"company_id" gorm:"column:company_id" form:"company_id" binding:"required"` // 公司名称
	UpdateAt     string `json:"update_at" gorm:"column:update_at"`                                        // 更新时间
	Class        string `json:"class" gorm:"column:class" form:"class" binding:"required"`                // 属于那一类
	RegisterCode string `json:"register_code" gorm:"column:register_code"`
}

type UpdateTradmark struct {
	ID     int    `json:"id" form:"tradmark_id" gorm:"column:id" binding:"required"`
	Reamrk string `json:"reamrk" form:"remark" gorm:"column:remark"`
	State  int    `json:"state" form:"state" gorm:"column:state"`
}

func NewTradmark() *Tradmark {
	return &Tradmark{}
}
func (*UpdateTradmark) TableName() string {
	return "tradmark"

}

func (m *Tradmark) TableName() string {
	return "tradmark"
}

func (this *Tradmark) AddTradmark() error {
	res := tool.DB.Create(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (this *Tradmark) GetTradmarkList(db *gorm.DB) (Tradmark *[]Tradmark, err error) {
	res := db.Order("id desc").Find(&Tradmark)
	if res.Error != nil {
		return Tradmark, res.Error
	}
	return
}
func (this *UpdateTradmark) UpdateTradmark() error {
	res := tool.DB.Save(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (this Tradmark) GetTradmarkInfo(where string, value string) (tradmarkInfo *Tradmark, err error) {

	res := tool.DB.Where(where, value).Find(&tradmarkInfo)
	if res.Error != nil {
		return tradmarkInfo, res.Error
	}
	return tradmarkInfo, nil
}
