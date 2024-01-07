package site

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
	"time"
)

type SiteMetaRepo struct {
}

func (dal *SiteMetaRepo) Save(ctx context.Context, siteMetaPO *SiteMeta) {
	metaPO := &SiteMeta{}
	result := infra.DB.WithContext(ctx).
		Where("site_id = ? and meta_key = ? and meta_value = ?",
			siteMetaPO.SiteID, siteMetaPO.MetaKey, siteMetaPO.MetaValue).
		First(&metaPO)
	if result.Error == nil && metaPO.ID != 0 {
		return
	}
	siteMetaPO.CreatedAt = time.Now()
	siteMetaPO.UpdatedAt = time.Now()
	result = infra.DB.WithContext(ctx).Create(siteMetaPO)
	if result.Error != nil {
		klog.CtxErrorf(ctx, "create site meta error is %v", result.Error)
	}
	return
}

func (dal *SiteMetaRepo) GetSiteMetaBySiteID(ctx context.Context, siteID int64) (*SiteMeta, error) {
	siteMeta := &SiteMeta{}
	result := infra.DB.WithContext(ctx).
		Where("site_id = ?", siteID).
		First(&siteMeta)
	if result.Error == nil {
		return siteMeta, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}
