package mysql

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TEquipmentDetailModel = (*customTEquipmentDetailModel)(nil)

type (
	// TEquipmentDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTEquipmentDetailModel.
	TEquipmentDetailModel interface {
		tEquipmentDetailModel
		ListByPage(ctx context.Context, page, pageSize int64) ([]*TEquipmentDetail, error)
		Count(ctx context.Context) (int64, error)
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

// 分页查询列表
func (m *customTEquipmentDetailModel) ListByPage(ctx context.Context, page, pageSize int64) ([]*TEquipmentDetail, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT ?, ?", tEquipmentDetailRows, m.table)
	var resp []*TEquipmentDetail
	err := m.conn.QueryRowsCtx(ctx, &resp, query, offset, pageSize)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 获取总数量
func (m *customTEquipmentDetailModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}
