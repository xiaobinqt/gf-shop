// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoodsCommentsInfo is the golang structure for table order_goods_comments_info.
type OrderGoodsCommentsInfo struct {
	Id             int         `json:"id"             orm:"id"               description:""`
	OrderId        int         `json:"orderId"        orm:"order_id"         description:"订单id"`
	GoodsId        int         `json:"goodsId"        orm:"goods_id"         description:"商品id"`
	GoodsOptionsId int         `json:"goodsOptionsId" orm:"goods_options_id" description:"商品规格id"`
	ParentId       int         `json:"parentId"       orm:"parent_id"        description:"父级评论id"`
	Content        string      `json:"content"        orm:"content"          description:"评论内容"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       description:""`
}
