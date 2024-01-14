package site

import (
	"speedy/read/biz/infra"
)

type Site struct {
	infra.Model
	Url         string
	Description string
	Tag         string
}

type SiteMeta struct {
	infra.Model
	SiteID    int64
	MetaType  string
	MetaKey   string
	MetaValue string
}

func (SiteMeta) TableName() string {
	return "site_metas"
}
