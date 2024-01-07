package article_summary

import (
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/article_summary"
)

func ConvertArticleSummaryPOListTODO(summaryDOList []*ArticleSummary) []*article_summary.ArticleSummary {
	result := make([]*article_summary.ArticleSummary, 0)
	for _, po := range summaryDOList {
		result = append(result, ConvertArticleSummaryPOToDO(po))
	}
	return result
}

func ConvertArticleSummaryPOToDO(summaryPO *ArticleSummary) *article_summary.ArticleSummary {
	if summaryPO == nil {
		return nil
	}
	return &article_summary.ArticleSummary{
		ID: summaryPO.ID,
		Article: &article.Article{
			ID: summaryPO.ID,
		},
		Title:          summaryPO.Title,
		Summary:        summaryPO.Summary,
		ContentSummary: summaryPO.ContentSummary,
		Outline:        summaryPO.Outline,
		CreatedAt:      summaryPO.CreatedAt,
		UpdatedAt:      summaryPO.UpdatedAt,
	}
}

func ConvertLabelPOListTODO(labelList []*Label) []*article_summary.Label {
	result := make([]*article_summary.Label, 0)
	for _, po := range labelList {
		result = append(result, ConvertLabelPOToDO(po))
	}
	return result
}

func ConvertLabelPOToDO(label *Label) *article_summary.Label {
	if label == nil {
		return nil
	}
	return &article_summary.Label{
		ID:          label.ID,
		Description: label.Description,
	}
}

func ConvertArticleSummaryDOToPO(do *article_summary.ArticleSummary) *ArticleSummary {
	if do == nil || do.Article == nil {
		return nil
	}
	return &ArticleSummary{
		ArticleID:      do.Article.ID,
		Title:          do.Title,
		Summary:        do.Summary,
		ContentSummary: do.ContentSummary,
		Outline:        do.Outline,
	}
}
