package rpc

import (
	"context"
	"speedy/read/biz/app"
	"speedy/read/biz/conversion"
	"speedy/read/kitex_gen/speedy_read"

	"github.com/cloudwego/kitex/pkg/klog"
)

type SymbolHandlerI interface {
	Import(ctx context.Context, req *speedy_read.Request) (resp *speedy_read.Response, err error)
	GetSymbolList(ctx context.Context, req *speedy_read.SymbolListRequest) (resp *speedy_read.SymbolListResponse, err error)
	SearchSymbol(ctx context.Context, req *speedy_read.SearchSymbolRequest) (resp *speedy_read.SearchSymbolResponse, err error)
	UpdateSymbol(ctx context.Context, req *speedy_read.UpdateSymbolRequest) (resp *speedy_read.UpdateSymbolResponse, err error)
	GetSymbol(ctx context.Context, req *speedy_read.GetSymbolRequest) (resp *speedy_read.GetSybmolResponse, err error)
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
	klog.CtxErrorf(ctx, "import symbol start")
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

func (s *SymbolHandler) SearchSymbol(ctx context.Context, req *speedy_read.SearchSymbolRequest) (resp *speedy_read.SearchSymbolResponse, err error) {
	resp = &speedy_read.SearchSymbolResponse{}
	list, err := s.SymbolSvc.SearchSymbolByKeyword(ctx, req.GetKeyWord())
	if err != nil {
		klog.CtxErrorf(ctx, "import symbol err %v", err)
		return resp, err
	}
	symbolList := make([]*speedy_read.Symbol, 0)
	for _, info := range list {
		symbolList = append(symbolList, conversion.SymbolDOToThrift(info))
	}
	resp.SymbolList = symbolList
	return resp, nil
}

func (s *SymbolHandler) UpdateSymbol(ctx context.Context, req *speedy_read.UpdateSymbolRequest) (resp *speedy_read.UpdateSymbolResponse, err error) {
	resp = &speedy_read.UpdateSymbolResponse{
		Success: false,
	}
	err = s.SymbolSvc.UpdateSymbol(ctx, app.UpdateSymbolParams{
		ID:              req.GetID(),
		Company:         req.Company,
		CompanyZH:       req.CompanyZH,
		CompanyUrl:      req.CompanyUrl,
		CompanyAddress:  req.CompanyAddress,
		CompanyBusiness: req.CompanyBusiness,
		Description:     req.Description,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "update symbol err %v", err)
		return resp, err
	}
	resp.Success = true
	return resp, nil
}

func (s *SymbolHandler) GetSymbol(ctx context.Context, req *speedy_read.GetSymbolRequest) (resp *speedy_read.GetSybmolResponse, err error) {
	resp = &speedy_read.GetSybmolResponse{}
	if req.GetID() != "" {
		info, err := s.SymbolSvc.GetSymbolByID(ctx, req.GetID())
		if err != nil {
			klog.CtxErrorf(ctx, "import symbol err %v", err)
			return resp, err
		}
		resp.Symbol = conversion.SymbolDOToThrift(info)
	} else if req.GetSymbolTag() != "" {
		info, err := s.SymbolSvc.GetBySymbol(ctx, req.GetSymbolTag())
		if err != nil {
			klog.CtxErrorf(ctx, "import symbol err %v", err)
			return resp, err
		}
		resp.Symbol = conversion.SymbolDOToThrift(info)
	}
	return resp, nil
}
