package service

import (
	"context"
	"speedy/read/biz/infra"
	"testing"
)

func TestCrawData(t *testing.T) {
	infra.Init()
	crawSvc := NewDataCrawService()
	_ = crawSvc.CrawArticle(context.Background())
}
