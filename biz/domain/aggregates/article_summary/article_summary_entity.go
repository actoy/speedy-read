package article_summary

import (
	"speedy/read/biz/domain/aggregates/article"
	"time"
)

type ArticleSummary struct {
	ID             int64
	Article        *article.Article
	LabelList      []*Label
	Title          string
	Summary        string
	ContentSummary string
	Outline        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Label struct {
	ID          int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
