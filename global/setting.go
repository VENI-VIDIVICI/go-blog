package global

import (
	"github.com/VENI-VIDIVICI/go-blog/pkg/logger"
	"github.com/VENI-VIDIVICI/go-blog/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DbEngine        *gorm.DB
	Logger          *logger.Logger
)
