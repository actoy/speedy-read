package app

import (
	"context"
	site "speedy/read/biz/domain/aggregates/site"
	siteInfra "speedy/read/biz/infra/repository/site"
)

type SiteApplicationI interface {
	CreateSite(ctx context.Context, siteDO *site.Site) (int64, error)
	GetSiteInfo(ctx context.Context) ([]*site.Site, error)
}

type SiteApplication struct {
	siteRepo site.SiteRepo
}

func NewSiteApplication() SiteApplicationI {
	return &SiteApplication{
		siteRepo: siteInfra.NewSiteRepository(),
	}
}

func (impl *SiteApplication) CreateSite(ctx context.Context, siteDO *site.Site) (int64, error) {
	return impl.siteRepo.Create(ctx, siteDO)
}

func (impl *SiteApplication) GetSiteInfo(ctx context.Context) ([]*site.Site, error) {
	return impl.siteRepo.GetSiteList(ctx)
}
