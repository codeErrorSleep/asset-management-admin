package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TEquipmentClassModel = (*customTEquipmentClassModel)(nil)

type (
	// TEquipmentClassModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTEquipmentClassModel.
	TEquipmentClassModel interface {
		tEquipmentClassModel
	}

	customTEquipmentClassModel struct {
		*defaultTEquipmentClassModel
	}
)

// NewTEquipmentClassModel returns a model for the database table.
func NewTEquipmentClassModel(conn sqlx.SqlConn) TEquipmentClassModel {
	return &customTEquipmentClassModel{
		defaultTEquipmentClassModel: newTEquipmentClassModel(conn),
	}
}
