package mysql

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TBaseStationModel = (*customTBaseStationModel)(nil)

type (
	// TBaseStationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTBaseStationModel.
	TBaseStationModel interface {
		tBaseStationModel
		ListByPage(ctx context.Context, page, pageSize int64) ([]*TBaseStation, error)
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

func (m *customTBaseStationModel) ListByPage(ctx context.Context, page, pageSize int64) ([]*TBaseStation, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT ?, ?", tBaseStationRows, m.table)
	var resp []*TBaseStation
	err := m.conn.QueryRowsCtx(ctx, &resp, query, offset, pageSize)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
