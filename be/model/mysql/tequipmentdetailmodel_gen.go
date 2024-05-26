// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tEquipmentDetailFieldNames          = builder.RawFieldNames(&TEquipmentDetail{})
	tEquipmentDetailRows                = strings.Join(tEquipmentDetailFieldNames, ",")
	tEquipmentDetailRowsExpectAutoSet   = strings.Join(stringx.Remove(tEquipmentDetailFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tEquipmentDetailRowsWithPlaceHolder = strings.Join(stringx.Remove(tEquipmentDetailFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tEquipmentDetailModel interface {
		Insert(ctx context.Context, data *TEquipmentDetail) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TEquipmentDetail, error)
		Update(ctx context.Context, data *TEquipmentDetail) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTEquipmentDetailModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TEquipmentDetail struct {
		Id            int64          `db:"id"`
		ClassId       int64          `db:"class_id"`       // 分类id
		Name          string         `db:"name"`           // 设备名称
		Type          string         `db:"type"`           // 设备类型
		SpecificModel sql.NullString `db:"specific_model"` // 具体型号
		Unit          sql.NullString `db:"unit"`           // 单位
		Status        sql.NullString `db:"status"`         // 状态
		Image         sql.NullString `db:"image"`          // 图片
		CreateUserId  sql.NullInt64  `db:"create_user_id"` // 创建用户id
	}
)

func newTEquipmentDetailModel(conn sqlx.SqlConn) *defaultTEquipmentDetailModel {
	return &defaultTEquipmentDetailModel{
		conn:  conn,
		table: "`t_equipment_detail`",
	}
}

func (m *defaultTEquipmentDetailModel) withSession(session sqlx.Session) *defaultTEquipmentDetailModel {
	return &defaultTEquipmentDetailModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`t_equipment_detail`",
	}
}

func (m *defaultTEquipmentDetailModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTEquipmentDetailModel) FindOne(ctx context.Context, id int64) (*TEquipmentDetail, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tEquipmentDetailRows, m.table)
	var resp TEquipmentDetail
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTEquipmentDetailModel) Insert(ctx context.Context, data *TEquipmentDetail) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, tEquipmentDetailRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ClassId, data.Name, data.Type, data.SpecificModel, data.Unit, data.Status, data.Image, data.CreateUserId)
	return ret, err
}

func (m *defaultTEquipmentDetailModel) Update(ctx context.Context, data *TEquipmentDetail) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tEquipmentDetailRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ClassId, data.Name, data.Type, data.SpecificModel, data.Unit, data.Status, data.Image, data.CreateUserId, data.Id)
	return err
}

func (m *defaultTEquipmentDetailModel) tableName() string {
	return m.table
}
