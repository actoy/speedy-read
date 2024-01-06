package service

import "speedy/read/biz/domain/aggregates/article_summary"

type DataCrawingServiceI interface {
	CreateArticleSummary(ctx context.Context, articleSummaryDO *article_summary.ArticleSummary) (int64, error)
	GetArticleSummaryList(ctx context.Context, params article_summary.SummaryListParams) (
		resp []*article_summary.ArticleSummary, err error)
}
