package article_summary

import (
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/article_summary"
)

func convertArticleSummaryDO(summaryDOList []*ArticleSummary, summaryContentList []*SummaryContent,
	outlineList []*SummaryOutline) []*article_summary.ArticleSummary {
	result := make([]*article_summary.ArticleSummary, 0)
	for _, po := range summaryDOList {
		result = append(result, convertArticleSummaryPOToDO(po, summaryContentList, outlineList))
	}
	return result
}

func convertArticleSummaryPOToDO(summaryPO *ArticleSummary, summaryContentList []*SummaryContent,
	outlineList []*SummaryOutline) *article_summary.ArticleSummary {
	if summaryPO == nil {
		return nil
	}
	return &article_summary.ArticleSummary{
		ID: summaryPO.ID,
		Article: &article.Article{
			ID: summaryPO.ArticleID,
		},
		Title:           summaryPO.Title,
		Summary:         summaryPO.Summary,
		ContentSummary:  covertSummaryContentToDO(summaryContentList, summaryPO.ID),
		Outline:         convertSummaryOutlineDO(outlineList, summaryPO.ID),
		TradingProposal: summaryPO.TradingProposal,
		CreatedAt:       summaryPO.CreatedAt,
		UpdatedAt:       summaryPO.UpdatedAt,
	}
}

func covertSummaryContentToDO(summaryContentList []*SummaryContent, summaryID int64) *article_summary.SummaryContent {
	for _, summaryContent := range summaryContentList {
		if summaryContent.SummaryID == summaryID {
			return &article_summary.SummaryContent{
				ID:          summaryContent.ID,
				SummaryID:   summaryContent.SummaryID,
				Original:    summaryContent.Original,
				Translation: summaryContent.Translation,
			}
		}
	}
	return nil
}

func convertSummaryOutlineDO(outlineList []*SummaryOutline, summaryID int64) []*article_summary.SummaryOutline {
	result := make([]*article_summary.SummaryOutline, 0)
	for _, outline := range outlineList {
		if outline.SummaryID == summaryID {
			result = append(result, &article_summary.SummaryOutline{
				ID:        outline.ID,
				SummaryID: outline.SummaryID,
				Title:     outline.Title,
				Content:   outline.Content,
			})
		}
	}
	return result
}

func convertLabelPOListTODO(labelList []*Label) []*article_summary.Label {
	result := make([]*article_summary.Label, 0)
	for _, po := range labelList {
		result = append(result, convertLabelPOToDO(po))
	}
	return result
}

func convertLabelPOToDO(label *Label) *article_summary.Label {
	if label == nil {
		return nil
	}
	return &article_summary.Label{
		ID:          label.ID,
		Description: label.Description,
	}
}

func convertArticleSummaryDOToPO(do *article_summary.ArticleSummary) *ArticleSummary {
	if do == nil || do.Article == nil {
		return nil
	}
	return &ArticleSummary{
		ArticleID:       do.Article.ID,
		Title:           do.Title,
		Summary:         do.Summary,
		TradingProposal: do.TradingProposal,
	}
}

func convertContentSummaryDoToPo(do *article_summary.SummaryContent, summaryID int64) *SummaryContent {
	return &SummaryContent{
		SummaryID:   summaryID,
		Original:    do.Original,
		Translation: do.Translation,
	}
}

func convertSummaryOutlineDoToPo(doList []*article_summary.SummaryOutline, summaryID int64) []*SummaryOutline {
	result := make([]*SummaryOutline, 0)
	for _, do := range doList {
		result = append(result, &SummaryOutline{
			SummaryID: summaryID,
			Title:     do.Title,
			Content:   do.Content,
		})
	}
	return result
}
