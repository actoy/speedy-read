package symbol

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"speedy/read/biz/infra"
	"time"
)

type SymbolRepo struct {
}

func (dal *SymbolRepo) Save(ctx context.Context, symbolPO *Symbol) (int64, error) {
	exit, err := dal.GetBySymbol(ctx, symbolPO.Symbol)
	if err == nil && exit != nil {
		exit.Company = symbolPO.Company
		exit.Source = symbolPO.Source
		exit.UpdatedAt = time.Now()
		infra.DB.WithContext(ctx).Save(exit)
		return exit.ID, nil
	}
	symbolPO.ID = infra.IdGenerate()
	symbolPO.CreatedAt = time.Now()
	symbolPO.UpdatedAt = time.Now()
	result := infra.DB.WithContext(ctx).Create(symbolPO)
	if result.Error != nil {
		return int64(0), result.Error
	}
	return symbolPO.ID, nil
}

func (dal *SymbolRepo) FindByID(ctx context.Context, ID string) (*Symbol, error) {
	symbolPO := &Symbol{}
	result := infra.DB.WithContext(ctx).First(&symbolPO, ID)
	if result.Error == nil {
		return symbolPO, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *SymbolRepo) GetBySymbol(ctx context.Context, symbol string) (*Symbol, error) {
	symbolPO := &Symbol{}
	result := infra.DB.WithContext(ctx).Where("symbol = ?", symbol).First(&symbolPO)
	if result.Error == nil {
		return symbolPO, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *SymbolRepo) GetBySymbolList(ctx context.Context, symbolList []string) ([]*Symbol, error) {
	list := make([]*Symbol, 0)
	result := infra.DB.WithContext(ctx).Where("symbol in ?", symbolList).Find(&list)
	if result.Error == nil {
		return list, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *SymbolRepo) GetList(ctx context.Context) ([]*Symbol, error) {
	list := make([]*Symbol, 0)
	result := infra.DB.WithContext(ctx).Find(&list)
	if result.Error == nil {
		return list, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *SymbolRepo) SearchSymbolByKeyWord(ctx context.Context, keyword string) ([]*Symbol, error) {
	list := make([]*Symbol, 0)
	result := infra.DB.WithContext(ctx).
		Where("symbol like ?", "%"+keyword+"%").
		Or("company like ?", "%"+keyword+"%").
		Find(&list)
	if result.Error == nil {
		return list, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, result.Error
}

func (dal *SymbolRepo) Update(ctx context.Context, symbolPO *Symbol) error {
	result := infra.DB.WithContext(ctx).Save(symbolPO)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
