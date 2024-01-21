package site

import "time"

const (
	SeekingAlphaTag = "SeekingAlpha"

	SiteTypeRss  = "rss"
	SiteTypeCraw = "craw"
)

type Site struct {
	ID          int64
	SiteMeta    *SiteMeta
	Url         string
	Description string
	Tag         string
	Type        string
	TypeKey     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	StockMeteType = "stock"
)

type SiteMeta struct {
	ID        int64
	SiteID    int64
	MetaType  string
	MetaKey   string
	MetaValue string
	CreatedAt time.Time
	UpdatedAt time.Time
}
