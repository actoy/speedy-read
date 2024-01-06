package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"
	"log"
	"net/http"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/site"
	"speedy/read/biz/domain/service/craw_data"
	"time"
)

type DataCrawServiceI interface {
	CrawArticle(ctx context.Context) error
}

type DataCrawService struct {
}

func NewDataCrawService() DataCrawServiceI {
	return &DataCrawService{}
}

func (impl *DataCrawService) CrawArticle(ctx context.Context) error {
	url := "https://seekingalpha.com/api/sa/combined/TSLA.xml"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making HTTP request to %s: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body from URL %s: %v", url, err)
	}

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
		id, err := articleSvc.CreateArticle(ctx, &article.Article{
			Author: &article.Author{
				AuthorName: item.AuthorName,
			},
			SourceSite: &site.Site{
				Url: url,
			},
			Status:    article.StatusInit,
			Url:       item.ArticleUrl,
			Title:     item.Title,
			PublishAt: publishAt,
		})
		if err != nil {
			klog.CtxErrorf(ctx, "create article error is %v", err)
			continue
		}
		klog.CtxInfof(ctx, "create article success, id is %d", id)
	}

	return nil
}
