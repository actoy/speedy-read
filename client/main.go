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
	//testGetArticleList(client)
	//testSaveArticleSummary(client)
	//testArticleSummaryList(client)
	testEcho(client)
	//testArticleCount(client)
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func testCreateSite(client speedyread.Client) {
	createParams := &speedy_read.CreateSiteRequest{
		Url:         "https://seekingalpha.com/api/sa/combined/TSLA.xml",
		Description: "seeking alpha Tesla",
		Tag:         "SeekingAlpha",
		MetaType:    stringPtr("stock"),
		MetaValue:   stringPtr("TSLA"),
		MetaKey:     stringPtr("Tesla, Inc."),
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
	articleList, err := client.ArticleList(context.Background(), &speedy_read.GetArticleListRequest{
		ArticleType: stringPtr(speedy_read.TypeNew),
		SiteIdList:  []string{"3"},
		Offset:      0,
		Limit:       10,
	})
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(articleList)
}

func testSaveArticleSummary(client speedyread.Client) {
	createParams := &speedy_read.SaveArticleSummaryRequest{
		ArticleID: "30",
		Title:     "this is Title",
		Content:   "this is content",
		Summary:   "this is Summary",
		ContentSummary: &speedy_read.ArticleContentSummary{
			Original:    "original",
			Translation: "translation",
		},
		Outline: []*speedy_read.SummaryOutline{
			{
				Title:   "title1",
				Content: "content1",
			},
			{
				Title:   "title2",
				Content: "content2",
			},
		},
		Tags:            []string{"tag1", "tag2"},
		TradingProposal: 1,
	}
	id, err := client.SaveArticleSummary(context.Background(), createParams)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(id)
}

func testArticleSummaryList(client speedyread.Client) {
	SummaryList, err := client.GetArticleSummaryList(context.Background(), &speedy_read.ArticleSummaryListRequest{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(SummaryList)
}

func testEcho(client speedyread.Client) {
	req := &speedy_read.Request{
		Message: "message",
	}
	resp, err := client.Echo(context.Background(), req)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(resp)
}

func testArticleCount(client speedyread.Client) {
	req := &speedy_read.ArticleCountRequest{
		Status: 1,
	}
	resp, err := client.ArticleCount(context.Background(), req)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(resp)
}
