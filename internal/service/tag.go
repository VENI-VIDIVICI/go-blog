package service

import (
	"github.com/VENI-VIDIVICI/go-blog/internal/model"
	"github.com/VENI-VIDIVICI/go-blog/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	Id         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"required,min=3,max=100"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteTagRequest struct {
	Id uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(params *UpdateTagRequest) error {
	return svc.dao.UpdateTag(params.Id, params.Name, params.State, params.ModifiedBy)
}

func (svc *Service) DeleteTag(params *DeleteTagRequest) error {
	return svc.dao.DeleteTag(params.Id)
}
