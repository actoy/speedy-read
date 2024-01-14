package article

import (
	"speedy/read/biz/infra"
	"speedy/read/biz/infra/repository/site"
	"time"
)

type Article struct {
	infra.Model
	AuthorID     int64
	Author       *Author
	SourceSiteID int64
	SourceSite   *site.Site
	Language     string
	PublishAt    time.Time
	Url          string
	Type         string
	Title        string
	Content      string
	Status       int32
	Score        int32
}

type Author struct {
	infra.Model
	Url        string
	AuthorName string
	Image      string
}

type ArticleMeta struct {
	infra.Model
	ArticleID int64
	MetaType  string
	MetaKey   string
	MetaValue string
}

func (ArticleMeta) TableName() string {
	return "article_metas"
}
