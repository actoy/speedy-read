package article_summary

import "time"

type ArticleSummary struct {
	ID             int64
	ArticleID      int64
	Title          string
	Summary        string
	ContentSummary string
	Outline        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (ArticleSummary) TableName() string {
	return "article_summarys"
}

type Label struct {
	ID          int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type LabelRef struct {
	ID         int64
	SourceID   int64
	SourceType string
	LabelID    int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
