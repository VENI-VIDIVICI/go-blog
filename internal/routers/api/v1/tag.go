package v1

import (
	"fmt"
	"github.com/VENI-VIDIVICI/go-blog/global"
	"github.com/VENI-VIDIVICI/go-blog/internal/service"
	"github.com/VENI-VIDIVICI/go-blog/pkg/app"
	"github.com/VENI-VIDIVICI/go-blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.SwaggerTag "成功"
// @Router /api/v1/tags [post]
func (t Tag) CreateTag(c *gin.Context) {
	//生成请求包装，外加校验
	params := service.CreateTagRequest{}
	fmt.Println(params)
	//生成返回包装
	response := app.NewResponse(c)
	//校验
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.CreateTag(&params); err != nil {
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (t Tag) CountTag(c *gin.Context) {
	params := service.CountTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	count, err := svc.CountTag(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	response.ToResponse(gin.H{
		"count": count,
	})
	return
}

// 删除指定标签

func (t Tag) DeleteTagById(c *gin.Context) {
	params := service.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if valid == false {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{
		"message": "删除成功",
	})
}

//更新指定标签

func (t Tag) UpdateTagById(c *gin.Context) {
	params := service.UpdateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if valid == false {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{
		"message": "修改成功",
	})
}

//获取全部标签

func (t Tag) GetAll(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == false {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pageConfig := app.Pager{Page: 1, PageSize: 10}
	list, err := svc.GetTagList(&param, &pageConfig)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	//list, err :=
	//response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	response.ToResponse(gin.H{
		"data": list,
	})
}
