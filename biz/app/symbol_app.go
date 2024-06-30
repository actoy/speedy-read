package app

import (
	"context"
	"encoding/csv"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"

	"speedy/read/biz/domain/aggregates/symbol"
	symbolInfra "speedy/read/biz/infra/repository/symbol"
)

type SymbolApplicationI interface {
	Import(ctx context.Context) error
	GetSymbolList(ctx context.Context) ([]*symbol.Symbol, error)
	SearchSymbolByKeyword(ctx context.Context, keyword string) ([]*symbol.Symbol, error)
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
