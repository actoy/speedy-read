package article_summary

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
	"time"
)

type ArticleSummaryRepo struct {
}

func (r *ArticleSummaryRepo) Save(ctx context.Context, articleSummaryPO *ArticleSummary) (int64, error) {
	if articleSummaryPO == nil || articleSummaryPO.ArticleID == int64(0) {
		return int64(0), errors.New("params po is error")
	}
	existArticleSummaryPO := &ArticleSummary{}
	result := infra.DB.Where("article_id", articleSummaryPO.ArticleID).Find(&existArticleSummaryPO)
	// update
	if result.Error == nil && existArticleSummaryPO.ID != 0 {
		existArticleSummaryPO.Title = articleSummaryPO.Title
		existArticleSummaryPO.Summary = articleSummaryPO.Summary
		existArticleSummaryPO.ContentSummary = articleSummaryPO.ContentSummary
		existArticleSummaryPO.Content = articleSummaryPO.Content
		existArticleSummaryPO.Outline = articleSummaryPO.Outline
		existArticleSummaryPO.UpdatedAt = time.Now()
		infra.DB.Save(existArticleSummaryPO)
		return existArticleSummaryPO.ID, nil
	}
	articleSummaryPO.CreatedAt = time.Now()
	articleSummaryPO.UpdatedAt = time.Now()
	result = infra.DB.Create(articleSummaryPO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return articleSummaryPO.ID, nil
}

func (r *ArticleSummaryRepo) GetArticleSummaryList(ctx context.Context, limit, offSet int32) ([]*ArticleSummary, error) {
	summaryList := make([]*ArticleSummary, 0)
	result := infra.DB.Limit(int(limit)).
		Offset(int(offSet)).
		Find(&summaryList)
	if result.Error == nil {
		return summaryList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return summaryList, nil
	}
	return summaryList, result.Error
}
