package model

import "github.com/gogf/gf/v2/os/gtime"

// RotationCreateUpdateBase 创建/修改轮播图基类
type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

// RotationCreateInput 创建轮播图
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建轮播图返回结果
type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}

// RotationUpdateInput 修改轮播图
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id int
}

// RotationGetListInput 获取内容列表
type RotationGetListInput struct {
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	UserId     uint   // 要查询的用户ID
}

// RotationGetListOutput 查询列表结果
type RotationGetListOutput struct {
	List  []RotationGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

// RotationSearchInput 搜索列表
type RotationSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationSearchOutput 搜索列表结果
type RotationSearchOutput struct {
	List  []RotationSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

type RotationGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	Sort      int         `json:"sort"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type RotationSearchOutputItem struct {
	RotationGetListOutputItem
}

// RotationListItem 主要用于列表展示
type RotationListItem struct {
	Id        uint        `json:"id"` // 自增ID
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	Sort      int         `json:"sort"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

// RotationListCategoryItem 绑定到Rotation列表中的栏目信息
type RotationListCategoryItem struct {
	Id           uint   `json:"id"`            // 分类ID，自增主键
	Name         string `json:"name"`          // 分类名称
	Thumb        string `json:"thumb"`         // 封面图
	RotationType string `json:"rotation_type"` // 内容类型

}

// RotationListUserItem 绑定到Rotation列表中的用户信息
type RotationListUserItem struct {
	Id       uint   `json:"id"`       // UID
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像地址
}
