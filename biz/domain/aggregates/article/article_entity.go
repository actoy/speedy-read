package article

import (
	"speedy/read/biz/domain/aggregates/site"
	"time"
)

type Article struct {
	ID              int64
	Author          *Author
	SourceSite      *site.Site
	ArticleMetaList []*ArticleMeta
	Language        string
	PublishAt       time.Time
	Url             string
	Type            string
	Title           string
	Content         string
	Status          int32
	Score           int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

const (
	StatusInit = iota + 1
	StatusReject
	StatusPass

	TypeArticle = "article"
	TypeNew     = "new"
)

type Author struct {
	ID         int64
	Url        string
	AuthorName string
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

const (
	StockMeteType = "stock"
)

type ArticleMeta struct {
	ID        int64
	ArticleID int64
	MetaType  string
	MetaKey   string
	MetaValue string
	CreatedAt time.Time
	UpdatedAt time.Time
}
