package mysql

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TTowerEquipmentModel = (*customTTowerEquipmentModel)(nil)

type (
	// TTowerEquipmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTTowerEquipmentModel.
	TTowerEquipmentModel interface {
		tTowerEquipmentModel
		// FindListByTowerId 方法根据 towerId 查询 TTowerEquipment 列表数据
		FindListByTowerId(ctx context.Context, towerId int64) ([]*TTowerEquipment, error)
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

// FindListByTowerId 方法根据 towerId 查询 TTowerEquipment 列表数据
func (m *customTTowerEquipmentModel) FindListByTowerId(ctx context.Context, towerId int64) ([]*TTowerEquipment, error) {
	query := fmt.Sprintf("select %s from %s where `tower_id` = ?", tTowerEquipmentRows, m.table)
	var resp []*TTowerEquipment
	err := m.conn.QueryRowsCtx(ctx, &resp, query, towerId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
