package site

import (
	"context"
	site "speedy/read/biz/domain/aggregates/site"
)

type Repository struct {
	SiteRepo     *SiteRepo
	SiteMetaRepo *SiteMetaRepo
}

func NewSiteRepository() site.SiteRepo {
	return &Repository{
		SiteRepo:     &SiteRepo{},
		SiteMetaRepo: &SiteMetaRepo{},
	}
}

func (r *Repository) Create(ctx context.Context, siteDO *site.Site) (int64, error) {
	siteID, err := r.SiteRepo.Save(ctx, CovertPO(siteDO))
	if err != nil {
		return int64(0), err
	}
	if siteDO.SiteMeta != nil {
		siteDO.SiteMeta.SiteID = siteID
		r.SiteMetaRepo.Save(ctx, ConvertMetaDOToPO(siteDO.SiteMeta))
	}
	return siteID, nil
}

func (r *Repository) GetSiteList(ctx context.Context) ([]*site.Site, error) {
	sitePOList, err := r.SiteRepo.GetSiteList(ctx)
	if err != nil {
		return nil, err
	}
	siteList := make([]*site.Site, 0)
	for _, po := range sitePOList {
		siteMetaPO, err := r.SiteMetaRepo.GetSiteMetaBySiteID(ctx, po.ID)
		if err != nil {
			continue
		}
		siteList = append(siteList, CovertDO(po, siteMetaPO))
	}
	return siteList, nil
}

func (r *Repository) GetSiteByUrl(ctx context.Context, url string) (*site.Site, error) {
	sitePO, err := r.SiteRepo.GetSiteByUrl(ctx, url)
	if err != nil {
		return nil, err
	}
	siteMetaPO, err := r.SiteMetaRepo.GetSiteMetaBySiteID(ctx, sitePO.ID)
	if err != nil {
		return nil, err
	}
	return CovertDO(sitePO, siteMetaPO), nil
}

func (r *Repository) GetSiteByID(ctx context.Context, id int64) (*site.Site, error) {
	sitePO, err := r.SiteRepo.GetSiteByID(ctx, id)
	if err != nil {
		return nil, err
	}
	siteMetaPO, err := r.SiteMetaRepo.GetSiteMetaBySiteID(ctx, sitePO.ID)
	if err != nil {
		return nil, err
	}
	return CovertDO(sitePO, siteMetaPO), nil
}
