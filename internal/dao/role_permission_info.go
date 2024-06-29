// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-shop/internal/dao/internal"
)

// internalRolePermissionInfoDao is internal type for wrapping internal DAO implements.
type internalRolePermissionInfoDao = *internal.RolePermissionInfoDao

// rolePermissionInfoDao is the data access object for table role_permission_info.
// You can define custom methods on it to extend its functionality as you wish.
type rolePermissionInfoDao struct {
	internalRolePermissionInfoDao
}

var (
	// RolePermissionInfo is globally public accessible object for table role_permission_info operations.
	RolePermissionInfo = rolePermissionInfoDao{
		internal.NewRolePermissionInfoDao(),
	}
)

// Fill with you ideas below.
