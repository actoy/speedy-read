package site

import (
	"context"
)

type SiteRepo interface {
	Create(ctx context.Context, siteDO *Site) (int64, error)
	GetSiteList(ctx context.Context) ([]*Site, error)
	GetSiteByUrl(ctx context.Context, url string) (*Site, error)
}
