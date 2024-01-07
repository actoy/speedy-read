package api

import (
	"speedy/read/biz/domain/aggregates/site"
	"speedy/read/biz/utils"
	"speedy/read/kitex_gen/speedy_read"
)

func SiteDOToThrift(siteDO *site.Site) *speedy_read.SiteInfo {
	return &speedy_read.SiteInfo{
		ID:          utils.Int64ToString(siteDO.ID),
		SourceID:    utils.Int64ToString(siteDO.SourceID),
		SourceType:  siteDO.SourceType,
		Url:         siteDO.Url,
		Description: siteDO.Description,
		Tag:         siteDO.Tag,
	}
}
