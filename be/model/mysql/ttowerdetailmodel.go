package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TTowerDetailModel = (*customTTowerDetailModel)(nil)

type (
	// TTowerDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTowerDetailModel.
	TTowerDetailModel interface {
		tTowerDetailModel
	}

	customTTowerDetailModel struct {
		*defaultTTowerDetailModel
	}
)

// NewTTowerDetailModel returns a model for the database table.
func NewTTowerDetailModel(conn sqlx.SqlConn) TTowerDetailModel {
	return &customTTowerDetailModel{
		defaultTTowerDetailModel: newTTowerDetailModel(conn),
	}
}
