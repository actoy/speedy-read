package article_summary

import (
	"context"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
	"time"
)

type LabelRefRepo struct {
}

func (r *LabelRefRepo) Save(ctx context.Context, labelRefPO *LabelRef) (int64, error) {
	existLabelRefPO := &LabelRef{}
	result := infra.DB.WithContext(ctx).Where("source_id = ? and source_type = ? and label_id = ?",
		labelRefPO.SourceID, labelRefPO.SourceType, labelRefPO.LabelID).Find(&existLabelRefPO)
	if result.Error == nil && existLabelRefPO.ID != 0 {
		return existLabelRefPO.ID, nil
	}
	labelRefPO.CreatedAt = time.Now()
	labelRefPO.UpdatedAt = time.Now()
	result = infra.DB.WithContext(ctx).Create(labelRefPO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return labelRefPO.ID, nil
}

func (r *LabelRefRepo) GetLabelRefListBySourceIDs(ctx context.Context, sourceID int64, sourceType string) ([]*LabelRef, error) {
	refs := make([]*LabelRef, 0)
	result := infra.DB.WithContext(ctx).
		Where("source_id = ? AND source_type = ?", sourceID, sourceType).
		Find(&refs)
	if result.Error == nil {
		return refs, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return refs, nil
	}
	return refs, result.Error
}
