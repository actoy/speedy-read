package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/site"
	articleInfra "speedy/read/biz/infra/repository/article"
	siteInfra "speedy/read/biz/infra/repository/site"
)

type ArticleServiceI interface {
	CreateArticle(ctx context.Context, articleDO *article.Article) (int64, error)
}

type ArticleService struct {
	articleRepo     article.ArticleRepo
	authorRepo      article.AuthorRepo
	siteRepo        site.SiteRepo
	articleMeteRepo article.ArticleMetaRepo
}

func NewArticleService() ArticleServiceI {
	return &ArticleService{
		articleRepo:     articleInfra.NewArticleRepository(),
		authorRepo:      articleInfra.NewAuthorRepository(),
		siteRepo:        siteInfra.NewSiteRepository(),
		articleMeteRepo: articleInfra.NewArticleMetaRepository(),
	}
}

func (impl *ArticleService) CreateArticle(ctx context.Context, articleDO *article.Article) (int64, error) {
	// create author
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
	// create site
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
	// create article
	articleID, err := impl.articleRepo.Create(ctx, articleDO)
	if err != nil {
		return int64(0), err
	}
	// create article meta
	for _, meta := range articleDO.ArticleMetaList {
		meta.ArticleID = articleID
		_, err := impl.articleMeteRepo.CreateArticleMeta(ctx, meta)
		if err != nil {
			klog.CtxErrorf(ctx, "create meta error, meteKey is ")
		}
	}
	return articleID, nil
}
