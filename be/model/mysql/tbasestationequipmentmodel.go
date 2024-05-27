package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TBaseStationEquipmentModel = (*customTBaseStationEquipmentModel)(nil)

type (
	// TBaseStationEquipmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTBaseStationEquipmentModel.
	TBaseStationEquipmentModel interface {
		tBaseStationEquipmentModel
	}

	customTBaseStationEquipmentModel struct {
		*defaultTBaseStationEquipmentModel
	}
)

// NewTBaseStationEquipmentModel returns a model for the database table.
func NewTBaseStationEquipmentModel(conn sqlx.SqlConn) TBaseStationEquipmentModel {
	return &customTBaseStationEquipmentModel{
		defaultTBaseStationEquipmentModel: newTBaseStationEquipmentModel(conn),
	}
}
