package main

import (
	"context"
	"encoding/json"
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
	//testEcho(client)
	//testArticleCount(client)
	//testArticleDetail(client)
	//testImport(client)
	//testGetSymbolList(client)
	//testSearchSymbolList(client)
	testArticleSummaryCount(client)
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func testCreateSite(client speedyread.Client) {
	type SiteTypeKey struct {
		TaskID   string
		TaskType string
	}
	//key := []SiteTypeKey{
	//	{
	//		TaskID:   "338cfd8f-476a-493c-97d5-83c2092cebe7",
	//		TaskType: "all",
	//	},
	//	{
	//		TaskID:   "56c58ff4-a2e9-41e6-ae7e-3e7df36bbf69",
	//		TaskType: "all",
	//	},
	//}
	//keyJson, _ := json.Marshal(key)
	//
	//createParams := &speedy_read.CreateSiteRequest{
	//	Url:         "https://www.fool.com/investing-news/",
	//	Description: "fool.com",
	//	Tag:         "Fool",
	//	Type:        speedy_read.SiteTypeCraw,
	//	TypeKey:     string(keyJson),
	//}
	key := []SiteTypeKey{
		{
			TaskID:   "93d09025-ca44-29a8-41ad-effa2a731214",
			TaskType: "all",
		},
		{
			TaskID:   "0f5b9355-83b0-4f16-8ff6-9cd7fab818ee",
			TaskType: "all",
		},
		{
			TaskID:   "2207c8a2-80f6-4d42-a352-1fd16a6112ce",
			TaskType: "all",
		},
	}
	keyJson, _ := json.Marshal(key)

	createParams := &speedy_read.CreateSiteRequest{
		Url:         "https://www.thestreet.com/investing/stocks",
		Description: "thestreet.com",
		Tag:         "TheStreet",
		Type:        speedy_read.SiteTypeCraw,
		TypeKey:     string(keyJson),
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
		//ArticleType: stringPtr(speedy_read.TypeNew),
		SymbolIdList: []string{"1769215432683687936", "1769215428673933312"},
		SiteIdList:   []string{"1751484813790941184", "1769262327066005504", "1746385683192221696"},
		Offset:       0,
		Limit:        10,
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
		Limit:       10,
		Offset:      0,
		Symbol:      stringPtr("TSLA"),
		ArticleType: stringPtr("article"),
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
		Status:       1,
		SymbolIdList: []string{"1769215432683687936", "1769215428673933312"},
		SiteIdList:   []string{"1751484813790941184", "1769262327066005504", "1746385683192221696"},
		ArticleType:  stringPtr("news"),
	}
	resp, err := client.ArticleCount(context.Background(), req)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(resp)
}

func testArticleSummaryCount(client speedyread.Client) {
	req := &speedy_read.ArticleSummaryCountRequest{
		ArticleType: stringPtr("news"),
	}
	resp, err := client.ArticleSummaryCount(context.Background(), req)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(resp)
}

func testImport(client speedyread.Client) {
	req := &speedy_read.Request{
		Message: "message",
	}
	resp, err := client.ImportSymbol(context.Background(), req)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(resp)
}

func testGetSymbolList(client speedyread.Client) {
	list, err := client.GetSymbolList(context.Background(), &speedy_read.SymbolListRequest{})
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(list)
}

func testArticleDetail(client speedyread.Client) {
	req := &speedy_read.ArticleSummaryDetailRequest{
		SummaryID: "1769266733262049280",
	}
	resp, err := client.ArticleSummaryDetail(context.Background(), req)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(resp)
}

func testSearchSymbolList(client speedyread.Client) {
	list, err := client.SearchSymbol(context.Background(), &speedy_read.SearchSymbolRequest{
		KeyWord: "tesla",
	})
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(list)
}
