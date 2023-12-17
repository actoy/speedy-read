package article

import "context"

type Repository struct {
	articleRepo *ArticleRepo
}

func (r *Repository) Create(ctx context.Context) {
	r.articleRepo.Save(ctx)
}