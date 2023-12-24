package article_summary

import (
	"context"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
	"time"
)

type LabelRepo struct {
}

func (r *LabelRepo) Save(ctx context.Context, labelPO *Label) (int64, error) {
	existLabelPO := &Label{}
	result := infra.DB.Where("description", labelPO.Description).Find(&existLabelPO)
	if result.Error == nil && existLabelPO.ID != 0 {
		return existLabelPO.ID, nil
	}
	labelPO.CreatedAt = time.Now()
	labelPO.UpdatedAt = time.Now()
	result = infra.DB.Create(labelPO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return labelPO.ID, nil
}

func (r *LabelRepo) GetLabelListBySource(ctx context.Context, sourceID int64, sourceType string) ([]*Label, error) {
	labelList := make([]*Label, 0)
	result := infra.DB.Table("labels").Select("labels.description").
		Joins("inner join label_refs on labels.id = label_ref.label_id").
		Where("source_id = ? AND source_type = ?", sourceID, sourceType).
		Find(&labelList)
	if result.Error == nil {
		return labelList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return labelList, nil
	}
	return labelList, result.Error
}
