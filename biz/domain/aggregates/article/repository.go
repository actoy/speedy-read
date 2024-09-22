package article

import "context"

type ArticleListParams struct {
	SymbolIdList []int64
	SiteIdList   []int64
	ArticleType  string
	Limit        int32
	OffSet       int32
}

type ArticleRepo interface {
	Create(ctx context.Context, articleDO *Article) (int64, error)
	ArticleList(ctx context.Context, params ArticleListParams) ([]*Article, error)
	SetStatusReject(ctx context.Context, articleID int64) error
	GetArticleByID(ctx context.Context, articleIds []int64) ([]*Article, error)
	SetStatusPass(ctx context.Context, articleID int64, content string) error
	GetArticleCount(ctx context.Context, status int32, params ArticleListParams) (int32, error)
	GetArticleByUrl(ctx context.Context, url string) (*Article, error)
}

type AuthorRepo interface {
	CreateAuthor(ctx context.Context, authorDO *Author) (int64, error)
	GetAuthorByID(ctx context.Context, id int64) (*Author, error)
	GetAuthorByAuthorName(ctx context.Context, AuthorName string) (*Author, error)
}

type ArticleMetaRepo interface {
	CreateArticleMeta(ctx context.Context, articleMetaOD *ArticleMeta) (int64, error)
	GetArticleMetaByArticleID(ctx context.Context, articleID int64) ([]*ArticleMeta, error)
}
