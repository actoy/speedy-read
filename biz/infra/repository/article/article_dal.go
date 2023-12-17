package article

import (
	"context"
	"fmt"
)

type ArticleRepo struct {

}

func (a *ArticleRepo) Save(ctx context.Context) {
	fmt.Println("1")
}
