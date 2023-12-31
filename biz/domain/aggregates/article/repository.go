package article

import "context"

type ArticleRepo interface {
	Create(ctx context.Context, articleDO *Article) (int64, error)
	ArticleList(ctx context.Context, limit, offSet int32) ([]*Article, error)
	SetStatusReject(ctx context.Context, articleID int64) error
	GetArticleByID(ctx context.Context, articleID int64) (*Article, error)
	SetStatusPass(ctx context.Context, articleID int64, content string) error
	GetArticleCount(ctx context.Context, status int32) (int32, error)
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
