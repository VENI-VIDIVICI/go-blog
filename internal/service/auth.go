package service

import (
	"errors"
	"fmt"
)

type AuthRequest struct {
	Appkey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc Service) GetAuth(params AuthRequest) error {
	auth, err := svc.dao.GetAuth(params.Appkey, params.AppSecret)
	if err != nil {
		return err
	}
	fmt.Println(auth)
	if auth.Model.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
