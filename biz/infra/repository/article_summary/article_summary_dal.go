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
	result := infra.DB.WithContext(ctx).Where("article_id", articleSummaryPO.ArticleID).Find(&existArticleSummaryPO)
	// update
	if result.Error == nil && existArticleSummaryPO.ID != 0 {
		existArticleSummaryPO.Title = articleSummaryPO.Title
		existArticleSummaryPO.Summary = articleSummaryPO.Summary
		existArticleSummaryPO.TradingProposal = articleSummaryPO.TradingProposal
		existArticleSummaryPO.UpdatedAt = time.Now()
		infra.DB.Save(existArticleSummaryPO)
		return existArticleSummaryPO.ID, nil
	}
	articleSummaryPO.ID = infra.IdGenerate()
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
	result := infra.DB.WithContext(ctx).Limit(int(limit)).
		Offset(int(offSet * limit)).
		Order("created_at").
		Find(&summaryList)
	if result.Error == nil {
		return summaryList, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return summaryList, nil
	}
	return summaryList, result.Error
}

func (r *ArticleSummaryRepo) GetArticleSummaryCount(ctx context.Context) (int32, error) {
	var count int64
	result := infra.DB.WithContext(ctx).Model(&ArticleSummary{}).Count(&count)
	if result.Error != nil {
		return int32(0), nil
	}
	return int32(count), result.Error
}

func (r *ArticleSummaryRepo) GetArticleSummaryDetailByID(ctx context.Context, id int64) (*ArticleSummary, error) {
	var summaryPO *ArticleSummary
	result := infra.DB.WithContext(ctx).First(&summaryPO, id)
	if result.Error == nil {
		return summaryPO, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}
