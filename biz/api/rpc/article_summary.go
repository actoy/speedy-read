package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/biz/app"
	"speedy/read/biz/conversion"
	"speedy/read/biz/domain/aggregates/article"
	"speedy/read/biz/domain/aggregates/article_summary"
	"speedy/read/biz/utils"
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
	if len(req.ArticleID) == 0 {
		klog.CtxErrorf(ctx, "params error")
		return nil, err
	}
	id, err := s.articleSummarySvc.CreateArticle(ctx, &article_summary.ArticleSummary{
		Article: &article.Article{
			ID: utils.StringToInt64(req.ArticleID),
		},
		LabelList:       conversion.CovertLabelToDO(req.Tags),
		Title:           req.Title,
		ContentSummary:  conversion.CovertSummaryContentToDO(req.ContentSummary),
		Summary:         req.Summary,
		Outline:         conversion.CovertSummaryOutlineListToDO(req.Outline),
		TradingProposal: req.TradingProposal,
	}, req.Content)
	if err != nil {
		klog.CtxErrorf(ctx, "create article summary error %v", err)
		return nil, err
	}
	return &speedy_read.SaveArticleSummaryResponse{
		ID: utils.Int64ToString(id),
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

func (s *ArticleSummaryHandler) ArticleSummaryCount(ctx context.Context, req *speedy_read.ArticleSummaryCountRequest) (resp *speedy_read.ArticleSummaryCountResponse, err error) {
	count := s.articleSummarySvc.ArticleSummaryCount(ctx)
	return &speedy_read.ArticleSummaryCountResponse{
		Count: count,
	}, nil
}
