package api

import (
	"speedy/read/biz/domain/aggregates/site"
	"speedy/read/biz/utils"
	"speedy/read/kitex_gen/speedy_read"
)

func SiteDOToThrift(siteDO *site.Site) *speedy_read.SiteInfo {
	return &speedy_read.SiteInfo{
		ID:          utils.Int64ToString(siteDO.ID),
		SiteMeta:    SiteMetaDOToThrift(siteDO.SiteMeta),
		Url:         siteDO.Url,
		Description: siteDO.Description,
		Tag:         siteDO.Tag,
	}
}

func SiteMetaDOToThrift(siteMetaDO *site.SiteMeta) *speedy_read.SiteMeta {
	if siteMetaDO == nil {
		return &speedy_read.SiteMeta{}
	}
	return &speedy_read.SiteMeta{
		ID:        utils.Int64ToString(siteMetaDO.ID),
		SiteID:    utils.Int64ToString(siteMetaDO.SiteID),
		MetaType:  siteMetaDO.MetaType,
		MetaValue: siteMetaDO.MetaValue,
		MetaKey:   siteMetaDO.MetaKey,
		CreatedAt: siteMetaDO.CreatedAt.String(),
		UpdatedAt: siteMetaDO.UpdatedAt.String(),
	}
}
