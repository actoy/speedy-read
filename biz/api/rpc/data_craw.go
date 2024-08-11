package rpc

import (
	"context"
	"speedy/read/biz/app"
	"speedy/read/kitex_gen/speedy_read"
)

type DateCrawI interface {
	Echo(ctx context.Context, req *speedy_read.Request) (resp *speedy_read.Response, err error)
	CrawData(ctx context.Context, req *speedy_read.CrawDataRequest) (resp *speedy_read.Response, err error)
}

type DataCrawHandler struct {
	crawSvc app.DateCrawApplicationI
}

func NewDataCrawHandler() DateCrawI {
	return &DataCrawHandler{
		crawSvc: app.NewDateCrawApplication(),
	}
}

// Echo implements the SpeedyReadImpl interface.
func (s *DataCrawHandler) Echo(ctx context.Context, req *speedy_read.Request) (resp *speedy_read.Response, err error) {
	s.crawSvc.Craw(ctx, "")
	return &speedy_read.Response{Message: req.Message}, nil
}

func (s *DataCrawHandler) CrawData(ctx context.Context, req *speedy_read.CrawDataRequest) (resp *speedy_read.Response, err error) {
	s.crawSvc.Craw(ctx, req.GetSource())
	return &speedy_read.Response{Message: "success"}, nil
}
