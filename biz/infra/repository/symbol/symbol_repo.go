package symbol

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
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

func (r *Repository) GetBySymbolList(ctx context.Context, symbolList []string) ([]*symbol.Symbol, error) {
	poList, err := r.SymbolRepo.GetBySymbolList(ctx, symbolList)
	if err != nil {
		return nil, err
	}
	list := make([]*symbol.Symbol, 0)
	for _, po := range poList {
		list = append(list, CovertDO(po))
	}
	return list, nil
}

func (r *Repository) FindByID(ctx context.Context, ID string) (*symbol.Symbol, error) {
	po, err := r.SymbolRepo.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return CovertDO(po), nil
}

func (r *Repository) SearchSymbol(ctx context.Context, keyword string) ([]*symbol.Symbol, error) {
	klog.CtxInfof(ctx, "repo search symbol")
	poList, err := r.SymbolRepo.SearchSymbolByKeyWord(ctx, keyword)
	if err != nil {
		return nil, err
	}
	list := make([]*symbol.Symbol, 0)
	for _, po := range poList {
		list = append(list, CovertDO(po))
	}
	return list, nil
}

func (r *Repository) UpdateSymbol(ctx context.Context, symbolDO *symbol.Symbol) error {
	err := r.SymbolRepo.Update(ctx, CovertPO(symbolDO))
	if err != nil {
		return err
	}
	return nil
}
