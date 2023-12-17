// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package generated

import (
	api "speedy/read/biz/conversion/api"
	article "speedy/read/biz/domain/aggregates/article"
	site "speedy/read/biz/domain/aggregates/site"
	speedyread "speedy/read/kitex_gen/speedy_read"
)

type ArticleConvertImpl struct{}

func (c *ArticleConvertImpl) ArticleDOToThrift(source *article.Article) *speedyread.Article {
	var pSpeedy_readArticle *speedyread.Article
	if source != nil {
		var speedy_readArticle speedyread.Article
		speedy_readArticle.ID = (*source).ID
		speedy_readArticle.Author = c.pArticleAuthorToPSpeedy_readAuthor((*source).Author)
		speedy_readArticle.Site = c.pSiteSiteToPSpeedy_readSiteInfo((*source).SourceSite)
		speedy_readArticle.Language = (*source).Language
		speedy_readArticle.PublishAt = api.TimeToString((*source).PublishAt)
		speedy_readArticle.Url = (*source).Url
		speedy_readArticle.Type = (*source).Type
		speedy_readArticle.Title = (*source).Title
		speedy_readArticle.Content = (*source).Content
		speedy_readArticle.Status = (*source).Status
		speedy_readArticle.Score = (*source).Score
		speedy_readArticle.CreatedAt = api.TimeToString((*source).CreatedAt)
		speedy_readArticle.UpdatedAt = api.TimeToString((*source).UpdatedAt)
		pSpeedy_readArticle = &speedy_readArticle
	}
	return pSpeedy_readArticle
}
func (c *ArticleConvertImpl) ArticleThriftToDO(source *speedyread.Article) *article.Article {
	var pArticleArticle *article.Article
	if source != nil {
		var articleArticle article.Article
		articleArticle.ID = (*source).ID
		articleArticle.Author = c.pSpeedy_readAuthorToPArticleAuthor((*source).Author)
		articleArticle.SourceSite = c.pSpeedy_readSiteInfoToPSiteSite((*source).Site)
		articleArticle.Language = (*source).Language
		articleArticle.PublishAt = api.StringToTime((*source).PublishAt)
		articleArticle.Url = (*source).Url
		articleArticle.Type = (*source).Type
		articleArticle.Title = (*source).Title
		articleArticle.Content = (*source).Content
		articleArticle.Status = (*source).Status
		articleArticle.Score = (*source).Score
		articleArticle.CreatedAt = api.StringToTime((*source).CreatedAt)
		articleArticle.UpdatedAt = api.StringToTime((*source).UpdatedAt)
		pArticleArticle = &articleArticle
	}
	return pArticleArticle
}
func (c *ArticleConvertImpl) pArticleAuthorToPSpeedy_readAuthor(source *article.Author) *speedyread.Author {
	var pSpeedy_readAuthor *speedyread.Author
	if source != nil {
		var speedy_readAuthor speedyread.Author
		speedy_readAuthor.ID = (*source).ID
		speedy_readAuthor.Url = (*source).Url
		speedy_readAuthor.AuthorName = (*source).AuthorName
		speedy_readAuthor.Image = (*source).Image
		speedy_readAuthor.CreatedAt = api.TimeToString((*source).CreatedAt)
		speedy_readAuthor.UpdatedAt = api.TimeToString((*source).UpdatedAt)
		pSpeedy_readAuthor = &speedy_readAuthor
	}
	return pSpeedy_readAuthor
}
func (c *ArticleConvertImpl) pSiteSiteToPSpeedy_readSiteInfo(source *site.Site) *speedyread.SiteInfo {
	var pSpeedy_readSiteInfo *speedyread.SiteInfo
	if source != nil {
		var speedy_readSiteInfo speedyread.SiteInfo
		speedy_readSiteInfo.ID = (*source).ID
		speedy_readSiteInfo.SourceID = (*source).SourceID
		speedy_readSiteInfo.SourceType = (*source).SourceType
		speedy_readSiteInfo.Url = (*source).Url
		speedy_readSiteInfo.Description = (*source).Description
		speedy_readSiteInfo.CreatedAt = api.TimeToString((*source).CreatedAt)
		speedy_readSiteInfo.UpdatedAt = api.TimeToString((*source).UpdatedAt)
		pSpeedy_readSiteInfo = &speedy_readSiteInfo
	}
	return pSpeedy_readSiteInfo
}
func (c *ArticleConvertImpl) pSpeedy_readAuthorToPArticleAuthor(source *speedyread.Author) *article.Author {
	var pArticleAuthor *article.Author
	if source != nil {
		var articleAuthor article.Author
		articleAuthor.ID = (*source).ID
		articleAuthor.Url = (*source).Url
		articleAuthor.AuthorName = (*source).AuthorName
		articleAuthor.Image = (*source).Image
		articleAuthor.CreatedAt = api.StringToTime((*source).CreatedAt)
		articleAuthor.UpdatedAt = api.StringToTime((*source).UpdatedAt)
		pArticleAuthor = &articleAuthor
	}
	return pArticleAuthor
}
func (c *ArticleConvertImpl) pSpeedy_readSiteInfoToPSiteSite(source *speedyread.SiteInfo) *site.Site {
	var pSiteSite *site.Site
	if source != nil {
		var siteSite site.Site
		siteSite.ID = (*source).ID
		siteSite.SourceID = (*source).SourceID
		siteSite.SourceType = (*source).SourceType
		siteSite.Url = (*source).Url
		siteSite.Description = (*source).Description
		siteSite.CreatedAt = api.StringToTime((*source).CreatedAt)
		siteSite.UpdatedAt = api.StringToTime((*source).UpdatedAt)
		pSiteSite = &siteSite
	}
	return pSiteSite
}
