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

func (r *Repository) ArticleList(ctx context.Context, limit, offSet int32) ([]*article.Article, error) {
	articlePOList, err := r.articleRepo.GetArticleList(ctx, limit, offSet)
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

func (r *Repository) GetArticleByID(ctx context.Context, articleID int64) (*article.Article, error) {
	articlePOList, err := r.articleRepo.GetArticleListByIDs(ctx, []int64{articleID})
	if err != nil {
		return nil, err
	}
	if len(articlePOList) == 0 {
		return nil, nil
	}
	articlePO := articlePOList[0]
	authorPO, err := r.authorRepo.GetAuthorByID(ctx, articlePO.AuthorID)
	if err != nil {
		klog.Error(ctx, "get author by id error: %v", err)
		return nil, err
	}
	sitePO, err := r.siteRepo.GetSiteByID(ctx, articlePO.SourceSiteID)
	if err != nil {
		klog.Error(ctx, "get site by id error: %v", err)
		return nil, err
	}

	return ConvertArticlePOToDO(articlePO, authorPO, sitePO), nil
}

func (r *Repository) SetStatusReject(ctx context.Context, articleID int64) error {
	return r.articleRepo.SetStatusReject(ctx, articleID)
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
