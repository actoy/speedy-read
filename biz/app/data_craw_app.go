package app

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/domain/aggregates/site"
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
	//go func() {
	err := impl.dataCrawSvc.CrawArticle(ctx, site.SiteTypeRss)
	if err != nil {
		klog.CtxErrorf(ctx, "craw article, types is rss, err: %v", err)
	}
	//}()
	//go func() {
	err = impl.dataCrawSvc.CrawArticle(ctx, site.SiteTypeCraw)
	if err != nil {
		klog.CtxErrorf(ctx, "craw article, types is craw, err: %v", err)
	}
	//}()
	return nil
}
