package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"speedy/read/biz/domain/aggregates/symbol"
	symbolInfra "speedy/read/biz/infra/repository/symbol"
	"speedy/read/biz/utils"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/site"
	"speedy/read/biz/domain/service/craw_data"
	siteInfra "speedy/read/biz/infra/repository/site"
)

type DataCrawServiceI interface {
	CrawArticle(ctx context.Context) error
}

type DataCrawService struct {
	siteRepo   site.SiteRepo
	symbolRepo symbol.SymbolRepo
}

func NewDataCrawService() DataCrawServiceI {
	return &DataCrawService{
		siteRepo:   siteInfra.NewSiteRepository(),
		symbolRepo: symbolInfra.NewSymbolRepository(),
	}
}

func (impl *DataCrawService) CrawArticle(ctx context.Context) error {
	siteList, err := impl.siteRepo.GetSiteList(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "get site list err: %v", err)
		return err
	}
	var (
		rssErr  error
		crawErr error
	)
	for _, siteDO := range siteList {
		switch siteDO.Type {
		case site.SiteTypeRss:
			//rssErr = impl.dealArticle4Rss(ctx, siteDO)
			continue
		case site.SiteTypeCraw:
			crawErr = impl.dealArticle4Craw(ctx, siteDO)
		default:
			return nil
		}
	}
	if rssErr != nil {
		return rssErr
	}
	if crawErr != nil {
		return crawErr
	}
	return nil
}

func (impl *DataCrawService) dealArticle4Rss(ctx context.Context, siteDO *site.Site) error {
	if siteDO == nil {
		return nil
	}
	switch siteDO.Tag {
	case site.SeekingAlphaTag:
		symbolList, err := impl.symbolRepo.GetList(ctx)
		if err != nil {
			klog.CtxErrorf(ctx, "Error get symbol list: %v", err)
		}
		return impl.dealSeekingAlpha(ctx, symbolList, siteDO)
	default:
		return nil
	}
}

