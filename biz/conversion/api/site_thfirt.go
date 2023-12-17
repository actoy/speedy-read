package api

import (
	"speedy/read/biz/domain/aggregates/site"
	"speedy/read/kitex_gen/speedy_read"
)

func SiteDOToThrift(siteDO *site.Site) *speedy_read.SiteInfo {
	return &speedy_read.SiteInfo{
		ID:          siteDO.ID,
		SourceID:    siteDO.SourceID,
		SourceType:  siteDO.SourceType,
		Url:         siteDO.Url,
		Description: siteDO.Description,
	}
}
