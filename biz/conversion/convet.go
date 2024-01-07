package conversion

import (
	"speedy/read/biz/constants"
	"speedy/read/biz/conversion/api"
	"speedy/read/biz/conversion/api/generated"
	"speedy/read/biz/domain/aggregates/article_summary"
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
	return &speedy_read.ArticleSummary{
		ID:             utils.Int64ToString(do.ID),
		Article:        ArticleDOToThrift(do.Article),
		Title:          do.Title,
		Summary:        do.Summary,
		ContentSummary: do.ContentSummary,
		Outline:        do.Outline,
		Tags:           GetLabelDescription(do.LabelList),
		CreatedAt:      do.CreatedAt.Format(constants.DateTimeLayout),
	}
}

func GetLabelDescription(labelList []*article_summary.Label) []string {
	result := make([]string, 0)
	for _, label := range labelList {
		result = append(result, label.Description)
	}
	return result
}
