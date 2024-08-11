// Code generated by Kitex v0.8.0. DO NOT EDIT.

package speedyread

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	speedy_read "speedy/read/kitex_gen/speedy_read"
)

func serviceInfo() *kitex.ServiceInfo {
	return speedyReadServiceInfo
}

var speedyReadServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SpeedyRead"
	handlerType := (*speedy_read.SpeedyRead)(nil)
	methods := map[string]kitex.MethodInfo{
		"echo":                  kitex.NewMethodInfo(echoHandler, newSpeedyReadEchoArgs, newSpeedyReadEchoResult, false),
		"GetSiteInfo":           kitex.NewMethodInfo(getSiteInfoHandler, newSpeedyReadGetSiteInfoArgs, newSpeedyReadGetSiteInfoResult, false),
		"CreateSiteInfo":        kitex.NewMethodInfo(createSiteInfoHandler, newSpeedyReadCreateSiteInfoArgs, newSpeedyReadCreateSiteInfoResult, false),
		"ArticleList":           kitex.NewMethodInfo(articleListHandler, newSpeedyReadArticleListArgs, newSpeedyReadArticleListResult, false),
		"CreateArticle":         kitex.NewMethodInfo(createArticleHandler, newSpeedyReadCreateArticleArgs, newSpeedyReadCreateArticleResult, false),
		"RejectArticle":         kitex.NewMethodInfo(rejectArticleHandler, newSpeedyReadRejectArticleArgs, newSpeedyReadRejectArticleResult, false),
		"ArticleCount":          kitex.NewMethodInfo(articleCountHandler, newSpeedyReadArticleCountArgs, newSpeedyReadArticleCountResult, false),
		"SaveArticleSummary":    kitex.NewMethodInfo(saveArticleSummaryHandler, newSpeedyReadSaveArticleSummaryArgs, newSpeedyReadSaveArticleSummaryResult, false),
		"GetArticleSummaryList": kitex.NewMethodInfo(getArticleSummaryListHandler, newSpeedyReadGetArticleSummaryListArgs, newSpeedyReadGetArticleSummaryListResult, false),
		"ArticleSummaryCount":   kitex.NewMethodInfo(articleSummaryCountHandler, newSpeedyReadArticleSummaryCountArgs, newSpeedyReadArticleSummaryCountResult, false),
		"ArticleSummaryDetail":  kitex.NewMethodInfo(articleSummaryDetailHandler, newSpeedyReadArticleSummaryDetailArgs, newSpeedyReadArticleSummaryDetailResult, false),
		"importSymbol":          kitex.NewMethodInfo(importSymbolHandler, newSpeedyReadImportSymbolArgs, newSpeedyReadImportSymbolResult, false),
		"UpdateSymbol":          kitex.NewMethodInfo(updateSymbolHandler, newSpeedyReadUpdateSymbolArgs, newSpeedyReadUpdateSymbolResult, false),
		"GetSymbolList":         kitex.NewMethodInfo(getSymbolListHandler, newSpeedyReadGetSymbolListArgs, newSpeedyReadGetSymbolListResult, false),
		"SearchSymbol":          kitex.NewMethodInfo(searchSymbolHandler, newSpeedyReadSearchSymbolArgs, newSpeedyReadSearchSymbolResult, false),
		"CrawData":              kitex.NewMethodInfo(crawDataHandler, newSpeedyReadCrawDataArgs, newSpeedyReadCrawDataResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "speedy_read",
		"ServiceFilePath": `thrift/speedy_read.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadEchoArgs)
	realResult := result.(*speedy_read.SpeedyReadEchoResult)
	success, err := handler.(speedy_read.SpeedyRead).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadEchoArgs() interface{} {
	return speedy_read.NewSpeedyReadEchoArgs()
}

func newSpeedyReadEchoResult() interface{} {
	return speedy_read.NewSpeedyReadEchoResult()
}

func getSiteInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadGetSiteInfoArgs)
	realResult := result.(*speedy_read.SpeedyReadGetSiteInfoResult)
	success, err := handler.(speedy_read.SpeedyRead).GetSiteInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadGetSiteInfoArgs() interface{} {
	return speedy_read.NewSpeedyReadGetSiteInfoArgs()
}

func newSpeedyReadGetSiteInfoResult() interface{} {
	return speedy_read.NewSpeedyReadGetSiteInfoResult()
}

func createSiteInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadCreateSiteInfoArgs)
	realResult := result.(*speedy_read.SpeedyReadCreateSiteInfoResult)
	success, err := handler.(speedy_read.SpeedyRead).CreateSiteInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadCreateSiteInfoArgs() interface{} {
	return speedy_read.NewSpeedyReadCreateSiteInfoArgs()
}

func newSpeedyReadCreateSiteInfoResult() interface{} {
	return speedy_read.NewSpeedyReadCreateSiteInfoResult()
}

func articleListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadArticleListArgs)
	realResult := result.(*speedy_read.SpeedyReadArticleListResult)
	success, err := handler.(speedy_read.SpeedyRead).ArticleList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadArticleListArgs() interface{} {
	return speedy_read.NewSpeedyReadArticleListArgs()
}

func newSpeedyReadArticleListResult() interface{} {
	return speedy_read.NewSpeedyReadArticleListResult()
}

func createArticleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadCreateArticleArgs)
	realResult := result.(*speedy_read.SpeedyReadCreateArticleResult)
	success, err := handler.(speedy_read.SpeedyRead).CreateArticle(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadCreateArticleArgs() interface{} {
	return speedy_read.NewSpeedyReadCreateArticleArgs()
}

func newSpeedyReadCreateArticleResult() interface{} {
	return speedy_read.NewSpeedyReadCreateArticleResult()
}

func rejectArticleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadRejectArticleArgs)
	realResult := result.(*speedy_read.SpeedyReadRejectArticleResult)
	success, err := handler.(speedy_read.SpeedyRead).RejectArticle(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadRejectArticleArgs() interface{} {
	return speedy_read.NewSpeedyReadRejectArticleArgs()
}

func newSpeedyReadRejectArticleResult() interface{} {
	return speedy_read.NewSpeedyReadRejectArticleResult()
}

func articleCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadArticleCountArgs)
	realResult := result.(*speedy_read.SpeedyReadArticleCountResult)
	success, err := handler.(speedy_read.SpeedyRead).ArticleCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadArticleCountArgs() interface{} {
	return speedy_read.NewSpeedyReadArticleCountArgs()
}

func newSpeedyReadArticleCountResult() interface{} {
	return speedy_read.NewSpeedyReadArticleCountResult()
}

func saveArticleSummaryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadSaveArticleSummaryArgs)
	realResult := result.(*speedy_read.SpeedyReadSaveArticleSummaryResult)
	success, err := handler.(speedy_read.SpeedyRead).SaveArticleSummary(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadSaveArticleSummaryArgs() interface{} {
	return speedy_read.NewSpeedyReadSaveArticleSummaryArgs()
}

func newSpeedyReadSaveArticleSummaryResult() interface{} {
	return speedy_read.NewSpeedyReadSaveArticleSummaryResult()
}

func getArticleSummaryListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadGetArticleSummaryListArgs)
	realResult := result.(*speedy_read.SpeedyReadGetArticleSummaryListResult)
	success, err := handler.(speedy_read.SpeedyRead).GetArticleSummaryList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadGetArticleSummaryListArgs() interface{} {
	return speedy_read.NewSpeedyReadGetArticleSummaryListArgs()
}

func newSpeedyReadGetArticleSummaryListResult() interface{} {
	return speedy_read.NewSpeedyReadGetArticleSummaryListResult()
}

func articleSummaryCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadArticleSummaryCountArgs)
	realResult := result.(*speedy_read.SpeedyReadArticleSummaryCountResult)
	success, err := handler.(speedy_read.SpeedyRead).ArticleSummaryCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadArticleSummaryCountArgs() interface{} {
	return speedy_read.NewSpeedyReadArticleSummaryCountArgs()
}

func newSpeedyReadArticleSummaryCountResult() interface{} {
	return speedy_read.NewSpeedyReadArticleSummaryCountResult()
}

func articleSummaryDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadArticleSummaryDetailArgs)
	realResult := result.(*speedy_read.SpeedyReadArticleSummaryDetailResult)
	success, err := handler.(speedy_read.SpeedyRead).ArticleSummaryDetail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadArticleSummaryDetailArgs() interface{} {
	return speedy_read.NewSpeedyReadArticleSummaryDetailArgs()
}

func newSpeedyReadArticleSummaryDetailResult() interface{} {
	return speedy_read.NewSpeedyReadArticleSummaryDetailResult()
}

func importSymbolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadImportSymbolArgs)
	realResult := result.(*speedy_read.SpeedyReadImportSymbolResult)
	success, err := handler.(speedy_read.SpeedyRead).ImportSymbol(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadImportSymbolArgs() interface{} {
	return speedy_read.NewSpeedyReadImportSymbolArgs()
}

func newSpeedyReadImportSymbolResult() interface{} {
	return speedy_read.NewSpeedyReadImportSymbolResult()
}

func updateSymbolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadUpdateSymbolArgs)
	realResult := result.(*speedy_read.SpeedyReadUpdateSymbolResult)
	success, err := handler.(speedy_read.SpeedyRead).UpdateSymbol(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadUpdateSymbolArgs() interface{} {
	return speedy_read.NewSpeedyReadUpdateSymbolArgs()
}

func newSpeedyReadUpdateSymbolResult() interface{} {
	return speedy_read.NewSpeedyReadUpdateSymbolResult()
}

func getSymbolListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadGetSymbolListArgs)
	realResult := result.(*speedy_read.SpeedyReadGetSymbolListResult)
	success, err := handler.(speedy_read.SpeedyRead).GetSymbolList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadGetSymbolListArgs() interface{} {
	return speedy_read.NewSpeedyReadGetSymbolListArgs()
}

func newSpeedyReadGetSymbolListResult() interface{} {
	return speedy_read.NewSpeedyReadGetSymbolListResult()
}

func searchSymbolHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadSearchSymbolArgs)
	realResult := result.(*speedy_read.SpeedyReadSearchSymbolResult)
	success, err := handler.(speedy_read.SpeedyRead).SearchSymbol(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadSearchSymbolArgs() interface{} {
	return speedy_read.NewSpeedyReadSearchSymbolArgs()
}

func newSpeedyReadSearchSymbolResult() interface{} {
	return speedy_read.NewSpeedyReadSearchSymbolResult()
}

func crawDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*speedy_read.SpeedyReadCrawDataArgs)
	realResult := result.(*speedy_read.SpeedyReadCrawDataResult)
	success, err := handler.(speedy_read.SpeedyRead).CrawData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSpeedyReadCrawDataArgs() interface{} {
	return speedy_read.NewSpeedyReadCrawDataArgs()
}

func newSpeedyReadCrawDataResult() interface{} {
	return speedy_read.NewSpeedyReadCrawDataResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, req *speedy_read.Request) (r *speedy_read.Response, err error) {
	var _args speedy_read.SpeedyReadEchoArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetSiteInfo(ctx context.Context, req *speedy_read.GetSiteRequest) (r *speedy_read.GetSiteResponse, err error) {
	var _args speedy_read.SpeedyReadGetSiteInfoArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadGetSiteInfoResult
	if err = p.c.Call(ctx, "GetSiteInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateSiteInfo(ctx context.Context, req *speedy_read.CreateSiteRequest) (r *speedy_read.CreateSiteResponse, err error) {
	var _args speedy_read.SpeedyReadCreateSiteInfoArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadCreateSiteInfoResult
	if err = p.c.Call(ctx, "CreateSiteInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ArticleList(ctx context.Context, req *speedy_read.GetArticleListRequest) (r *speedy_read.GetArticleListResponse, err error) {
	var _args speedy_read.SpeedyReadArticleListArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadArticleListResult
	if err = p.c.Call(ctx, "ArticleList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateArticle(ctx context.Context, req *speedy_read.CreateArticleRequest) (r *speedy_read.CreateArticleResponse, err error) {
	var _args speedy_read.SpeedyReadCreateArticleArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadCreateArticleResult
	if err = p.c.Call(ctx, "CreateArticle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RejectArticle(ctx context.Context, req *speedy_read.RejectArticleRequest) (r *speedy_read.RejectArticleResponse, err error) {
	var _args speedy_read.SpeedyReadRejectArticleArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadRejectArticleResult
	if err = p.c.Call(ctx, "RejectArticle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ArticleCount(ctx context.Context, req *speedy_read.ArticleCountRequest) (r *speedy_read.ArticleCountResponse, err error) {
	var _args speedy_read.SpeedyReadArticleCountArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadArticleCountResult
	if err = p.c.Call(ctx, "ArticleCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SaveArticleSummary(ctx context.Context, req *speedy_read.SaveArticleSummaryRequest) (r *speedy_read.SaveArticleSummaryResponse, err error) {
	var _args speedy_read.SpeedyReadSaveArticleSummaryArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadSaveArticleSummaryResult
	if err = p.c.Call(ctx, "SaveArticleSummary", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetArticleSummaryList(ctx context.Context, req *speedy_read.ArticleSummaryListRequest) (r *speedy_read.ArticleSummaryListResponse, err error) {
	var _args speedy_read.SpeedyReadGetArticleSummaryListArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadGetArticleSummaryListResult
	if err = p.c.Call(ctx, "GetArticleSummaryList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ArticleSummaryCount(ctx context.Context, req *speedy_read.ArticleSummaryCountRequest) (r *speedy_read.ArticleSummaryCountResponse, err error) {
	var _args speedy_read.SpeedyReadArticleSummaryCountArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadArticleSummaryCountResult
	if err = p.c.Call(ctx, "ArticleSummaryCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ArticleSummaryDetail(ctx context.Context, req *speedy_read.ArticleSummaryDetailRequest) (r *speedy_read.ArticleSummaryDetailResponse, err error) {
	var _args speedy_read.SpeedyReadArticleSummaryDetailArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadArticleSummaryDetailResult
	if err = p.c.Call(ctx, "ArticleSummaryDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ImportSymbol(ctx context.Context, req *speedy_read.Request) (r *speedy_read.Response, err error) {
	var _args speedy_read.SpeedyReadImportSymbolArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadImportSymbolResult
	if err = p.c.Call(ctx, "importSymbol", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateSymbol(ctx context.Context, req *speedy_read.UpdateSymbolRequest) (r *speedy_read.UpdateSymbolResponse, err error) {
	var _args speedy_read.SpeedyReadUpdateSymbolArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadUpdateSymbolResult
	if err = p.c.Call(ctx, "UpdateSymbol", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetSymbolList(ctx context.Context, req *speedy_read.SymbolListRequest) (r *speedy_read.SymbolListResponse, err error) {
	var _args speedy_read.SpeedyReadGetSymbolListArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadGetSymbolListResult
	if err = p.c.Call(ctx, "GetSymbolList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchSymbol(ctx context.Context, req *speedy_read.SearchSymbolRequest) (r *speedy_read.SearchSymbolResponse, err error) {
	var _args speedy_read.SpeedyReadSearchSymbolArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadSearchSymbolResult
	if err = p.c.Call(ctx, "SearchSymbol", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CrawData(ctx context.Context, req *speedy_read.CrawDataRequest) (r *speedy_read.Response, err error) {
	var _args speedy_read.SpeedyReadCrawDataArgs
	_args.Req = req
	var _result speedy_read.SpeedyReadCrawDataResult
	if err = p.c.Call(ctx, "CrawData", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
