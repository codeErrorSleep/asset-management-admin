package mysql

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTowerDetailModel = (*customTTowerDetailModel)(nil)

type (
	// TTowerDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTowerDetailModel.
	TTowerDetailModel interface {
		tTowerDetailModel
		List(ctx context.Context, page, pageSize int64) ([]*TTowerDetail, error)
		Count(ctx context.Context) (int64, error)
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

// 分页查询列表
func (m *customTTowerDetailModel) List(ctx context.Context, page, pageSize int64) ([]*TTowerDetail, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT ?,?", tTowerDetailRows, m.table)
	var resp []*TTowerDetail
	err := m.conn.QueryRowsCtx(ctx, &resp, query, offset, pageSize)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 获取总数量
func (m *customTTowerDetailModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}
