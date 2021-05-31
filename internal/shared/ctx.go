package shared

import (
	"context"

	"github.com/antbiz/antchat/internal/types"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Ctx = &ctxShared{
	ctxKey: "ctxKey",
}

type ctxShared struct {
	ctxKey string
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *ctxShared) Init(r *ghttp.Request, customCtx *types.Context) {
	r.SetCtxVar(s.ctxKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *ctxShared) Get(ctx context.Context) *types.Context {
	value := ctx.Value(s.ctxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*types.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 设置上下文中的 User 信息
func (s *ctxShared) SetUser(ctx context.Context, ctxUser *types.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// SetData 设置上下文中的其他信息
func (s *ctxShared) SetData(ctx context.Context, data g.Map) {
	s.Get(ctx).Data = data
}
