package site

import (
	"speedy/read/biz/domain/aggregates/site"
)

func CovertPO(siteDO *site.Site) *Site {
	return &Site{
		ID:          siteDO.ID,
		SourceID:    siteDO.SourceID,
		SourceType:  siteDO.SourceType,
		Url:         siteDO.Url,
		Description: siteDO.Description,
		CreatedAt:   siteDO.CreatedAt,
		UpdatedAt:   siteDO.UpdatedAt,
	}
}

func CovertDO(sitePO *Site) *site.Site {
	return &site.Site{
		ID:          sitePO.ID,
		SourceID:    sitePO.SourceID,
		SourceType:  sitePO.SourceType,
		Url:         sitePO.Url,
		Description: sitePO.Description,
		CreatedAt:   sitePO.CreatedAt,
		UpdatedAt:   sitePO.UpdatedAt,
	}
}
