package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/app"
	"speedy/read/biz/conversion"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/article_summary"
	"speedy/read/kitex_gen/speedy_read"
)

type ArticleSummaryHandlerI interface {
	Save(ctx context.Context, req *speedy_read.SaveArticleSummaryRequest) (resp *speedy_read.SaveArticleSummaryResponse, err error)
	ArticleSummaryList(ctx context.Context, req *speedy_read.ArticleSummaryListRequest) (resp *speedy_read.ArticleSummaryListResponse, err error)
}

type ArticleSummaryHandler struct {
	articleSummarySvc app.ArticleSummaryApplicationI
}

func NewArticleSummaryHandler() *ArticleSummaryHandler {
	return &ArticleSummaryHandler{
		articleSummarySvc: app.NewArticleSummaryApplication(),
	}
}

func (s *ArticleSummaryHandler) Save(ctx context.Context, req *speedy_read.SaveArticleSummaryRequest) (resp *speedy_read.SaveArticleSummaryResponse, err error) {
	if req.ArticleID == 0 {
		klog.CtxErrorf(ctx, "params error")
		return nil, err
	}
	labels := make([]*article_summary.Label, 0)
	for _, desc := range req.Tags {
		labels = append(labels, &article_summary.Label{
			Description: desc,
		})
	}
	id, err := s.articleSummarySvc.CreateArticle(ctx, &article_summary.ArticleSummary{
		Article: &article.Article{
			ID: req.ArticleID,
		},
		LabelList:      labels,
		Title:          req.Title,
		Content:        req.Content,
		ContentSummary: req.ContentSummary,
		Summary:        req.Summary,
		Outline:        req.Outline,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "create article summary error %v", err)
		return nil, err
	}
	return &speedy_read.SaveArticleSummaryResponse{
		ID: id,
	}, nil
}

func (s *ArticleSummaryHandler) ArticleSummaryList(ctx context.Context, req *speedy_read.ArticleSummaryListRequest) (resp *speedy_read.ArticleSummaryListResponse, err error) {
	articleSummaryList, err := s.articleSummarySvc.GetArticleSummaryList(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		klog.CtxErrorf(ctx, "get article summary list error %v", err)
		return nil, err
	}
	summaryList := make([]*speedy_read.ArticleSummary, 0)
	for _, summaryInfo := range articleSummaryList {
		summaryList = append(summaryList, conversion.ArticleSummaryDOToThrift(summaryInfo))
	}
	return &speedy_read.ArticleSummaryListResponse{
		ArticleSummaryList: summaryList,
	}, nil
}
