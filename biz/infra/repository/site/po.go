package site

import (
	"time"
)

type Site struct {
	ID          int64
	Url         string
	Description string
	Tag         string
	Type        string
	TypeKey     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SiteMeta struct {
	ID        int64
	SiteID    int64
	MetaType  string
	MetaKey   string
	MetaValue string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SiteMeta) TableName() string {
	return "site_metas"
}
