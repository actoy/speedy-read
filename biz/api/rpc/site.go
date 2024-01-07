package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/app"
	apiConvert "speedy/read/biz/conversion/api"
	site "speedy/read/biz/domain/aggregates/site"
	"speedy/read/biz/utils"
	"speedy/read/kitex_gen/speedy_read"
)

type SiteHandlerI interface {
	GetSiteInfo(ctx context.Context, req *speedy_read.GetSiteRequest) (resp *speedy_read.GetSiteResponse, err error)
	CreateSiteInfo(ctx context.Context, req *speedy_read.CreateSiteRequest) (resp *speedy_read.CreateSiteResponse, err error)
}

type SiteHandler struct {
	SiteSvc app.SiteApplicationI
}

func NewSiteHandler() SiteHandlerI {
	return &SiteHandler{
		SiteSvc: app.NewSiteApplication(),
	}
}

func (s *SiteHandler) GetSiteInfo(ctx context.Context, req *speedy_read.GetSiteRequest) (resp *speedy_read.GetSiteResponse, err error) {
	siteInfoList, err := s.SiteSvc.GetSiteInfo(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "get site info error %v", err)
		return nil, err
	}
	siteList := make([]*speedy_read.SiteInfo, 0)
	for _, siteInfo := range siteInfoList {
		siteList = append(siteList, apiConvert.SiteDOToThrift(siteInfo))
	}
	return &speedy_read.GetSiteResponse{
		SiteList: siteList,
	}, nil
}

func (s *SiteHandler) CreateSiteInfo(ctx context.Context, req *speedy_read.CreateSiteRequest) (resp *speedy_read.CreateSiteResponse, err error) {
	siteDO := &site.Site{
		Url:         req.Url,
		Description: req.Description,
		SiteMeta: &site.SiteMeta{
			MetaType:  req.MetaType,
			MetaValue: req.MetaValue,
			MetaKey:   req.MetaKey,
		},
		Tag: req.Tag,
	}
	id, err := s.SiteSvc.CreateSite(ctx, siteDO)
	if err != nil {
		klog.CtxErrorf(ctx, "create site error %v", err)
		return nil, err
	}
	return &speedy_read.CreateSiteResponse{
		ID: utils.Int64ToString(id),
	}, nil
}
