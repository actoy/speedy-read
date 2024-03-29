package symbol

import (
	"context"
	"speedy/read/biz/domain/aggregates/symbol"
)

type Repository struct {
	SymbolRepo *SymbolRepo
}

func NewSymbolRepository() symbol.SymbolRepo {
	return &Repository{
		SymbolRepo: &SymbolRepo{},
	}
}

func (r *Repository) Create(ctx context.Context, symbolDO *symbol.Symbol) (int64, error) {
	id, err := r.SymbolRepo.Save(ctx, CovertPO(symbolDO))
	if err != nil {
		return int64(0), err
	}
	return id, nil
}

func (r *Repository) GetList(ctx context.Context) ([]*symbol.Symbol, error) {
	poList, err := r.SymbolRepo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*symbol.Symbol, 0)
	for _, po := range poList {
		list = append(list, CovertDO(po))
	}
	return list, nil
}

func (r *Repository) GetBySymbol(ctx context.Context, symbol string) (*symbol.Symbol, error) {
	po, err := r.SymbolRepo.GetBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}
	return CovertDO(po), nil
}
