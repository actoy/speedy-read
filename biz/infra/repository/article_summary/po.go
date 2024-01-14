package article_summary

import (
	"speedy/read/biz/infra"
)

type ArticleSummary struct {
	infra.Model
	ArticleID       int64
	Title           string
	Summary         string
	TradingProposal int32
}

func (ArticleSummary) TableName() string {
	return "article_summarys"
}

type Label struct {
	infra.Model
	Description string
}

type LabelRef struct {
	infra.Model
	SourceID   int64
	SourceType string
	LabelID    int64
}

type SummaryContent struct {
	infra.Model
	SummaryID   int64
	Original    string
	Translation string
}

func (SummaryContent) TableName() string {
	return "summary_contents"
}

type SummaryOutline struct {
	infra.Model
	SummaryID int64
	Title     string
	Content   string
}

func (SummaryOutline) TableName() string {
	return "summary_outlines"
}
