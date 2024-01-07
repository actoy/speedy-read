package site

import (
	"speedy/read/biz/domain/aggregates/site"
)

func CovertPO(siteDO *site.Site) *Site {
	if siteDO == nil {
		return nil
	}
	return &Site{
		ID:          siteDO.ID,
		SourceID:    siteDO.SourceID,
		SourceType:  siteDO.SourceType,
		Url:         siteDO.Url,
		Description: siteDO.Description,
		Tag:         siteDO.Tag,
		CreatedAt:   siteDO.CreatedAt,
		UpdatedAt:   siteDO.UpdatedAt,
	}
}

func CovertDO(sitePO *Site) *site.Site {
	if sitePO == nil {
		return nil
	}
	return &site.Site{
		ID:          sitePO.ID,
		SourceID:    sitePO.SourceID,
		SourceType:  sitePO.SourceType,
		Url:         sitePO.Url,
		Description: sitePO.Description,
		Tag:         sitePO.Tag,
		CreatedAt:   sitePO.CreatedAt,
		UpdatedAt:   sitePO.UpdatedAt,
	}
}
