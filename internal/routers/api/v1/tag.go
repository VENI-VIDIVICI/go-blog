package v1

import (
	"github.com/VENI-VIDIVICI/go-blog/global"
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

}

// 删除指定标签

func (t Tag) DeleteTagById(c *gin.Context) {

}

//更新指定标签

func (t Tag) UpdateTagById(c *gin.Context) {

}

//获取全部标签

func (t Tag) GetAll(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == false {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	//response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	response.ToResponse(gin.H{})
}
