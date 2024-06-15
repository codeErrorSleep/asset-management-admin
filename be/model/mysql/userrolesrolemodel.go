package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserRolesRoleModel = (*customUserRolesRoleModel)(nil)

type (
	// UserRolesRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRolesRoleModel.
	UserRolesRoleModel interface {
		userRolesRoleModel
	}

	customUserRolesRoleModel struct {
		*defaultUserRolesRoleModel
	}
)

// NewUserRolesRoleModel returns a model for the database table.
func NewUserRolesRoleModel(conn sqlx.SqlConn) UserRolesRoleModel {
	return &customUserRolesRoleModel{
		defaultUserRolesRoleModel: newUserRolesRoleModel(conn),
	}
}
