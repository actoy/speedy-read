package article

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/infra/repository/site"
)

type Repository struct {
	articleRepo *ArticleRepo
	authorRepo  *AuthorRepo
	siteRepo    *site.SiteRepo
}

func NewArticleRepository() article.ArticleRepo {
	return &Repository{
		articleRepo: &ArticleRepo{},
		authorRepo:  &AuthorRepo{},
		siteRepo:    &site.SiteRepo{},
	}
}

func NewAuthorRepository() article.AuthorRepo {
	return &Repository{
		authorRepo: &AuthorRepo{},
	}
}

func (r *Repository) Create(ctx context.Context, articleDO *article.Article) (int64, error) {
	return r.articleRepo.Save(ctx, ConvertArticleDOToPO(articleDO))
}

func (r *Repository) ArticleList(ctx context.Context) ([]*article.Article, error) {
	articlePOList, err := r.articleRepo.GetArticleList(ctx)
	if err != nil {
		return nil, err
	}
	articleList := make([]*article.Article, 0)
	for _, po := range articlePOList {
		authorPO, err := r.authorRepo.GetAuthorByID(ctx, po.AuthorID)
		if err != nil {
			klog.Error(ctx, "get author by id error: %v", err)
			continue
		}
		sitePO, err := r.siteRepo.GetSiteByID(ctx, po.SourceSiteID)
		if err != nil {
			klog.Error(ctx, "get author by id error: %v", err)
			continue
		}
		articleList = append(articleList, ConvertArticlePOToDO(po, authorPO, sitePO))
	}
	return articleList, nil
}

func (r *Repository) GetAuthorByAuthorName(ctx context.Context, AuthorName string) (*article.Author, error) {
	author, err := r.authorRepo.GetAuthorByAuthorName(ctx, AuthorName)
	if err != nil {
		return nil, err
	}
	return ConvertAuthorPOToDO(author), nil
}

func (r *Repository) GetAuthorByID(ctx context.Context, id int64) (*article.Author, error) {
	author, err := r.authorRepo.GetAuthorByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ConvertAuthorPOToDO(author), nil
}

func (r *Repository) CreateAuthor(ctx context.Context, authorDO *article.Author) (int64, error) {
	return r.authorRepo.Save(ctx, ConvertAuthorDOToPO(authorDO))
}
