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
	return r.SiteRepo.Save(ctx, CovertPO(siteDO))
}

func (r *Repository) GetSiteList(ctx context.Context) ([]*site.Site, error) {
	sitePOList, err := r.SiteRepo.GetSiteList(ctx)
	if err != nil {
		return nil, err
	}
	siteList := make([]*site.Site, 0)
	for _, po := range sitePOList {
		siteList = append(siteList, CovertDO(po))
	}
	return siteList, nil
}

func (r *Repository) GetSiteByUrl(ctx context.Context, url string) (*site.Site, error) {
	sitePO, err := r.SiteRepo.GetSiteByUrl(ctx, url)
	if err != nil {
		return nil, err
	}
	return CovertDO(sitePO), nil
}
