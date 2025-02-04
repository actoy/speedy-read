package app

import (
	"context"
	"encoding/csv"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"

	"speedy/read/biz/domain/aggregates/symbol"
	symbolInfra "speedy/read/biz/infra/repository/symbol"
)

type SymbolApplicationI interface {
	Import(ctx context.Context) error
	GetSymbolList(ctx context.Context) ([]*symbol.Symbol, error)
	SearchSymbolByKeyword(ctx context.Context, keyword string) ([]*symbol.Symbol, error)
	UpdateSymbol(ctx context.Context, params UpdateSymbolParams) error
	GetSymbolByID(ctx context.Context, id string) (*symbol.Symbol, error)
	GetBySymbol(ctx context.Context, symbolTag string) (*symbol.Symbol, error)
	// GetSymbolMapInfo
	// response key is Symbol
	GetSymbolMapInfo(ctx context.Context, symbols []string) (map[string]*symbol.Symbol, error)
}

type SymbolApplication struct {
	symbolRepo symbol.SymbolRepo
}

func NewSymbolApplication() SymbolApplicationI {
	return &SymbolApplication{
		symbolRepo: symbolInfra.NewSymbolRepository(),
	}
}

func (impl *SymbolApplication) ImportCSVData(ctx context.Context, source, path string) {
	file, err := os.Open(path)
	if err != nil {
		klog.CtxErrorf(ctx, "open csv file error is %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// 读取CSV文件的所有记录
	records, err := reader.ReadAll()
	if err != nil {
		klog.CtxErrorf(ctx, "read csv data error is %v", err)
		return
	}

	// 读取特定列的数据
	columnIndex := 1 // 指定要读取的列索引，假设这里是第二列（从0开始计数）
	isFirst := true
	for _, record := range records {
		if isFirst {
			isFirst = false
			continue
		}
		// 检查记录是否有足够的列数
		if len(record) > columnIndex {
			// 输出指定列的值
			do := &symbol.Symbol{
				Symbol:  record[0],
				Company: record[1],
				Source:  source,
			}
			_, err := impl.symbolRepo.Create(ctx, do)
			if err != nil {

			}
		}
	}
}

func (impl *SymbolApplication) Import(ctx context.Context) error {
	pathList := map[string]string{
		"AMEX":   "AMEX.csv",
		"NASDAQ": "NASDAQ.csv",
		"NYSE":   "NYSE.csv",
	}
	currentDir, err := os.Getwd()
	if err != nil {
		klog.CtxErrorf(ctx, "get wd error is %v", err)
		return err
	}
	for source, fileName := range pathList {
		filePath := currentDir + "/biz/statics/" + fileName
		impl.ImportCSVData(ctx, source, filePath)
	}
	return nil
}

func (impl *SymbolApplication) GetSymbolList(ctx context.Context) ([]*symbol.Symbol, error) {
	return impl.symbolRepo.GetList(ctx)
}

func (impl *SymbolApplication) SearchSymbolByKeyword(ctx context.Context, keyword string) ([]*symbol.Symbol, error) {
	return impl.symbolRepo.SearchSymbol(ctx, keyword)
}

type UpdateSymbolParams struct {
	ID              string
	Company         *string
	CompanyZH       *string
	CompanyUrl      *string
	CompanyAddress  *string
	Description     *string
	CompanyBusiness *string
}

func (impl *SymbolApplication) UpdateSymbol(ctx context.Context, params UpdateSymbolParams) error {
	sym, err := impl.symbolRepo.FindByID(ctx, params.ID)
	if err != nil {
		klog.CtxErrorf(ctx, "find by id error is %v", err)
		return err
	}
	if params.Company != nil {
		sym.Company = *params.Company
	}
	if params.CompanyZH != nil {
		sym.CompanyZH = *params.CompanyZH
	}
	if params.CompanyUrl != nil {
		sym.CompanyUrl = *params.CompanyUrl
	}
	if params.CompanyAddress != nil {
		sym.CompanyAddress = *params.CompanyAddress
	}
	if params.Description != nil {
		sym.Description = *params.Description
	}
	if params.CompanyBusiness != nil {
		sym.CompanyBusiness = *params.CompanyBusiness
	}
	return impl.symbolRepo.UpdateSymbol(ctx, sym)
}

func (impl *SymbolApplication) GetSymbolByID(ctx context.Context, id string) (*symbol.Symbol, error) {
	return impl.symbolRepo.FindByID(ctx, id)
}

func (impl *SymbolApplication) GetBySymbol(ctx context.Context, symbolTag string) (*symbol.Symbol, error) {
	return impl.symbolRepo.GetBySymbol(ctx, symbolTag)
}

func (impl *SymbolApplication) GetSymbolMapInfo(ctx context.Context, symbols []string) (map[string]*symbol.Symbol, error) {
	symbolList, err := impl.symbolRepo.GetBySymbolList(ctx, symbols)
	if err != nil {
		klog.CtxErrorf(ctx, "get symbol list error is %v", err)
		return nil, err
	}
	symbolMap := make(map[string]*symbol.Symbol)
	for _, sym := range symbolList {
		symbolMap[sym.Symbol] = sym
	}
	return symbolMap, nil
}
