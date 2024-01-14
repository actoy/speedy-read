package site

import (
	"context"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
	"time"
)

type SiteRepo struct {
}

func (dal *SiteRepo) Save(ctx context.Context, sitePO *Site) (int64, error) {
	sitePO.ID = infra.IdGenerate()
	sitePO.CreatedAt = time.Now()
	sitePO.UpdatedAt = time.Now()
	result := infra.DB.WithContext(ctx).Create(sitePO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return sitePO.ID, nil
}

func (dal *SiteRepo) GetSiteList(ctx context.Context) ([]*Site, error) {
	siteList := make([]*Site, 0)
	result := infra.DB.WithContext(ctx).Find(&siteList)
	if result.Error == nil {
		return siteList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *SiteRepo) GetSiteByUrl(ctx context.Context, url string) (*Site, error) {
	sitePO := &Site{}
	result := infra.DB.WithContext(ctx).Where("url = ?", url).First(&sitePO)
	if result.Error == nil {
		return sitePO, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *SiteRepo) GetSiteByID(ctx context.Context, id int64) (*Site, error) {
	sitePO := &Site{}
	result := infra.DB.WithContext(ctx).First(&sitePO, id)
	if result.Error == nil {
		return sitePO, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}
