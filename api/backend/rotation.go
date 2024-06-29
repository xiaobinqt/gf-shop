package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationCreateReq struct {
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"You first api"`

	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}
type RotationCreateRes struct {
	// todo
	//g.Meta `mime:"text/html" example:"string"`
	RotationId int `json:"rotation_id"`
}
