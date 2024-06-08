package svc

import (
	"be/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	// sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config: c,
	}
}
