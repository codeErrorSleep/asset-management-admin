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
	rolePermissionsPermissionFieldNames          = builder.RawFieldNames(&RolePermissionsPermission{})
	rolePermissionsPermissionRows                = strings.Join(rolePermissionsPermissionFieldNames, ",")
	rolePermissionsPermissionRowsExpectAutoSet   = strings.Join(stringx.Remove(rolePermissionsPermissionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	rolePermissionsPermissionRowsWithPlaceHolder = strings.Join(stringx.Remove(rolePermissionsPermissionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	rolePermissionsPermissionModel interface {
		Insert(ctx context.Context, data *RolePermissionsPermission) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RolePermissionsPermission, error)
		Update(ctx context.Context, data *RolePermissionsPermission) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRolePermissionsPermissionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RolePermissionsPermission struct {
		Id           int64 `db:"id"`
		RoleId       int64 `db:"roleId"`
		PermissionId int64 `db:"permissionId"`
	}
)

func newRolePermissionsPermissionModel(conn sqlx.SqlConn) *defaultRolePermissionsPermissionModel {
	return &defaultRolePermissionsPermissionModel{
		conn:  conn,
		table: "`role_permissions_permission`",
	}
}

func (m *defaultRolePermissionsPermissionModel) withSession(session sqlx.Session) *defaultRolePermissionsPermissionModel {
	return &defaultRolePermissionsPermissionModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`role_permissions_permission`",
	}
}

func (m *defaultRolePermissionsPermissionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRolePermissionsPermissionModel) FindOne(ctx context.Context, id int64) (*RolePermissionsPermission, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rolePermissionsPermissionRows, m.table)
	var resp RolePermissionsPermission
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

func (m *defaultRolePermissionsPermissionModel) Insert(ctx context.Context, data *RolePermissionsPermission) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, rolePermissionsPermissionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RoleId, data.PermissionId)
	return ret, err
}

func (m *defaultRolePermissionsPermissionModel) Update(ctx context.Context, data *RolePermissionsPermission) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rolePermissionsPermissionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.RoleId, data.PermissionId, data.Id)
	return err
}

func (m *defaultRolePermissionsPermissionModel) tableName() string {
	return m.table
}
