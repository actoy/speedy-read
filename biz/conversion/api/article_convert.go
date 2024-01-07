package api

import (
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/utils"
	"speedy/read/kitex_gen/speedy_read"
	"time"
)

// goverter:converter
// goverter:extend TimeToString
// goverter:extend StringToTime
// goverter:extend Int64ToString
// goverter:extend StringToInt64
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

func Int64ToString(i int64) string {
	return utils.Int64ToString(i)
}

func StringToInt64(s string) int64 {
	return utils.StringToInt64(s)
}
