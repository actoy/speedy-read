package article

import (
	"context"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
	"time"
)

type AuthorRepo struct {
}

func (dal *AuthorRepo) GetAuthorByID(ctx context.Context, id int64) (*Author, error) {
	authorPO := &Author{}
	result := infra.DB.First(&authorPO, id)
	if result.Error == nil {
		return authorPO, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *AuthorRepo) GetAuthorByAuthorName(ctx context.Context, AuthorName string) (*Author, error) {
	authorPO := &Author{}
	result := infra.DB.Where("author_name = ?", AuthorName).First(&authorPO)
	if result.Error == nil {
		return authorPO, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *AuthorRepo) Save(ctx context.Context, authorPO *Author) (int64, error) {
	authorPO.CreatedAt = time.Now()
	authorPO.UpdatedAt = time.Now()
	result := infra.DB.Create(authorPO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return authorPO.ID, nil
}
