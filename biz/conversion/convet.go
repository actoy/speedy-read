package conversion

import (
	"encoding/json"
	"speedy/read/biz/constants"
	"speedy/read/biz/conversion/api"
	"speedy/read/biz/conversion/api/generated"
	"speedy/read/biz/domain/aggregates/article_summary"
	"speedy/read/biz/domain/aggregates/symbol"
	"speedy/read/biz/utils"
	"speedy/read/kitex_gen/speedy_read"
)

var (
	articleConvert api.ArticleConvert = &generated.ArticleConvertImpl{}
)

var (
	ArticleDOToThrift = articleConvert.ArticleDOToThrift
	ArticleThriftToDO = articleConvert.ArticleThriftToDO
)

func ArticleSummaryDOToThrift(do *article_summary.ArticleSummary) *speedy_read.ArticleSummary {
	if do == nil {
		return &speedy_read.ArticleSummary{}
	}
	return &speedy_read.ArticleSummary{
		ID:              utils.Int64ToString(do.ID),
		Article:         ArticleDOToThrift(do.Article),
		Title:           do.Title,
		Summary:         do.Summary,
		ContentSummary:  ConvertContentSummaryToThrift(do.ContentSummary),
		Outline:         CovertSummaryOutlineListToThrift(do.Outline),
		Tags:            GetLabelDescription(do.LabelList),
		CreatedAt:       do.CreatedAt.Format(constants.DateTimeLayout),
		TradingProposal: do.TradingProposal,
	}
}

func GetLabelDescription(labelList []*article_summary.Label) []string {
	result := make([]string, 0)
	for _, label := range labelList {
		result = append(result, label.Description)
	}
	return result
}

func CovertLabelToDO(tagList []string) []*article_summary.Label {
	labels := make([]*article_summary.Label, 0)
	for _, desc := range tagList {
		labels = append(labels, &article_summary.Label{
			Description: desc,
		})
	}
	return labels
}

func CovertSummaryContentToDO(contentSummary *speedy_read.ArticleContentSummary) *article_summary.SummaryContent {
	return &article_summary.SummaryContent{
		Original:    contentSummary.Original,
		Translation: contentSummary.Translation,
	}
}

func ConvertContentSummaryToThrift(contentSummary *article_summary.SummaryContent) *speedy_read.ArticleContentSummary {
	if contentSummary == nil {
		return &speedy_read.ArticleContentSummary{}
	}
	return &speedy_read.ArticleContentSummary{
		Original:    contentSummary.Original,
		Translation: contentSummary.Translation,
	}
}

func CovertSummaryOutlineListToDO(outlineList []*speedy_read.SummaryOutline) []*article_summary.SummaryOutline {
	outlineDoList := make([]*article_summary.SummaryOutline, 0)
	for _, outline := range outlineList {
		outlineDoList = append(outlineDoList, &article_summary.SummaryOutline{
			Title:   outline.Title,
			Content: outline.Content,
		})
	}
	return outlineDoList
}

func CovertSummaryOutlineListToThrift(outlineList []*article_summary.SummaryOutline) []*speedy_read.SummaryOutline {
	outlineDoList := make([]*speedy_read.SummaryOutline, 0)
	for _, outline := range outlineList {
		outlineDoList = append(outlineDoList, &speedy_read.SummaryOutline{
			Title:   outline.Title,
			Content: outline.Content,
		})
	}
	return outlineDoList
}

func ConvertSummaryOutlineStringToDO(outlineStrings []string) []*article_summary.SummaryOutline {
	outlineDoList := make([]*article_summary.SummaryOutline, 0)
	for _, outline := range outlineStrings {
		var outlineDO article_summary.SummaryOutline
		_ = json.Unmarshal([]byte(outline), &outlineDO)
		outlineDoList = append(outlineDoList, &outlineDO)
	}
	return outlineDoList
}

func SymbolDOToThrift(symbolDO *symbol.Symbol) *speedy_read.Symbol {
	return &speedy_read.Symbol{
		ID:              utils.Int64ToString(symbolDO.ID),
		Symbol:          symbolDO.Symbol,
		Company:         symbolDO.Company,
		Source:          symbolDO.Source,
		CompanyZH:       symbolDO.CompanyZH,
		CompanyUrl:      symbolDO.CompanyUrl,
		CompanyAddress:  symbolDO.CompanyAddress,
		Description:     symbolDO.Description,
		CompanyBusiness: symbolDO.CompanyBusiness,
	}
}
