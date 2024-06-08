package svc

import (
	"be/api/internal/config"
	"be/model/mysql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	TEquipmentClassModel  mysql.TEquipmentClassModel
	TEquipmentDetailModel mysql.TEquipmentDetailModel
	TTowerDetailModel     mysql.TTowerDetailModel
	TTowerEquipmentModel  mysql.TTowerEquipmentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:                c,
		TEquipmentClassModel:  mysql.NewTEquipmentClassModel(sqlConn),
		TTowerDetailModel:     mysql.NewTTowerDetailModel(sqlConn),
		TTowerEquipmentModel:  mysql.NewTTowerEquipmentModel(sqlConn),
		TEquipmentDetailModel: mysql.NewTEquipmentDetailModel(sqlConn),
	}
}
