package app

import (
	"context"
	"speedy/read/biz/domain/service"
)

type DateCrawApplicationI interface {
	Craw(ctx context.Context) error
}

type DateCrawApplication struct {
	dataCrawSvc service.DataCrawServiceI
}

func NewDateCrawApplication() DateCrawApplicationI {
	return &DateCrawApplication{
		dataCrawSvc: service.NewDataCrawService(),
	}
}

func (impl *DateCrawApplication) Craw(ctx context.Context) error {
	return impl.dataCrawSvc.CrawArticle(ctx)
}
