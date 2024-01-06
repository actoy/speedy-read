package service

import (
	"context"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/site"
	articleInfra "speedy/read/biz/infra/repository/article"
	siteInfra "speedy/read/biz/infra/repository/site"
)

type ArticleServiceI interface {
	CreateArticle(ctx context.Context, articleDO *article.Article) (int64, error)
}

type ArticleService struct {
	articleRepo article.ArticleRepo
	authorRepo  article.AuthorRepo
	siteRepo    site.SiteRepo
}

func NewArticleService() ArticleServiceI {
	return &ArticleService{
		articleRepo: articleInfra.NewArticleRepository(),
		authorRepo:  articleInfra.NewAuthorRepository(),
		siteRepo:    siteInfra.NewSiteRepository(),
	}
}

func (impl *ArticleService) CreateArticle(ctx context.Context, articleDO *article.Article) (int64, error) {
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
