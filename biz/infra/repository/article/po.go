package article

import (
	"speedy/read/biz/infra/repository/site"
	"time"
)

type Article struct {
	ID           int64
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
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Author struct {
	ID         int64
	Url        string
	AuthorName string
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ArticleMeta struct {
	ID        int64
	ArticleID int64
	MetaType  string
	MetaKey   string
	MetaValue string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ArticleMeta) TableName() string {
	return "article_metas"
}
