package article_summary

import "context"

type SummaryListParams struct {
	Limit  int32
	OffSet int32
}

type ArticleSummaryRepo interface {
	CreateSummary(ctx context.Context, articleSummaryDO *ArticleSummary) (int64, error)
	ArticleSummaryList(ctx context.Context, params SummaryListParams) ([]*ArticleSummary, error)
	GetArticleSummaryCount(ctx context.Context) (int32, error)
}

type SourceType string

const (
	SourceTypeSummary = "article_summary"
)

type LabelRepo interface {
	GetLabelListBySource(ctx context.Context, sourceID int64, sourceType SourceType) ([]*Label, error)
	CreateLabel(ctx context.Context, descriptionList []string, sourceID int64, sourceType SourceType) error
}
