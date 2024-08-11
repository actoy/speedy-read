package symbol

import "speedy/read/biz/domain/aggregates/symbol"

func CovertPO(symbolDO *symbol.Symbol) *Symbol {
	if symbolDO == nil {
		return nil
	}
	return &Symbol{
		ID:              symbolDO.ID,
		Symbol:          symbolDO.Symbol,
		Company:         symbolDO.Company,
		CompanyZH:       symbolDO.CompanyZH,
		CompanyUrl:      symbolDO.CompanyUrl,
		Description:     symbolDO.Description,
		CompanyAddress:  symbolDO.CompanyAddress,
		CompanyBusiness: symbolDO.CompanyBusiness,
		Source:          symbolDO.Source,
		CreatedAt:       symbolDO.CreatedAt,
		UpdatedAt:       symbolDO.UpdatedAt,
	}
}

func CovertDO(symbolPO *Symbol) *symbol.Symbol {
	if symbolPO == nil {
		return nil
	}
	return &symbol.Symbol{
		ID:              symbolPO.ID,
		Symbol:          symbolPO.Symbol,
		Company:         symbolPO.Company,
		CompanyZH:       symbolPO.CompanyZH,
		CompanyUrl:      symbolPO.CompanyUrl,
		Description:     symbolPO.Description,
		CompanyAddress:  symbolPO.CompanyAddress,
		CompanyBusiness: symbolPO.CompanyBusiness,
		Source:          symbolPO.Source,
		CreatedAt:       symbolPO.CreatedAt,
		UpdatedAt:       symbolPO.UpdatedAt,
	}
}
