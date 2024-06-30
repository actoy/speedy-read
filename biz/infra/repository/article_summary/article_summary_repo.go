package article_summary

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/domain/aggregates/article_summary"
)

type Repository struct {
	articleSummaryRepo *ArticleSummaryRepo
	summaryContentRepo *SummaryContentRepo
	summaryOutlineRepo *SummaryOutlineRepo
	labelRepo          *LabelRepo
	labelRefRepo       *LabelRefRepo
}

func NewArticleSummaryRepository() article_summary.ArticleSummaryRepo {
	return &Repository{
		articleSummaryRepo: &ArticleSummaryRepo{},
		summaryContentRepo: &SummaryContentRepo{},
		summaryOutlineRepo: &SummaryOutlineRepo{},
	}
}

func NewLabelRepository() article_summary.LabelRepo {
	return &Repository{
		labelRepo:    &LabelRepo{},
		labelRefRepo: &LabelRefRepo{},
	}
}

func (r *Repository) CreateSummary(ctx context.Context, articleSummaryDO *article_summary.ArticleSummary) (int64, error) {
	summaryID, err := r.articleSummaryRepo.Save(ctx, convertArticleSummaryDOToPO(articleSummaryDO))
	if err != nil {
		return summaryID, err
	}
	if articleSummaryDO.ContentSummary != nil {
		err = r.summaryContentRepo.Save(ctx, convertContentSummaryDoToPo(articleSummaryDO.ContentSummary, summaryID))
		if err != nil {
			klog.CtxErrorf(ctx, "save content summary error %v", err)
		}
	}
	if len(articleSummaryDO.Outline) > 0 {
		err = r.summaryOutlineRepo.Save(ctx, convertSummaryOutlineDoToPo(articleSummaryDO.Outline, summaryID))
		if err != nil {
			klog.CtxErrorf(ctx, "save summary outline error %v", err)
		}
	}
	return summaryID, nil
}

func (r *Repository) ArticleSummaryList(ctx context.Context, params article_summary.SummaryListParams) ([]*article_summary.ArticleSummary, error) {
	summaryPOList, err := r.articleSummaryRepo.GetArticleSummaryList(ctx, params.Limit, params.OffSet, params.Symbol, params.ArticleType)
	if err != nil {
		return nil, err
	}
	summaryIDs := make([]int64, 0)
	for _, summary := range summaryPOList {
		summaryIDs = append(summaryIDs, summary.ID)
	}
	summaryOutlineList, err := r.summaryOutlineRepo.GetOutlineListBySummaryID(ctx, summaryIDs)
	if err != nil {
		klog.CtxErrorf(ctx, "get summary outline error %v", err)
		return nil, err
	}
	summaryContentList, err := r.summaryContentRepo.GetSummaryContentBySummaryID(ctx, summaryIDs)
	if err != nil {
		klog.CtxErrorf(ctx, "get summary content error %v", err)
		return nil, err
	}
	return convertArticleSummaryDO(summaryPOList, summaryContentList, summaryOutlineList), nil
}

func (r *Repository) GetLabelListBySource(ctx context.Context, sourceID int64, sourceType article_summary.SourceType) ([]*article_summary.Label, error) {
	labelPOList, err := r.labelRepo.GetLabelListBySource(ctx, sourceID, string(sourceType))
	if err != nil {
		return nil, err
	}
	return convertLabelPOListTODO(labelPOList), nil
}

func (r *Repository) CreateLabel(ctx context.Context, descriptionList []string, sourceID int64, sourceType article_summary.SourceType) error {
	r.labelRefRepo.Delete(ctx, sourceID, string(sourceType))
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

func (r *Repository) GetArticleSummaryCount(ctx context.Context, articleType string) (int32, error) {
	return r.articleSummaryRepo.GetArticleSummaryCount(ctx, articleType)
}

func (r *Repository) GetArticleSummaryByID(ctx context.Context, summaryID int64) (*article_summary.ArticleSummary, error) {
	summaryPO, err := r.articleSummaryRepo.GetArticleSummaryDetailByID(ctx, summaryID)
	if err != nil {
		klog.CtxErrorf(ctx, "get summary detail error %v", err)
		return nil, err
	}
	if summaryPO == nil {
		return &article_summary.ArticleSummary{}, nil
	}
	summaryOutlineList, err := r.summaryOutlineRepo.GetOutlineListBySummaryID(ctx, []int64{summaryPO.ID})
	if err != nil {
		klog.CtxErrorf(ctx, "get summary outline error %v", err)
		return nil, err
	}
	summaryContentList, err := r.summaryContentRepo.GetSummaryContentBySummaryID(ctx, []int64{summaryPO.ID})
	if err != nil {
		klog.CtxErrorf(ctx, "get summary content error %v", err)
		return nil, err
	}
	return convertArticleSummaryPOToDO(summaryPO, summaryContentList, summaryOutlineList), nil
}
