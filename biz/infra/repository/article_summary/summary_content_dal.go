package article_summary

import (
	"context"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
)

type SummaryContentRepo struct {
}

func (r *SummaryContentRepo) Save(ctx context.Context, summaryContent *SummaryContent) error {
	infra.DB.Where("summary_id = ?", summaryContent.SummaryID).Delete(&SummaryContent{})
	result := infra.DB.Create(summaryContent)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *SummaryContentRepo) GetSummaryContentBySummaryID(ctx context.Context, summaryIDs []int64) ([]*SummaryContent, error) {
	summaryContentList := make([]*SummaryContent, 0)
	result := infra.DB.WithContext(ctx).Where("summary_id in ?", summaryIDs).Find(&summaryContentList)
	if result.Error == nil {
		return summaryContentList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return summaryContentList, nil
	}
	return summaryContentList, result.Error
}
