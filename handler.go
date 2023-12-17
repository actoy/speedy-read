package main

import (
	"context"
	"speedy/read/biz/api/rpc"
	speedy_read "speedy/read/kitex_gen/speedy_read"
)

// SpeedyReadImpl implements the last service interface defined in the IDL.
type SpeedyReadImpl struct{}

// Echo implements the SpeedyReadImpl interface.
func (s *SpeedyReadImpl) Echo(ctx context.Context, req *speedy_read.Request) (resp *speedy_read.Response, err error) {
	return rpc.NewArticleHandler().Echo(ctx, req)
}

// GetSiteInfo implements the SpeedyReadImpl interface.
func (s *SpeedyReadImpl) GetSiteInfo(ctx context.Context, req *speedy_read.GetSiteRequest) (resp *speedy_read.GetSiteResponse, err error) {
	return rpc.NewSiteHandler().GetSiteInfo(ctx, req)
}

// CreateSiteInfo implements the SpeedyReadImpl interface.
func (s *SpeedyReadImpl) CreateSiteInfo(ctx context.Context, req *speedy_read.CreateSiteRequest) (resp *speedy_read.CreateSiteResponse, err error) {
	return rpc.NewSiteHandler().CreateSiteInfo(ctx, req)
}

// ArticleList implements the SpeedyReadImpl interface.
func (s *SpeedyReadImpl) ArticleList(ctx context.Context, req *speedy_read.GetArticleListRequest) (resp *speedy_read.GetArticleListResponse, err error) {
	return rpc.NewArticleHandler().GetArticleList(ctx, req)
}

// CreateArticle implements the SpeedyReadImpl interface.
func (s *SpeedyReadImpl) CreateArticle(ctx context.Context, req *speedy_read.CreateArticleRequest) (resp *speedy_read.CreateArticleResponse, err error) {
	return rpc.NewArticleHandler().CreateArticle(ctx, req)
}
