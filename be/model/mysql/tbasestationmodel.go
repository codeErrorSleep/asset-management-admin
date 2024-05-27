package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TBaseStationModel = (*customTBaseStationModel)(nil)

type (
	// TBaseStationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTBaseStationModel.
	TBaseStationModel interface {
		tBaseStationModel
	}

	customTBaseStationModel struct {
		*defaultTBaseStationModel
	}
)

// NewTBaseStationModel returns a model for the database table.
func NewTBaseStationModel(conn sqlx.SqlConn) TBaseStationModel {
	return &customTBaseStationModel{
		defaultTBaseStationModel: newTBaseStationModel(conn),
	}
}
