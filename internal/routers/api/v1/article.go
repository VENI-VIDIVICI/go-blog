package v1

import (
	"github.com/VENI-VIDIVICI/go-blog/internal/service"
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
	params := service.CreateAraicleRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.CreateArticle(&params); err != nil {
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

//删除指定文章

func (a Article) DeleteArticleById(c *gin.Context) {
	params := service.DeleteArticleRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.DeleteArticle(&params); err != nil {
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

//更新指定文章

func (a Article) UpdateArticleById(c *gin.Context) {
	params := service.UpdateArticleByIdRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.UpdateArticleById(&params); err != nil {
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

//获取指定文章

func (a Article) GetArticleById(c *gin.Context) {
	params := service.GetArticleByIdRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if valid == false {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	data, err := svc.GetArticleById(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{
		"data": data,
	})
	return
}

//获取文章列表

func (a Article) GetArticleList(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
