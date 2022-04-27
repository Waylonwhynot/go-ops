// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Plugin is the golang structure of table plugin for DAO operations like Where/Data.
type Plugin struct {
	g.Meta      `orm:"table:plugin, do:true"`
	Id          interface{} //
	Uuid        interface{} // 插件uuid
	Name        interface{} // 插件名
	PackageName interface{} // 包名
	Os          interface{} // 操作系统
	Arch        interface{} // 架构
	Md5         interface{} // 包md5名称
	Creater     interface{} // 创建人
	Updater     interface{} // 更新人
	Created     *gtime.Time //
	Updated     *gtime.Time //
}