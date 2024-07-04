package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PositionCreateReq struct {
	g.Meta    `path:"/backend/position/add" tags:"Position" method:"post" summary:"You first api"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name"    v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsID   int    `json:"goods_id"    v:"required#商品ID不能为空" dc:"商品ID"`
	Sort      int    `json:"sort"    dc:"排序"`
}
type PositionCreateRes struct {
	PositionId int `json:"position_id"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/backend/position/delete" method:"delete" tags:"手工位" summary:"删除手工位接口"`
	Id     int `json:"id" v:"min:1#请选择需要删除的手工位" dc:"手工位id"`
}

type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta    `path:"/backend/position/update/{Id}" method:"post" tags:"手工位" summary:"修改手工位接口"`
	Id        int    `json:"id"      v:"min:1#请选择需要修改的内容" dc:"手工位Id"`
	PicUrl    string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#图片跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name"    v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsID   int    `json:"goods_id"    v:"required#商品ID不能为空" dc:"商品ID"`
	Sort      int    `json:"int"   dc:"排序"`
}
type PositionUpdateRes struct{}

type PositionGetListCommonReq struct {
	g.Meta `path:"/backend/position/list" method:"get" tags:"手工位" summary:"手工位列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type PositionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总量"`
}
