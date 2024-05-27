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
}
type SwaggerArticle struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error
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
