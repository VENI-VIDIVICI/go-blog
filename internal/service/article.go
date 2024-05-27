package service

import (
	"github.com/VENI-VIDIVICI/go-blog/internal/model"
)

type CreateArticleRequest struct {
}

type CreateAraicleRequest struct {
	Title     string `form:"title" binding:"required,min=3,max=100"`
	Desc      string `form:"desc" binding:"required,min=3,max=255"`
	Content   string `form:"content" binding:"required,min=3,max=255"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	Id        uint32 `form:"id" binding:"required,gte=1"`
	DeletedBy string `form:"deleted_by" binding:"required"`
}

type GetArticleByIdRequest struct {
	Id uint32 `form:"id" binding:"required"`
}

type UpdateArticleByIdRequest struct {
	Title      string `form:"title"`
	Desc       string `form:"desc" binding:"min=3,max=255"`
	Content    string `form:"content" `
	ModifiedBy string `form:"modified_by"`
	Id         uint32 `form:"id" binding:"required"`
}

func (svc *Service) CreateArticle(params *CreateAraicleRequest) error {
	return svc.dao.CreateArticle(params.Title, params.Desc, params.Content, params.CreatedBy)
}

func (svc *Service) DeleteArticle(params *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(params.Id, params.DeletedBy)
}

func (svc *Service) GetArticleById(params *GetArticleByIdRequest) (model.Article, error) {
	return svc.dao.GetArticleById(params.Id)
}

func (svc *Service) UpdateArticleById(params *UpdateArticleByIdRequest) error {
	return svc.dao.UpdateArticleById(params.Id, params.Title, params.Desc, params.Title)
}
