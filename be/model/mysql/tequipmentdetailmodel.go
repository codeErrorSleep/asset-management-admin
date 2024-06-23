package mysql

import (
	"context"
	"fmt"
	"strings"

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
		FindByIds(ctx context.Context, ids []int64) ([]*TEquipmentDetail, error)
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

// 根据批量id查询数据的方法
func (m *customTEquipmentDetailModel) FindByIds(ctx context.Context, ids []int64) ([]*TEquipmentDetail, error) {
	// 构建查询条件中的占位符
	placeholders := make([]string, len(ids))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	// 将占位符拼接成字符串，用于IN语句
	inPlaceholders := strings.Join(placeholders, ",")
	// 构建查询语句
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `id` IN (%s)", tEquipmentDetailRows, m.table, inPlaceholders)
	// 将ids转换为interface{}类型的切片，用于传递给QueryRowsCtx
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	// 执行查询
	var resp []*TEquipmentDetail
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
