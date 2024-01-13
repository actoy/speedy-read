package article_summary

import (
	"speedy/read/biz/domain/aggregates/article"
	"time"
)

type ArticleSummary struct {
	ID              int64
	Article         *article.Article
	LabelList       []*Label
	Title           string
	Summary         string
	ContentSummary  *SummaryContent
	Outline         []*SummaryOutline
	TradingProposal int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Label struct {
	ID          int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SummaryContent struct {
	ID          int64
	SummaryID   int64
	Original    string
	Translation string
}

type SummaryOutline struct {
	ID        int64
	SummaryID int64
	Title     string
	Content   string
}
