package site

import (
	"speedy/read/biz/domain/aggregates/site"
)

func CovertPO(siteDO *site.Site) *Site {
	if siteDO == nil {
		return nil
	}
	return &Site{
		Url:         siteDO.Url,
		Description: siteDO.Description,
		Tag:         siteDO.Tag,
	}
}

func ConvertMetaDOToPO(metaDO *site.SiteMeta) *SiteMeta {
	if metaDO == nil {
		return nil
	}
	return &SiteMeta{
		SiteID:    metaDO.SiteID,
		MetaType:  metaDO.MetaType,
		MetaKey:   metaDO.MetaKey,
		MetaValue: metaDO.MetaValue,
	}
}

func CovertDO(sitePO *Site, metaPO *SiteMeta) *site.Site {
	if sitePO == nil {
		return nil
	}
	return &site.Site{
		ID:          sitePO.ID,
		SiteMeta:    ConvertMetaPOToDO(metaPO),
		Url:         sitePO.Url,
		Description: sitePO.Description,
		Tag:         sitePO.Tag,
		CreatedAt:   sitePO.CreatedAt,
		UpdatedAt:   sitePO.UpdatedAt,
	}
}

func ConvertMetaPOToDO(metaPO *SiteMeta) *site.SiteMeta {
	if metaPO == nil {
		return nil
	}
	return &site.SiteMeta{
		ID:        metaPO.ID,
		SiteID:    metaPO.SiteID,
		MetaType:  metaPO.MetaType,
		MetaKey:   metaPO.MetaKey,
		MetaValue: metaPO.MetaValue,
	}
}
