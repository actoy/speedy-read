package app

import (
	"context"
	"speedy/read/biz/domain/aggregates/article_summary"
	"speedy/read/biz/domain/service"
	articleInfra "speedy/read/biz/infra/repository/article"
	articleSummaryInfra "speedy/read/biz/infra/repository/article_summary"
)

type ArticleSummaryApplicationI interface {
	CreateArticle(ctx context.Context, articleSummary *article_summary.ArticleSummary) (int64, error)
	GetArticleSummaryList(ctx context.Context, limit, offSet int32) ([]*article_summary.ArticleSummary, error)
}

type ArticleSummaryApplication struct {
	articleSummarySvc service.ArticleSummaryServiceI
}

func NewArticleSummaryApplication() ArticleSummaryApplicationI {
	return &ArticleSummaryApplication{
		articleSummarySvc: service.NewArticleSummaryService(articleSummaryInfra.NewArticleSummaryRepository(),
			articleSummaryInfra.NewLabelRepository(), articleInfra.NewArticleRepository()),
	}
}

func (impl *ArticleSummaryApplication) CreateArticle(ctx context.Context, articleSummary *article_summary.ArticleSummary) (int64, error) {
	return impl.articleSummarySvc.CreateArticleSummary(ctx, articleSummary)
}

func (impl *ArticleSummaryApplication) GetArticleSummaryList(ctx context.Context, limit, offSet int32) ([]*article_summary.ArticleSummary, error) {
	return impl.articleSummarySvc.GetArticleSummaryList(ctx, article_summary.SummaryListParams{
		Limit:  limit,
		OffSet: offSet,
	})
}
