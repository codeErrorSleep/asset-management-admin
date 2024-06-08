package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TTowerEquipmentModel = (*customTTowerEquipmentModel)(nil)

type (
	// TTowerEquipmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTowerEquipmentModel.
	TTowerEquipmentModel interface {
		tTowerEquipmentModel
	}

	customTTowerEquipmentModel struct {
		*defaultTTowerEquipmentModel
	}
)

// NewTTowerEquipmentModel returns a model for the database table.
func NewTTowerEquipmentModel(conn sqlx.SqlConn) TTowerEquipmentModel {
	return &customTTowerEquipmentModel{
		defaultTTowerEquipmentModel: newTTowerEquipmentModel(conn),
	}
}
