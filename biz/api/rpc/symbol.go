package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/app"
	"speedy/read/biz/conversion"
	"speedy/read/kitex_gen/speedy_read"
)

type SymbolHandlerI interface {
	Import(ctx context.Context, req *speedy_read.Request) (resp *speedy_read.Response, err error)
	GetSymbolList(ctx context.Context, req *speedy_read.SymbolListRequest) (resp *speedy_read.SymbolListResponse, err error)
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

func (s *SymbolHandler) GetSymbolList(ctx context.Context, req *speedy_read.SymbolListRequest) (resp *speedy_read.SymbolListResponse, err error) {
	resp = &speedy_read.SymbolListResponse{}
	list, err := s.SymbolSvc.GetSymbolList(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "import symbol err %v", err)
		return resp, err
	}
	symbolList := make([]*speedy_read.Symbol, 0)
	for _, info := range list {
		symbolList = append(symbolList, conversion.SymbolDOToThrift(info))
	}
	resp.Symbol = symbolList
	return resp, nil
}
