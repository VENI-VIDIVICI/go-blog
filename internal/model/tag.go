package model

import "github.com/VENI-VIDIVICI/go-blog/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State string `json:"state"`
}

//swag init -g internal/routers/api/v1/tag.go

type SwaggerTag struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
