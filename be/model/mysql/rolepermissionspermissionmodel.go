package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RolePermissionsPermissionModel = (*customRolePermissionsPermissionModel)(nil)

type (
	// RolePermissionsPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRolePermissionsPermissionModel.
	RolePermissionsPermissionModel interface {
		rolePermissionsPermissionModel
	}

	customRolePermissionsPermissionModel struct {
		*defaultRolePermissionsPermissionModel
	}
)

// NewRolePermissionsPermissionModel returns a model for the database table.
func NewRolePermissionsPermissionModel(conn sqlx.SqlConn) RolePermissionsPermissionModel {
	return &customRolePermissionsPermissionModel{
		defaultRolePermissionsPermissionModel: newRolePermissionsPermissionModel(conn),
	}
}
