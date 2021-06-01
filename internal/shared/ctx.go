package shared

import (
	"context"

	"github.com/antbiz/antchat/internal/types"
	"github.com/gogf/gf/net/ghttp"
)

var Ctx = &ctxShared{
	ctxKey: "ctxKey",
}

type ctxShared struct {
	ctxKey string
}

// InitCtxUser .
func (s *ctxShared) InitCtxUser(r *ghttp.Request, customCtx *types.ContextUser) {
	r.SetCtxVar(s.ctxKey, customCtx)
}

// GetCtxUser .
func (s *ctxShared) GetCtxUser(ctx context.Context) *types.ContextUser {
	value := ctx.Value(s.ctxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*types.ContextUser); ok {
		return localCtx
	}
	return nil
}

// InitCtxVisitor .
func (s *ctxShared) InitCtxVisitor(r *ghttp.Request, customCtx *types.ContextVisitor) {
	r.SetCtxVar(s.ctxKey, customCtx)
}

// GetCtxVisitor .
func (s *ctxShared) GetCtxVisitor(ctx context.Context) *types.ContextVisitor {
	value := ctx.Value(s.ctxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*types.ContextVisitor); ok {
		return localCtx
	}
	return nil
}
