package article

import "context"

type ArticleRepo interface {
	Create(ctx context.Context)
}
