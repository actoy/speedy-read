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

func (dal *ArticleRepo) GetArticleByUrl(ctx context.Context, url string) (*Article, error) {
	article := &Article{}
	result := infra.DB.WithContext(ctx).Where("url = ?", url).First(&article)
	if result.Error == nil {
		return article, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *ArticleRepo) Save(ctx context.Context, articlePO *Article) (int64, error) {
	existArticle, _ := dal.GetArticleByUrl(ctx, articlePO.Url)
	if existArticle != nil {
		return existArticle.ID, nil
	}
	articlePO.ID = infra.IdGenerate()
	articlePO.CreatedAt = time.Now()
	articlePO.UpdatedAt = time.Now()
	result := infra.DB.WithContext(ctx).Create(articlePO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return articlePO.ID, nil
}

func (dal *ArticleRepo) GetArticleList(ctx context.Context, params article.ArticleListParams) ([]*Article, error) {
	articleList := make([]*Article, 0)
	result := infra.DB.WithContext(ctx).Where("status = ?", article.StatusInit)
	if len(params.ArticleType) != 0 {
		result = result.Where("type = ?", params.ArticleType)
	}
	if len(params.SiteIdList) > 0 {
		result = result.Where("source_site_id in ?", params.SiteIdList)
	}
	result = result.Limit(int(params.Limit)).
		Offset(int(params.OffSet * params.Limit)).
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
		Update("status", article.StatusReject).
		Update("updated_at", time.Now())
	if result.Error == nil {
		return nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return result.Error
}

func (dal *ArticleRepo) SetStatusPass(ctx context.Context, articleID int64, content string) error {
	result := infra.DB.WithContext(ctx).Model(&Article{}).
		Where("id = ?", articleID).
		Update("status", article.StatusPass).
		Update("content", content).
		Update("updated_at", time.Now())
	if result.Error == nil {
		return nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return result.Error
}

func (dal *ArticleRepo) GetArticleCount(ctx context.Context, status int32, params article.ArticleListParams) (int32, error) {
	var count int64
	result := infra.DB.WithContext(ctx).Model(&Article{}).
		Where("status = ?", status)
	if len(params.ArticleType) != 0 {
		result = result.Where("type = ?", params.ArticleType)
	}
	if len(params.SiteIdList) > 0 {
		result = result.Where("source_site_id in ?", params.SiteIdList)
	}
	result.Count(&count)
	if result.Error != nil {
		return int32(0), nil
	}
	return int32(count), result.Error
}
