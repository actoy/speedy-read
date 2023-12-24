package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/app"
	"speedy/read/biz/conversion"
	"speedy/read/kitex_gen/speedy_read"
)

type ArticleHandlerI interface {
	GetArticleList(ctx context.Context, req *speedy_read.GetArticleListRequest) (resp *speedy_read.GetArticleListResponse, err error)
	CreateArticle(ctx context.Context, req *speedy_read.CreateArticleRequest) (resp *speedy_read.CreateArticleResponse, err error)
	RejectArticle(ctx context.Context, req *speedy_read.RejectArticleRequest) (resp *speedy_read.RejectArticleResponse, err error)
}

type ArticleHandler struct {
	articleSvc app.ArticleApplicationI
}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{
		articleSvc: app.NewArticleApplication(),
	}
}

// Echo implements the SpeedyReadImpl interface.
func (s *ArticleHandler) Echo(ctx context.Context, req *speedy_read.Request) (resp *speedy_read.Response, err error) {
	return &speedy_read.Response{Message: req.Message}, nil
}

func (s *ArticleHandler) GetArticleList(ctx context.Context, req *speedy_read.GetArticleListRequest) (resp *speedy_read.GetArticleListResponse, err error) {
	articleInfoList, err := s.articleSvc.GetArticleList(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		klog.CtxErrorf(ctx, "get article list error %v", err)
		return nil, err
	}
	articleList := make([]*speedy_read.Article, 0)
	for _, articleInfo := range articleInfoList {
		articleList = append(articleList, conversion.ArticleDOToThrift(articleInfo))
	}
	return &speedy_read.GetArticleListResponse{
		ArticleList: articleList,
	}, nil
}

func (s *ArticleHandler) CreateArticle(ctx context.Context, req *speedy_read.CreateArticleRequest) (resp *speedy_read.CreateArticleResponse, err error) {
	id, err := s.articleSvc.CreateArticle(ctx, conversion.ArticleThriftToDO(req.GetArticle()))
	if err != nil {
		klog.CtxErrorf(ctx, "create article error %v", err)
		return nil, err
	}
	return &speedy_read.CreateArticleResponse{
		ID: id,
	}, nil
}

func (s *ArticleHandler) RejectArticle(ctx context.Context, req *speedy_read.RejectArticleRequest) (resp *speedy_read.RejectArticleResponse, err error) {
	err = s.articleSvc.RejectArticle(ctx, req.GetArticleID())
	if err != nil {
		klog.CtxErrorf(ctx, "create article error %v", err)
		return &speedy_read.RejectArticleResponse{
			Success: false,
		}, err
	}
	return &speedy_read.RejectArticleResponse{
		Success: true,
	}, nil
}
