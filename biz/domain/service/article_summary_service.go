package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/article_summary"
)

type ArticleSummaryServiceI interface {
	CreateArticleSummary(ctx context.Context, articleSummaryDO *article_summary.ArticleSummary, content string) (int64, error)
	GetArticleSummaryList(ctx context.Context, params article_summary.SummaryListParams) (
		resp []*article_summary.ArticleSummary, err error)
}

type ArticleSummaryService struct {
	summaryRepo article_summary.ArticleSummaryRepo
	labelRepo   article_summary.LabelRepo
	articleRepo article.ArticleRepo
}

func NewArticleSummaryService(summaryRepo article_summary.ArticleSummaryRepo,
	labelRepo article_summary.LabelRepo, articleRepo article.ArticleRepo) ArticleSummaryServiceI {
	return &ArticleSummaryService{
		summaryRepo: summaryRepo,
		labelRepo:   labelRepo,
		articleRepo: articleRepo,
	}
}

func (impl *ArticleSummaryService) CreateArticleSummary(ctx context.Context,
	articleSummaryDO *article_summary.ArticleSummary, content string) (int64, error) {
	id, err := impl.summaryRepo.CreateSummary(ctx, articleSummaryDO)
	if err != nil {
		return int64(0), err
	}
	description := make([]string, 0)
	for _, label := range articleSummaryDO.LabelList {
		if label == nil {
			continue
		}
		description = append(description, label.Description)
	}
	err = impl.labelRepo.CreateLabel(ctx, description, id, article_summary.SourceTypeSummary)
	if err != nil {
		return int64(0), err
	}
	impl.articleRepo.SetStatusPass(ctx, articleSummaryDO.Article.ID, content)
	return id, nil
}

func (impl *ArticleSummaryService) GetArticleSummaryList(ctx context.Context, params article_summary.SummaryListParams) (
	resp []*article_summary.ArticleSummary, err error) {
	resp, err = impl.summaryRepo.ArticleSummaryList(ctx, params)
	if err != nil {
		return resp, err
	}
	for _, summary := range resp {
		labelList, err := impl.labelRepo.GetLabelListBySource(ctx, summary.ID, article_summary.SourceTypeSummary)
		if err != nil {
			klog.CtxErrorf(ctx, "get label list error %v", err)
			continue
		}
		summary.LabelList = labelList
		if summary.Article != nil {
			article, err := impl.articleRepo.GetArticleByID(ctx, summary.Article.ID)
			if err != nil {
				klog.CtxErrorf(ctx, "get label list error %v", err)
				continue
			}
			summary.Article = article
		}
	}
	return resp, nil
}
