package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/app"
	"speedy/read/biz/conversion"
	"speedy/read/biz/utils"
	"speedy/read/kitex_gen/speedy_read"
)

type ArticleHandlerI interface {
	GetArticleList(ctx context.Context, req *speedy_read.GetArticleListRequest) (resp *speedy_read.GetArticleListResponse, err error)
	CreateArticle(ctx context.Context, req *speedy_read.CreateArticleRequest) (resp *speedy_read.CreateArticleResponse, err error)
	RejectArticle(ctx context.Context, req *speedy_read.RejectArticleRequest) (resp *speedy_read.RejectArticleResponse, err error)
	ArticleCount(ctx context.Context, req *speedy_read.ArticleCountRequest) (resp *speedy_read.ArticleCountResponse, err error)
}

type ArticleHandler struct {
	articleSvc app.ArticleApplicationI
}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{
		articleSvc: app.NewArticleApplication(),
	}
}

func (s *ArticleHandler) GetArticleList(ctx context.Context, req *speedy_read.GetArticleListRequest) (resp *speedy_read.GetArticleListResponse, err error) {
	articleInfoList, err := s.articleSvc.GetArticleList(ctx, app.ArticleListParams{
		SiteIdList:  utils.StringToInt64List(req.GetSiteIdList()),
		ArticleType: req.GetArticleType(),
		Limit:       req.GetLimit(),
		OffSet:      req.GetOffset(),
	})
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
		ID: utils.Int64ToString(id),
	}, nil
}

func (s *ArticleHandler) RejectArticle(ctx context.Context, req *speedy_read.RejectArticleRequest) (resp *speedy_read.RejectArticleResponse, err error) {
	err = s.articleSvc.RejectArticle(ctx, utils.StringToInt64(req.GetArticleID()))
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

func (s *ArticleHandler) ArticleCount(ctx context.Context, req *speedy_read.ArticleCountRequest) (resp *speedy_read.ArticleCountResponse, err error) {
	count, err := s.articleSvc.ArticleCount(ctx, req.GetStatus())
	if err != nil {
		klog.CtxErrorf(ctx, "get article count error %v", err)
		return &speedy_read.ArticleCountResponse{
			Count: 0,
		}, err
	}
	return &speedy_read.ArticleCountResponse{
		Count: count,
	}, nil
}
