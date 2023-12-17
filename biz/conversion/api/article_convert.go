package api

import (
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/kitex_gen/speedy_read"
	"time"
)

// goverter:converter
// goverter:extend TimeToString
// goverter:extend StringToTime
type ArticleConvert interface {
	// goverter:map SourceSite Site
	ArticleDOToThrift(articleDO *article.Article) *speedy_read.Article
	// goverter:map Site SourceSite
	ArticleThriftToDO(article *speedy_read.Article) *article.Article
}

func TimeToString(t time.Time) string {
	return t.String()
}

func StringToTime(t string) time.Time {
	return time.Now()
}
