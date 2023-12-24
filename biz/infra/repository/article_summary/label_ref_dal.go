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
	// todo: 重复插入
	labelRefPO.CreatedAt = time.Now()
	labelRefPO.UpdatedAt = time.Now()
	result := infra.DB.Create(labelRefPO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return labelRefPO.ID, nil
}

func (r *LabelRefRepo) GetLabelRefListBySourceIDs(ctx context.Context, sourceID int64, sourceType string) ([]*LabelRef, error) {
	refs := make([]*LabelRef, 0)
	result := infra.DB.
		Where("source_id = ? AND source_type = ?", sourceID, sourceType).
		Find(&refs)
	if result.Error == nil {
		return refs, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return refs, nil
	}
	return refs, result.Error
}
