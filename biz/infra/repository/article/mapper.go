package article

import (
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/infra/repository/site"
)

func ConvertArticleDOToPO(articleDO *article.Article) *Article {
	if articleDO == nil {
		return nil
	}
	return &Article{
		ID:         articleDO.ID,
		Author:     ConvertAuthorDOToPO(articleDO.Author),
		SourceSite: site.CovertPO(articleDO.SourceSite),
		Language:   articleDO.Language,
		PublishAt:  articleDO.PublishAt,
		Url:        articleDO.Url,
		Type:       articleDO.Type,
		Title:      articleDO.Title,
		Content:    articleDO.Content,
		Status:     articleDO.Status,
		CreatedAt:  articleDO.CreatedAt,
		UpdatedAt:  articleDO.UpdatedAt,
	}
}

func ConvertArticlePOToDO(articlePO *Article, authorPO *Author, sitePO *site.Site) *article.Article {
	if authorPO == nil {
		return nil
	}
	return &article.Article{
		ID:         articlePO.ID,
		Author:     ConvertAuthorPOToDO(authorPO),
		SourceSite: site.CovertDO(sitePO),
		Language:   articlePO.Language,
		PublishAt:  articlePO.PublishAt,
		Url:        articlePO.Url,
		Type:       articlePO.Type,
		Title:      articlePO.Title,
		Content:    articlePO.Content,
		Status:     articlePO.Status,
		CreatedAt:  articlePO.CreatedAt,
		UpdatedAt:  articlePO.UpdatedAt,
	}
}

func ConvertAuthorDOToPO(authorDO *article.Author) *Author {
	if authorDO == nil {
		return nil
	}
	return &Author{
		ID:         authorDO.ID,
		Url:        authorDO.Url,
		AuthorName: authorDO.AuthorName,
		Image:      authorDO.Image,
		CreatedAt:  authorDO.CreatedAt,
		UpdatedAt:  authorDO.UpdatedAt,
	}
}

func ConvertAuthorPOToDO(authorPO *Author) *article.Author {
	if authorPO == nil {
		return nil
	}
	return &article.Author{
		ID:         authorPO.ID,
		Url:        authorPO.Url,
		AuthorName: authorPO.AuthorName,
		Image:      authorPO.Image,
		CreatedAt:  authorPO.CreatedAt,
		UpdatedAt:  authorPO.UpdatedAt,
	}
}
