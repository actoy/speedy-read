package article

import (
	"context"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
	"time"
)

type ArticleRepo struct {
}

func (dal *ArticleRepo) Save(ctx context.Context, articlePO *Article) (int64, error) {
	articlePO.CreatedAt = time.Now()
	articlePO.UpdatedAt = time.Now()
	result := infra.DB.Create(articlePO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return articlePO.ID, nil
}

func (dal *ArticleRepo) GetArticleList(ctx context.Context) ([]*Article, error) {
	articleList := make([]*Article, 0)
	result := infra.DB.Find(&articleList)
	if result.Error == nil {
		return articleList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}
