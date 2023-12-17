package article

import "context"

type ArticleRepo interface {
	Create(ctx context.Context, articleDO *Article) (int64, error)
	ArticleList(ctx context.Context) ([]*Article, error)
}

type AuthorRepo interface {
	CreateAuthor(ctx context.Context, authorDO *Author) (int64, error)
	GetAuthorByID(ctx context.Context, id int64) (*Author, error)
	GetAuthorByAuthorName(ctx context.Context, AuthorName string) (*Author, error)
}
