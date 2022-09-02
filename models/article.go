package models

import "awesomeProject/tool"

type Article struct {
	ID       int    `json:"id" gorm:"column:id"`
	Lable    string `json:"lable" gorm:"column:lable" form:"lable" binding:"required"` // 文章标签逗号分隔
	Text     string `json:"text" gorm:"column:text" form:"text" binding:"required"`    // 文章内容
	CreateAt string `json:"create_at" gorm:"column:create_at"`                         // 创建时间
	UpdateAt string `json:"update_at" gorm:"column:update_at"`                         // 更新时间
	IsDelete int    `json:"is_delete" gorm:"column:is_delete"`
	Title    string `json:"title" gorm:"column:title" form:"title" binding:"required"`
}

type UpdateArticle struct {
	ID       int    `json:"id" gorm:"column:id" form:"article_id" binding:"required"`
	Lable    string `json:"lable" gorm:"column:lable" form:"lable" binding:"required"` // 文章标签逗号分隔
	Text     string `json:"text" gorm:"column:text" form:"text" binding:"required"`    // 文章内容
	UpdateAt string `json:"update_at" gorm:"column:update_at"`                         // 更新时间
	IsDelete int    `json:"is_delete" gorm:"column:is_delete" form:"is_delete"`
	Title    string `json:"title" gorm:"column:title" form:"title" binding:"required"`
}

func (m *Article) TableName() string {
	return "article"
}
func (m *UpdateArticle) TableName() string {
	return "article"
}

func (this *Article) AddArticle() error {
	res := tool.DB.Create(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (this Article) GetArticle(where string, value string) (*Article, error) {
	article := &Article{}
	res := tool.DB.Where(where, value).Find(article)
	if res.Error != nil {
		return article, res.Error
	}
	return article, nil
}

func (this *UpdateArticle) UpdateArticle() error {
	res := tool.DB.Save(this)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
