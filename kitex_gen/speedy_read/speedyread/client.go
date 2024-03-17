// Code generated by Kitex v0.8.0. DO NOT EDIT.

package speedyread

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	speedy_read "speedy/read/kitex_gen/speedy_read"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Echo(ctx context.Context, req *speedy_read.Request, callOptions ...callopt.Option) (r *speedy_read.Response, err error)
	GetSiteInfo(ctx context.Context, req *speedy_read.GetSiteRequest, callOptions ...callopt.Option) (r *speedy_read.GetSiteResponse, err error)
	CreateSiteInfo(ctx context.Context, req *speedy_read.CreateSiteRequest, callOptions ...callopt.Option) (r *speedy_read.CreateSiteResponse, err error)
	ArticleList(ctx context.Context, req *speedy_read.GetArticleListRequest, callOptions ...callopt.Option) (r *speedy_read.GetArticleListResponse, err error)
	CreateArticle(ctx context.Context, req *speedy_read.CreateArticleRequest, callOptions ...callopt.Option) (r *speedy_read.CreateArticleResponse, err error)
	RejectArticle(ctx context.Context, req *speedy_read.RejectArticleRequest, callOptions ...callopt.Option) (r *speedy_read.RejectArticleResponse, err error)
	ArticleCount(ctx context.Context, req *speedy_read.ArticleCountRequest, callOptions ...callopt.Option) (r *speedy_read.ArticleCountResponse, err error)
	SaveArticleSummary(ctx context.Context, req *speedy_read.SaveArticleSummaryRequest, callOptions ...callopt.Option) (r *speedy_read.SaveArticleSummaryResponse, err error)
	GetArticleSummaryList(ctx context.Context, req *speedy_read.ArticleSummaryListRequest, callOptions ...callopt.Option) (r *speedy_read.ArticleSummaryListResponse, err error)
	ArticleSummaryCount(ctx context.Context, req *speedy_read.ArticleSummaryCountRequest, callOptions ...callopt.Option) (r *speedy_read.ArticleSummaryCountResponse, err error)
	ImportSymbol(ctx context.Context, req *speedy_read.Request, callOptions ...callopt.Option) (r *speedy_read.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kSpeedyReadClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSpeedyReadClient struct {
	*kClient
}

func (p *kSpeedyReadClient) Echo(ctx context.Context, req *speedy_read.Request, callOptions ...callopt.Option) (r *speedy_read.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Echo(ctx, req)
}

func (p *kSpeedyReadClient) GetSiteInfo(ctx context.Context, req *speedy_read.GetSiteRequest, callOptions ...callopt.Option) (r *speedy_read.GetSiteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetSiteInfo(ctx, req)
}

func (p *kSpeedyReadClient) CreateSiteInfo(ctx context.Context, req *speedy_read.CreateSiteRequest, callOptions ...callopt.Option) (r *speedy_read.CreateSiteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateSiteInfo(ctx, req)
}

func (p *kSpeedyReadClient) ArticleList(ctx context.Context, req *speedy_read.GetArticleListRequest, callOptions ...callopt.Option) (r *speedy_read.GetArticleListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ArticleList(ctx, req)
}

func (p *kSpeedyReadClient) CreateArticle(ctx context.Context, req *speedy_read.CreateArticleRequest, callOptions ...callopt.Option) (r *speedy_read.CreateArticleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateArticle(ctx, req)
}

func (p *kSpeedyReadClient) RejectArticle(ctx context.Context, req *speedy_read.RejectArticleRequest, callOptions ...callopt.Option) (r *speedy_read.RejectArticleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RejectArticle(ctx, req)
}

func (p *kSpeedyReadClient) ArticleCount(ctx context.Context, req *speedy_read.ArticleCountRequest, callOptions ...callopt.Option) (r *speedy_read.ArticleCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ArticleCount(ctx, req)
}

func (p *kSpeedyReadClient) SaveArticleSummary(ctx context.Context, req *speedy_read.SaveArticleSummaryRequest, callOptions ...callopt.Option) (r *speedy_read.SaveArticleSummaryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SaveArticleSummary(ctx, req)
}

func (p *kSpeedyReadClient) GetArticleSummaryList(ctx context.Context, req *speedy_read.ArticleSummaryListRequest, callOptions ...callopt.Option) (r *speedy_read.ArticleSummaryListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetArticleSummaryList(ctx, req)
}

func (p *kSpeedyReadClient) ArticleSummaryCount(ctx context.Context, req *speedy_read.ArticleSummaryCountRequest, callOptions ...callopt.Option) (r *speedy_read.ArticleSummaryCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ArticleSummaryCount(ctx, req)
}

func (p *kSpeedyReadClient) ImportSymbol(ctx context.Context, req *speedy_read.Request, callOptions ...callopt.Option) (r *speedy_read.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ImportSymbol(ctx, req)
}
