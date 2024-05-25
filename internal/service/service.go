package service

import (
	"context"
	"github.com/VENI-VIDIVICI/go-blog/global"
	"github.com/VENI-VIDIVICI/go-blog/internal/dao"
)

type Service struct {
	dao *dao.Dao
	ctx context.Context
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DbEngine)
	return svc
}
