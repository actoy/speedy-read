package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/kitex_gen/speedy_read"
	"speedy/read/kitex_gen/speedy_read/speedyread"
)

func main() {
	client, err := speedyread.NewClient("speedy_read", client.WithHostPorts("127.0.0.1:3000"))
	if err != nil {
		klog.Error(err)
	}
	//testCreateSite(client)
	//testGetSiteList(client)
	//testCreateArticle(client)
	testGetArticleList(client)
}

func testCreateSite(client speedyread.Client) {
	createParams := &speedy_read.CreateSiteRequest{
		SourceID:    int64(1),
		SourceType:  "sourceType",
		Url:         "url",
		Description: "desc",
	}
	id, err := client.CreateSiteInfo(context.Background(), createParams)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(id)
}

func testGetSiteList(client speedyread.Client) {
	siteList, err := client.GetSiteInfo(context.Background(), &speedy_read.GetSiteRequest{})
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(siteList)
}

func testCreateArticle(client speedyread.Client) {
	createParams := &speedy_read.CreateArticleRequest{
		Article: &speedy_read.Article{
			Author: &speedy_read.Author{
				Url:        "Url",
				AuthorName: "authorName",
				Image:      "image",
			},
			Site: &speedy_read.SiteInfo{
				SourceID:    int64(1),
				SourceType:  "sourceType",
				Url:         "url",
				Description: "desc",
			},
			Language:  "language",
			PublishAt: "2023-12-17 10:00:00",
			Url:       "article_Url",
			Type:      "article Type",
			Title:     "article Title",
			Content:   "article Content",
			Status:    int32(1),
		},
	}
	id, err := client.CreateArticle(context.Background(), createParams)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(id)
}

func testGetArticleList(client speedyread.Client) {
	articleList, err := client.ArticleList(context.Background(), &speedy_read.GetArticleListRequest{})
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(articleList)
}
