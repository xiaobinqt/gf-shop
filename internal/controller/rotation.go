package controller

import (
	"context"
	"gf-shop/api/backend"

	"gf-shop/internal/model"
	"gf-shop/internal/service"
)

// Rotation 内容管理
var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *backend.RotationCreateReq) (res *backend.RotationCreateRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationCreateRes{RotationId: out.RotationId}, nil
}
