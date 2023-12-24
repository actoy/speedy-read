package article

import (
	"context"
	"gorm.io/gorm"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/infra"
	"time"
)

type ArticleRepo struct {
}

func (dal *ArticleRepo) Save(ctx context.Context, articlePO *Article) (int64, error) {
	articlePO.CreatedAt = time.Now()
	articlePO.UpdatedAt = time.Now()
	result := infra.DB.WithContext(ctx).Create(articlePO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return articlePO.ID, nil
}

func (dal *ArticleRepo) GetArticleList(ctx context.Context, limit, offSet int32) ([]*Article, error) {
	articleList := make([]*Article, 0)
	result := infra.DB.WithContext(ctx).
		Where("status =?", article.StatusInit).
		Limit(int(limit)).
		Offset(int(offSet)).
		Order("created_at").
		Find(&articleList)
	if result.Error == nil {
		return articleList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *ArticleRepo) GetArticleListByIDs(ctx context.Context, articleIDs []int64) ([]*Article, error) {
	articleList := make([]*Article, 0)
	result := infra.DB.WithContext(ctx).Find(&articleList, articleIDs)
	if result.Error == nil {
		return articleList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *ArticleRepo) SetStatusReject(ctx context.Context, articleID int64) error {
	result := infra.DB.WithContext(ctx).Model(&Article{}).
		Where("id = ?", articleID).
		Update("status", article.StatusReject)
	if result.Error == nil {
		return nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return result.Error
}
