package model

import (
	"github.com/VENI-VIDIVICI/go-blog/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
	Tag           []Tag  `gorm:"many2many:blog_article_tag" json:"tag"`
}
type SwaggerArticle struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) error {
	err := db.Omit("Tag").Create(&a).Error
	if err != nil {
		return err
	}
	return db.Model(&a).Preload("Tag").Association("Tag").Replace(a.Tag).Error
}

func (a Article) Delete(db *gorm.DB) error {
	// return db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error
	res := db.Model(&Article{}).Preload("Tag").Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a)
	return res.Error
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).First(&article).Error
	return article, err
}

func (a Article) Update(db *gorm.DB) error {
	// var article Article
	return db.Model(&Article{}).Where("id = ? AND is_del = ?", a.Model.ID, 0).Update(&a).Error
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]Article, error) {
	var article []Article
	var err error
	if pageSize > 0 && pageOffset > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	if err = db.Where("is_del = ?", 0).Find(&article).Error; err != nil {
		return article, err
	}
	return article, err
}
