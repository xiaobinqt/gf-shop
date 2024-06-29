// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure of table admin_info for DAO operations like Where/Data.
type AdminInfo struct {
	g.Meta    `orm:"table:admin_info, do:true"`
	Id        interface{} //
	Name      interface{} // 用户名
	Password  interface{} // 密码
	RoleIds   interface{} // 角色ids
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	UserSalt  interface{} // 加密盐
	IsAdmin   interface{} // 是否超级管理员
}
