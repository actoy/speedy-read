package site

import (
	"context"
	site "speedy/read/biz/domain/aggregates/site"
)

type Repository struct {
	SiteRepo *SiteRepo
}

func NewSiteRepository() site.SiteRepo {
	return &Repository{
		SiteRepo: &SiteRepo{},
	}
}

func (r *Repository) Create(ctx context.Context, siteDO *site.Site) (int64, error) {
	return r.SiteRepo.Save(ctx, siteDO)
}

func (r *Repository) GetSiteList(ctx context.Context) ([]*site.Site, error) {
	return r.SiteRepo.GetSiteList(ctx)
}
