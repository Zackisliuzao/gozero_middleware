// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneLogic {
	return &PhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneLogic) Phone(req *types.Demo) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{Message: req.Phone}
	return
}
