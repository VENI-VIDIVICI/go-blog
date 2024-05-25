package dao

import (
	"github.com/VENI-VIDIVICI/go-blog/internal/model"
	"github.com/VENI-VIDIVICI/go-blog/pkg/app"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(db *gorm.DB) *Dao {
	return &Dao{
		engine: db,
	}
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}
func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, ModifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{ModifiedBy: ModifiedBy, ID: id},
	}
	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: id, IsDel: 1},
	}
	return tag.Delete(d.engine)
}

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}

func (d *Dao) CreateArticle(title, desc, content, createdBy string) error {
	article := model.Article{
		Title:   title,
		Desc:    desc,
		Content: content,
		Model:   &model.Model{CreatedBy: createdBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) DeleteArticle(id uint32, modifiedBy string) error {
	article := model.Article{
		Model: &model.Model{ID: id, IsDel: 1, ModifiedBy: modifiedBy},
	}
	return article.Delete(d.engine)
}
