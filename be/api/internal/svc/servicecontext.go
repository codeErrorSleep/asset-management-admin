package svc

import (
	"be/api/internal/config"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	TBaseStation mysql.TBaseStationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:       c,
		TBaseStation: mysql.NewTBaseStationModel(sqlConn),
	}
}
