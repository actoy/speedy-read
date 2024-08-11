package app

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/domain/aggregates/site"
	"speedy/read/biz/domain/service"
)

type DateCrawApplicationI interface {
	Craw(ctx context.Context, source string) error
}

type DateCrawApplication struct {
	dataCrawSvc service.DataCrawServiceI
}

func NewDateCrawApplication() DateCrawApplicationI {
	return &DateCrawApplication{
		dataCrawSvc: service.NewDataCrawService(),
	}
}

func (impl *DateCrawApplication) Craw(ctx context.Context, source string) error {
	if source == "rss" {
		err := impl.dataCrawSvc.CrawArticle(ctx, site.SiteTypeRss)
		if err != nil {
			klog.CtxErrorf(ctx, "craw article, types is rss, err: %v", err)
		}
	} else if source == "craw" {
		err := impl.dataCrawSvc.CrawArticle(ctx, site.SiteTypeCraw)
		if err != nil {
			klog.CtxErrorf(ctx, "craw article, types is craw, err: %v", err)
		}
	} else {
		err := impl.dataCrawSvc.CrawArticle(ctx, site.SiteTypeRss)
		if err != nil {
			klog.CtxErrorf(ctx, "craw article, types is rss, err: %v", err)
		}
		err = impl.dataCrawSvc.CrawArticle(ctx, site.SiteTypeCraw)
		if err != nil {
			klog.CtxErrorf(ctx, "craw article, types is craw, err: %v", err)
		}
	}
	return nil
}