func (impl *DataCrawService) dealSeekingAlpha(ctx context.Context, symbolList []*symbol.Symbol, siteDO *site.Site) error {
	for _, symbolDO := range symbolList {
		if symbolDO == nil {
			continue
		}
		tmpSymbol := symbolDO.Symbol
		parts := strings.Split(symbolDO.Symbol, "^")
		if len(parts) > 1 {
			tmpSymbol = parts[0]
		}
		url := "https://seekingalpha.com/api/sa/combined/XXXX.xml"
		url = strings.Replace(url, "XXXX", tmpSymbol, -1)
		resp, httpErr := http.Get(url)
		if httpErr != nil {
			klog.CtxErrorf(ctx, "Error making HTTP request to %s: %v", siteDO.Url, httpErr)
			return httpErr
		}
		if resp.StatusCode != http.StatusOK {
			klog.CtxInfof(ctx, "url %s not exist", url)
			continue
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			klog.CtxErrorf(ctx, "Error reading response body from URL %s: %v", siteDO.Url, err)
			return err
		}
		defer resp.Body.Close()

		data := craw_data.SeekingAlpha{}
		err = xml.Unmarshal(body, &data)
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
			articleUrl := item.ArticleUrl
			if strings.Contains(item.ArticleUrl, "article") {
				itemType = article.TypeArticle
				articleUrl = item.ArticleUrl
			} else if strings.Contains(item.ArticleUrl, "news") {
				itemType = article.TypeNew
				flagList := strings.Split(item.Guid, ":")
				articleUrl = "https://seekingalpha.com/news/" + flagList[2]
			}
			id, err := articleSvc.CreateArticle(ctx, &article.Article{
				Author: &article.Author{
					AuthorName: item.AuthorName,
				},
				SourceSite: &site.Site{
					Url:         siteDO.Url,
					Tag:         site.SeekingAlphaTag,
					Description: site.SeekingAlphaTag,
					Type:        siteDO.Type,
					TypeKey:     siteDO.TypeKey,
				},
				ArticleMetaList: impl.dealSeekingAlphaMeta(ctx, item.Stock),
				Status:          article.StatusInit,
				Url:             articleUrl,
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
	}
	return nil
}

func (impl *DataCrawService) dealArticle4Craw(ctx context.Context, siteDO *site.Site) error {
	klog.CtxInfof(ctx, "dealArticle4Craw start")
	if siteDO == nil {
		return nil
	}
	typeKeyList := siteDO.GetTypeKeyList(ctx)
	for _, typeKey := range typeKeyList {
		// 获取未处理的导出数据
		exportedDataList, err := craw_data.GetNotExportedData(ctx, typeKey.TaskID)
		if err != nil {
			klog.CtxErrorf(ctx, "Error get not exported data %s: %v", siteDO.TypeKey, err)
			return err
		}

		// 创建article
		articleSvc := NewArticleService()
		for _, item := range exportedDataList {
			if filterArticle(item, siteDO.Tag) {
				continue
			}
			articleDO := &article.Article{
				Author: &article.Author{
					AuthorName: item.AuthorName,
					Url:        item.AuthorUrl,
				},
				SourceSite: &site.Site{
					Url:         siteDO.Url,
					Tag:         siteDO.Tag,
					Description: siteDO.Description,
					Type:        siteDO.Type,
					TypeKey:     siteDO.TypeKey,
				},
				Status:    article.StatusInit,
				Url:       item.ArticleUrl,
				Title:     item.ArticleTitle,
				Type:      article.TypeArticle,
				PublishAt: dealCrawPublishAt(item.PublishAt, siteDO.Tag),
				Content:   dealCrawContent(item.ArticleContent, siteDO.Tag),
			}
			var articleMetaList []*article.ArticleMeta
			if len(item.Stock) > 0 {
				articleMetaList = impl.dealCrawArticleMeta(ctx, item.Stock, siteDO.Tag)
			}
			if len(articleMetaList) > 0 {
				articleDO.ArticleMetaList = articleMetaList
			}
			id, err := articleSvc.CreateArticle(ctx, articleDO)
			if err != nil {
				klog.CtxErrorf(ctx, "dealArticle4Craw create article error is %v", err)
				continue
			}
			klog.CtxInfof(ctx, "dealArticle4Craw create article success, id is %d", id)
		}

		// 标记数据为已处理
		marked := craw_data.MarkExported(ctx, siteDO.TypeKey)
		if !marked {
			klog.CtxErrorf(ctx, "dealArticle4Craw failed mark exported %s", siteDO.TypeKey)
		}
	}
	return nil
}

func (impl *DataCrawService) dealSeekingAlphaMeta(ctx context.Context, stockList []craw_data.SeekingStock) []*article.ArticleMeta {
	result := make([]*article.ArticleMeta, 0)
	if len(stockList) == 0 {
		return result
	}
	for _, stock := range stockList {
		symbolDO, err := impl.symbolRepo.GetBySymbol(ctx, stock.Symbol)
		if err != nil {
			klog.Errorf("get symbol error %v", err)
			return result
		}
		if symbolDO == nil {
			return result
		}
		result = append(result, &article.ArticleMeta{
			MetaType:  article.StockMeteType,
			MetaKey:   symbolDO.Symbol,
			MetaValue: utils.Int64ToString(symbolDO.ID),
		})
	}
	return result
}

func dealCrawPublishAt(publishAt string, tag string) time.Time {
	switch tag {
	case site.FoolTag:
		tmpTime := strings.Trim(strings.Split(publishAt, "by")[0], "\n")
		layout := "Jan 2, 2006"
		publishTime, err := time.Parse(layout, tmpTime)
		//tmpTime := strings.ReplaceAll(strings.Split(publishAt, "|")[0], " ", "")
		//layout := "Jan.02,2006"
		//publishTime, err := time.Parse(layout, tmpTime)
		if err != nil {
			klog.Errorf("time parse error %v", err)
			return time.Time{}
		}
		return publishTime
	case site.TheStreetTag:
		// 定义时间字符串和格式
		//timeStr := "Mar 2, 2024 4:04 PM EST"
		layout := "Jan 2, 2006 3:04 PM MST" // 使用Go语言的基准时间

		// 将时间字符串解析为time.Time类型
		parsedTime, err := time.Parse(layout, publishAt)
		if err != nil {
			klog.Errorf("time parse error %v", err)
			return time.Time{}
		}
		return parsedTime
	}
	return time.Time{}
}

func dealCrawContent(content string, tag string) string {
	switch tag {
	case site.FoolTag:
		reg := regexp.MustCompile(`( )+|(\n)+`)
		crawContent := reg.ReplaceAllString(content, "$1$2")
		return strings.TrimSpace(crawContent)
	case site.TheStreetTag:
		return content
	}
	return ""
}

func filterArticle(item craw_data.ExportData, tag string) bool {
	switch tag {
	case site.FoolTag:
		if len(item.ArticleUrl) == 0 {
			return true
		}
		return !strings.Contains(item.ArticleUrl, "investing")
	case site.TheStreetTag:
		if len(item.ArticleTitle) == 0 {
			return true
		}
		return false
	}
	return false
}

func (impl *DataCrawService) dealCrawArticleMeta(ctx context.Context, stock string, tag string) []*article.ArticleMeta {
	result := make([]*article.ArticleMeta, 0)
	switch tag {
	case site.FoolTag:
		list := strings.Split(strings.Replace(stock, " ", "", -1), ":")
		symbolDO, err := impl.symbolRepo.GetBySymbol(ctx, list[1])
		if err != nil {
			klog.Errorf("get symbol error %v", err)
			return result
		}
		if symbolDO == nil {
			return result
		}
		result = append(result, &article.ArticleMeta{
			MetaType:  article.StockMeteType,
			MetaKey:   symbolDO.Symbol,
			MetaValue: utils.Int64ToString(symbolDO.ID),
		})
	case site.TheStreetTag:
		symbolDO, err := impl.symbolRepo.GetBySymbol(ctx, stock)
		if err != nil {
			klog.Errorf("get symbol error %v", err)
			return result
		}
		if symbolDO == nil {
			return result
		}
		result = append(result, &article.ArticleMeta{
			MetaType:  article.StockMeteType,
			MetaKey:   symbolDO.Symbol,
			MetaValue: utils.Int64ToString(symbolDO.ID),
		})
	}
	return result
}
