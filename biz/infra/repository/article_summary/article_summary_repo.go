package article_summary

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/domain/aggregates/article_summary"
)

type Repository struct {
	articleSummaryRepo *ArticleSummaryRepo
	labelRepo          *LabelRepo
	labelRefRepo       *LabelRefRepo
}

func NewArticleSummaryRepository() article_summary.ArticleSummaryRepo {
	return &Repository{
		articleSummaryRepo: &ArticleSummaryRepo{},
	}
}

func NewLabelRepository() article_summary.LabelRepo {
	return &Repository{
		labelRepo:    &LabelRepo{},
		labelRefRepo: &LabelRefRepo{},
	}
}

func (r *Repository) CreateSummary(ctx context.Context, articleSummaryDO *article_summary.ArticleSummary) (int64, error) {
	return r.articleSummaryRepo.Save(ctx, ConvertArticleSummaryDOToPO(articleSummaryDO))
}

func (r *Repository) ArticleSummaryList(ctx context.Context, params article_summary.SummaryListParams) ([]*article_summary.ArticleSummary, error) {
	summaryPOList, err := r.articleSummaryRepo.GetArticleSummaryList(ctx, params.Limit, params.OffSet)
	if err != nil {
		return nil, err
	}
	return ConvertArticleSummaryPOListTODO(summaryPOList), nil
}

func (r *Repository) GetLabelListBySource(ctx context.Context, sourceID int64, sourceType article_summary.SourceType) ([]*article_summary.Label, error) {
	labelPOList, err := r.labelRepo.GetLabelListBySource(ctx, sourceID, string(sourceType))
	if err != nil {
		return nil, err
	}
	return ConvertLabelPOListTODO(labelPOList), nil
}

func (r *Repository) CreateLabel(ctx context.Context, descriptionList []string, sourceID int64, sourceType article_summary.SourceType) error {
	for _, description := range descriptionList {
		id, err := r.labelRepo.Save(ctx, &Label{Description: description})
		if err != nil {
			klog.CtxErrorf(ctx, "save label error %v", err)
			continue
		}
		r.labelRefRepo.Save(ctx, &LabelRef{
			SourceID:   sourceID,
			SourceType: string(sourceType),
			LabelID:    id,
		})
	}
	return nil
}

func (r *Repository) GetArticleSummaryCount(ctx context.Context) (int32, error) {
	return r.articleSummaryRepo.GetArticleSummaryCount(ctx)
}
