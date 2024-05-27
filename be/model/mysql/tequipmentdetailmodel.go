package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TEquipmentDetailModel = (*customTEquipmentDetailModel)(nil)

type (
	// TEquipmentDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTEquipmentDetailModel.
	TEquipmentDetailModel interface {
		tEquipmentDetailModel
	}

	customTEquipmentDetailModel struct {
		*defaultTEquipmentDetailModel
	}
)

// NewTEquipmentDetailModel returns a model for the database table.
func NewTEquipmentDetailModel(conn sqlx.SqlConn) TEquipmentDetailModel {
	return &customTEquipmentDetailModel{
		defaultTEquipmentDetailModel: newTEquipmentDetailModel(conn),
	}
}
