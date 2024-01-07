package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"
	"net/http"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/site"
	"speedy/read/biz/domain/service/craw_data"
	siteInfra "speedy/read/biz/infra/repository/site"
	"strings"
	"time"
)

type DataCrawServiceI interface {
	CrawArticle(ctx context.Context) error
}

type DataCrawService struct {
	siteRepo site.SiteRepo
}

func NewDataCrawService() DataCrawServiceI {
	return &DataCrawService{
		siteRepo: siteInfra.NewSiteRepository(),
	}
}

func (impl *DataCrawService) CrawArticle(ctx context.Context) error {
	siteList, err := impl.siteRepo.GetSiteList(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "get site list err: %v", err)
		return err
	}
	var (
		resp = &http.Response{}
	)
	for _, siteDO := range siteList {
		if siteDO == nil {
			continue
		}
		resp, err = http.Get(siteDO.Url)
		if err != nil {
			klog.CtxErrorf(ctx, "Error making HTTP request to %s: %v", siteDO.Url, err)
			continue
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			klog.CtxErrorf(ctx, "Error reading response body from URL %s: %v", siteDO.Url, err)
			continue
		}
		switch siteDO.Tag {
		case site.SeekingAlphaTag:
			return impl.dealSeekingAlpha(ctx, body, siteDO)
		default:
			return nil
		}
	}
	defer resp.Body.Close()
	return nil
}

func (impl *DataCrawService) dealSeekingAlpha(ctx context.Context, body []byte, siteDO *site.Site) error {
	data := craw_data.SeekingAlpha{}
	err := xml.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	articleSvc := NewArticleService()
	for _, item := range data.Channel.Item {
		publishAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			continue
		}
		itemType := ""
		if strings.Contains(item.ArticleUrl, "article") {
			itemType = article.TypeArticle
		} else if strings.Contains(item.ArticleUrl, "news") {
			itemType = article.TypeNew
		}
		id, err := articleSvc.CreateArticle(ctx, &article.Article{
			Author: &article.Author{
				AuthorName: item.AuthorName,
			},
			SourceSite: &site.Site{
				Url:         siteDO.Url,
				Tag:         site.SeekingAlphaTag,
				Description: site.SeekingAlphaTag,
			},
			ArticleMetaList: dearSeekingAlphaMeta(item.Stock),
			Status:          article.StatusInit,
			Url:             item.ArticleUrl,
			Title:           item.Title,
			PublishAt:       publishAt,
			Type:            itemType,
		})
		if err != nil {
			klog.CtxErrorf(ctx, "create article error is %v", err)
			continue
		}
		klog.CtxInfof(ctx, "create article success, id is %d", id)
	}
	return nil
}

func dearSeekingAlphaMeta(stockList []craw_data.SeekingStock) []*article.ArticleMeta {
	result := make([]*article.ArticleMeta, 0)
	if len(stockList) == 0 {
		result = append(result, &article.ArticleMeta{
			MetaType: article.StockMeteType,
			//MetaKey:   siteDO.SourceID,
			//MetaValue: siteDO.SourceID,
		})
	}
	for _, stock := range stockList {
		result = append(result, &article.ArticleMeta{
			MetaType:  article.StockMeteType,
			MetaKey:   stock.Symbol,
			MetaValue: stock.Company,
		})
	}
	return result
}
