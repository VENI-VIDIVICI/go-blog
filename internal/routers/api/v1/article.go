package v1

import (
	"github.com/VENI-VIDIVICI/go-blog/pkg/app"
	"github.com/VENI-VIDIVICI/go-blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

//新增文章

func (a Article) CreateArticle(c *gin.Context) {
}

//删除指定文章

func (a Article) DeleteArticleById(c *gin.Context) {

}

//更新指定文章

func (a Article) UpdateArticleById(c *gin.Context) {

}

//获取指定文章

func (a Article) GetArticleById(c *gin.Context) {

}

//获取文章列表

func (a Article) GetArticleList(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
