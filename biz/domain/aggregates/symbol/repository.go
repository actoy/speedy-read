package symbol

import (
	"context"
)

type SymbolRepo interface {
	Create(ctx context.Context, SymbolDO *Symbol) (int64, error)
	GetList(ctx context.Context) ([]*Symbol, error)
	GetBySymbol(ctx context.Context, symbol string) (*Symbol, error)
	SearchSymbol(ctx context.Context, keyword string) ([]*Symbol, error)
	FindByID(ctx context.Context, ID string) (*Symbol, error)
	UpdateSymbol(ctx context.Context, symbolDO *Symbol) error
}
