package api

import (
	"fmt"
	"github.com/VENI-VIDIVICI/go-blog/global"
	"github.com/VENI-VIDIVICI/go-blog/internal/service"
	"github.com/VENI-VIDIVICI/go-blog/pkg/app"
	"github.com/VENI-VIDIVICI/go-blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	params := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	fmt.Println(params)
	global.Logger.Infof("%v", errs.Errors())
	if valid == false {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c)
	err := svc.GetAuth(params)
	fmt.Println(err)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetAuth)
		return
	}
	token, err := app.GenerateToken(params.Appkey, params.AppSecret)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetAuth)
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
