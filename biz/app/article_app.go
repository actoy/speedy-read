package app

import (
	"context"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/service"
	articleInfra "speedy/read/biz/infra/repository/article"
)

type ArticleApplicationI interface {
	CreateArticle(ctx context.Context, articleDO *article.Article) (int64, error)
	GetArticleList(ctx context.Context, limit, offSet int32) ([]*article.Article, error)
	RejectArticle(ctx context.Context, articleID int64) error
	ArticleCount(ctx context.Context, status int32) (int32, error)
}

type ArticleApplication struct {
	articleRepo article.ArticleRepo
	articleSvc  service.ArticleServiceI
}

func NewArticleApplication() ArticleApplicationI {
	return &ArticleApplication{
		articleRepo: articleInfra.NewArticleRepository(),
		articleSvc:  service.NewArticleService(),
	}
}

func (impl *ArticleApplication) CreateArticle(ctx context.Context, articleDO *article.Article) (int64, error) {
	return impl.articleSvc.CreateArticle(ctx, articleDO)
}

func (impl *ArticleApplication) GetArticleList(ctx context.Context, limit, offSet int32) ([]*article.Article, error) {
	return impl.articleRepo.ArticleList(ctx, limit, offSet)
}

func (impl *ArticleApplication) RejectArticle(ctx context.Context, articleID int64) error {
	return impl.articleRepo.SetStatusReject(ctx, articleID)
}

func (impl *ArticleApplication) ArticleCount(ctx context.Context, status int32) (int32, error) {
	return impl.articleRepo.GetArticleCount(ctx, status)
}
