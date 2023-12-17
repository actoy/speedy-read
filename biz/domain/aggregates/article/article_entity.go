package article

import (
	"speedy/read/biz/domain/aggregates/site"
	"time"
)

type Article struct {
	ID         int64
	Author     *Author
	SourceSite *site.Site
	Language   string
	PublishAt  time.Time
	Url        string
	Type       string
	Title      string
	Content    string
	Status     int32
	Score      int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

const (
	StatusInit = iota + 1
	StatusReject
	StatusPass
)

type Author struct {
	ID         int64
	Url        string
	AuthorName string
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
