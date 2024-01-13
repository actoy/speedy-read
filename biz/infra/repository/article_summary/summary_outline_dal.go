package article_summary

import (
	"context"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
)

type SummaryOutlineRepo struct {
}

func (r *SummaryOutlineRepo) Save(ctx context.Context, outlineList []*SummaryOutline) error {
	if len(outlineList) == 0 {
		return nil
	}
	infra.DB.Where("summary_id = ?", outlineList[0].SummaryID).Delete(&SummaryOutline{})
	result := infra.DB.Create(outlineList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *SummaryOutlineRepo) GetOutlineListBySummaryID(ctx context.Context, summaryIDs []int64) ([]*SummaryOutline, error) {
	outlineList := make([]*SummaryOutline, 0)
	result := infra.DB.WithContext(ctx).Where("summary_id in ?", summaryIDs).Find(&outlineList)
	if result.Error == nil {
		return outlineList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return outlineList, nil
	}
	return outlineList, result.Error
}
