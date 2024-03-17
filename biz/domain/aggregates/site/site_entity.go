package site

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

const (
	SeekingAlphaTag = "SeekingAlpha"
	FoolTag         = "Fool"
	TheStreetTag    = "TheStreet"

	SiteTypeRss        = "rss"
	SiteTypeCraw       = "craw"
	SiteTypeCrawDetail = "craw_detail"

	SiteTaskTypeList   = "list"
	SiteTaskTypeDetail = "detail"
	SiteTaskTypeAll    = "all" // list & detail
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

type SiteTypeKey struct {
	TaskID   string
	TaskType string
}

func (site *Site) GetTypeKeyList(ctx context.Context) []SiteTypeKey {
	result := make([]SiteTypeKey, 0)
	if len(site.TypeKey) == 0 {
		return result
	}
	err := json.Unmarshal([]byte(site.TypeKey), &result)
	if err != nil {
		klog.CtxErrorf(ctx, "Error get site type key err %v", err)
		return []SiteTypeKey{}
	}
	return result
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
