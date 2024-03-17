package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/app"
	"speedy/read/kitex_gen/speedy_read"
)

type SymbolHandlerI interface {
	Import(ctx context.Context, req *speedy_read.Request) (resp *speedy_read.Response, err error)
}

type SymbolHandler struct {
	SymbolSvc app.SymbolApplicationI
}

func NewSymbolHandler() SymbolHandlerI {
	return &SymbolHandler{
		SymbolSvc: app.NewSymbolApplication(),
	}
}

func (s *SymbolHandler) Import(ctx context.Context, req *speedy_read.Request) (resp *speedy_read.Response, err error) {
	err = s.SymbolSvc.Import(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "import symbol err %v", err)
	}
	return &speedy_read.Response{Message: req.Message}, nil
}
