package app

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/domain/aggregates/article_summary"
	"speedy/read/biz/domain/service"
	articleInfra "speedy/read/biz/infra/repository/article"
	articleSummaryInfra "speedy/read/biz/infra/repository/article_summary"
	"speedy/read/biz/utils"
)

type ArticleSummaryApplicationI interface {
	CreateArticle(ctx context.Context, articleSummary *article_summary.ArticleSummary, content string) (int64, error)
	GetArticleSummaryList(ctx context.Context, limit, offSet int32) ([]*article_summary.ArticleSummary, error)
	GetArticleSummaryDetailByID(ctx context.Context, summaryID string) (*article_summary.ArticleSummary, error)
	ArticleSummaryCount(ctx context.Context) int32
}

type ArticleSummaryApplication struct {
	articleSummarySvc  service.ArticleSummaryServiceI
	articleSummaryRepo article_summary.ArticleSummaryRepo
}

func NewArticleSummaryApplication() ArticleSummaryApplicationI {
	articleSummaryRepo := articleSummaryInfra.NewArticleSummaryRepository()
	return &ArticleSummaryApplication{
		articleSummaryRepo: articleSummaryRepo,
		articleSummarySvc: service.NewArticleSummaryService(articleSummaryRepo,
			articleSummaryInfra.NewLabelRepository(), articleInfra.NewArticleRepository()),
	}
}

func (impl *ArticleSummaryApplication) CreateArticle(ctx context.Context, articleSummary *article_summary.ArticleSummary, content string) (int64, error) {
	return impl.articleSummarySvc.CreateArticleSummary(ctx, articleSummary, content)
}

func (impl *ArticleSummaryApplication) GetArticleSummaryList(ctx context.Context, limit, offSet int32) ([]*article_summary.ArticleSummary, error) {
	return impl.articleSummarySvc.GetArticleSummaryList(ctx, article_summary.SummaryListParams{
		Limit:  limit,
		OffSet: offSet,
	})
}

func (impl *ArticleSummaryApplication) ArticleSummaryCount(ctx context.Context) int32 {
	count, err := impl.articleSummaryRepo.GetArticleSummaryCount(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "get article summary count error is %v", err)
		return int32(0)
	}
	return count
}

func (impl *ArticleSummaryApplication) GetArticleSummaryDetailByID(ctx context.Context, summaryID string) (*article_summary.ArticleSummary, error) {
	return impl.articleSummarySvc.GetArticleSummaryDetailByID(ctx, utils.StringToInt64(summaryID))
}
