package site

import (
	"context"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	site "speedy/read/biz/domain/aggregates/site"
	"speedy/read/biz/infra"
	"time"
)

type SiteRepo struct {
}

func (dal *SiteRepo) Save(ctx context.Context, siteDO *site.Site) (int64, error) {
	siteDO.CreatedAt = time.Now()
	siteDO.UpdatedAt = time.Now()
	result := infra.DB.Create(siteDO)
	if result.Error != nil {
		log.Info(ctx, "save site error: %v", result.Error)
	}
	return siteDO.ID, nil
}

func (dal *SiteRepo) GetSiteList(ctx context.Context) ([]*site.Site, error) {
	siteList := make([]*site.Site, 0)
	result := infra.DB.Find(&siteList)
	if result.Error != nil {
		log.Info(ctx, "save site error: %v", result.Error)
	}
	return siteList, nil
}
