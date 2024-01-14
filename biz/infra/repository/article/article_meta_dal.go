package article

import (
	"context"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
	"time"
)

type ArticleMetaRepo struct {
}

func (dal *ArticleMetaRepo) Save(ctx context.Context, articleMetaPO *ArticleMeta) (int64, error) {
	metaPO := &ArticleMeta{}
	result := infra.DB.WithContext(ctx).
		Where("article_id = ? and meta_key = ? and meta_value = ?",
			articleMetaPO.ArticleID, articleMetaPO.MetaKey, articleMetaPO.MetaValue).
		First(&metaPO)
	if result.Error == nil && metaPO.ID != 0 {
		return metaPO.ID, nil
	}
	articleMetaPO.ID = infra.IdGenerate()
	articleMetaPO.CreatedAt = time.Now()
	articleMetaPO.UpdatedAt = time.Now()
	result = infra.DB.WithContext(ctx).Create(articleMetaPO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return articleMetaPO.ID, nil
}

func (dal *ArticleMetaRepo) GetArticleMetaListByArticleID(ctx context.Context, articleID int64) ([]*ArticleMeta, error) {
	articleMetaList := make([]*ArticleMeta, 0)
	result := infra.DB.WithContext(ctx).
		Where("article_id = ?", articleID).
		Find(&articleMetaList)
	if result.Error == nil {
		return articleMetaList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}
