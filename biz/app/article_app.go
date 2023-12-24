package app

import (
	"context"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/site"
	articleInfra "speedy/read/biz/infra/repository/article"
	siteInfra "speedy/read/biz/infra/repository/site"
)

type ArticleApplicationI interface {
	CreateArticle(ctx context.Context, articleDO *article.Article) (int64, error)
	GetArticleList(ctx context.Context, limit, offSet int32) ([]*article.Article, error)
	RejectArticle(ctx context.Context, articleID int64) error
}

type ArticleApplication struct {
	articleRepo article.ArticleRepo
	authorRepo  article.AuthorRepo
	siteRepo    site.SiteRepo
}

func NewArticleApplication() ArticleApplicationI {
	return &ArticleApplication{
		articleRepo: articleInfra.NewArticleRepository(),
		authorRepo:  articleInfra.NewAuthorRepository(),
		siteRepo:    siteInfra.NewSiteRepository(),
	}
}

func (impl *ArticleApplication) CreateArticle(ctx context.Context, articleDO *article.Article) (int64, error) {
	author, err := impl.authorRepo.GetAuthorByAuthorName(ctx, articleDO.Author.AuthorName)
	if err != nil {
		return int64(0), err
	}
	if author == nil {
		authorID, err := impl.authorRepo.CreateAuthor(ctx, articleDO.Author)
		if err != nil {
			return int64(0), err
		}
		articleDO.Author.ID = authorID
	} else {
		articleDO.Author.ID = author.ID
	}
	site, err := impl.siteRepo.GetSiteByUrl(ctx, articleDO.SourceSite.Url)
	if err != nil {
		return int64(0), err
	}
	if site == nil {
		siteID, err := impl.siteRepo.Create(ctx, articleDO.SourceSite)
		if err != nil {
			return int64(0), err
		}
		articleDO.SourceSite.ID = siteID
	} else {
		articleDO.SourceSite.ID = site.ID
	}

	return impl.articleRepo.Create(ctx, articleDO)
}

func (impl *ArticleApplication) GetArticleList(ctx context.Context, limit, offSet int32) ([]*article.Article, error) {
	return impl.articleRepo.ArticleList(ctx, limit, offSet)
}

func (impl *ArticleApplication) RejectArticle(ctx context.Context, articleID int64) error {
	return impl.articleRepo.SetStatusReject(ctx, articleID)
}
