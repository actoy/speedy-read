namespace go speedy_read

const string TypeArticle = "article"
const string TypeNew     = "new"

const string SiteTypeRss = "rss"
const string siteTypeCraw = "craw"

const i32 TradingKong = 0 // 无
const i32 TradingStrongSell = 1 // 强烈卖出
const i32 TradingSell = 2 // 卖出
const i32 TradingMiddle = 3 // 中性
const i32 TradingStrongBuy = 4 // 强烈买入
const i32 TradingBuy = 5 // 买入

struct Request {
  1: string message
}

struct Response {
  1: string message
}

struct GetSiteRequest {
}

struct GetSiteResponse {
  1: required list<SiteInfo> SiteList
}

struct SiteInfo {
  1: required string ID
  2: optional SiteMeta SiteMeta
  3: required string Url
  4: required string Description
  5: required string Tag
  6: required string CreatedAt
  7: required string UpdatedAt
}

struct SiteMeta {
  1: required string ID
  2: required string SiteID
  3: required string MetaType
  4: required string MetaKey
  5: required string MetaValue
  6: required string CreatedAt
  7: required string UpdatedAt
}

struct CreateSiteRequest {
  1: optional string MetaType
  2: optional string MetaKey
  3: optional string MetaValue
  4: required string Url
  5: required string Description
  6: required string Tag
  7: required string Type
  8: required string TypeKey
}

struct CreateSiteResponse {
    1: required string ID
}

struct GetArticleListRequest {
    1: optional list<string> SymbolIdList
    2: optional string ArticleType
    3: optional list<string> SiteIdList
    4: required i32 Limit
    5: required i32 Offset
}

struct GetArticleListResponse {
  1: required list<Article> ArticleList
}

struct Article {
    1: required string ID
    2: required Author Author
    3: required SiteInfo Site
    4: optional list<ArticleMeta> ArticleMetaList
    5: required string Language
    6: required string PublishAt
    7: required string Url      // 原文连接
    8: required string Type     // 原文类型
    9: required string Title    // 原文标题
    10: required string Content // 原文内容
    11: required i32 Status
    12: required i32 Score
    13: required string CreatedAt
    14: required string UpdatedAt
}

struct Author {
	1: required string ID
	2: required string Url
	3: required string AuthorName
	4: required string Image
    5: required string CreatedAt
    6: required string UpdatedAt
}

struct ArticleMeta {
  1: required string ID
  2: required string ArticleID
  3: required string MetaType
  4: required string MetaKey
  5: required string MetaValue
  6: required string CreatedAt
  7: required string UpdatedAt
}

struct CreateArticleRequest {
    1: required Article Article
}


struct CreateArticleResponse {
    1: required string ID
}

struct RejectArticleRequest {
    1: required string ArticleID
}


struct RejectArticleResponse {
    1: required bool Success
}

struct SaveArticleSummaryRequest {
    1: required string ArticleID
    2: required string Title
    3: required string Content // 原文内容
    4: required string Summary // 摘要
    5: required ArticleContentSummary ContentSummary // 一句话原文
    6: required list<SummaryOutline> Outline
    7: required list<string> OutlineString
    8: required list<string> tags
    9: required i32 TradingProposal  // 买卖建议
}

struct SaveArticleSummaryResponse {
    1: required string ID
}

struct ArticleSummaryListRequest {
    1: required i32 Limit
    2: required i32 Offset
    3: optional string Symbol
    4: optional string ArticleType
}

struct ArticleSummary {
    1: required string ID
    2: required Article Article
    3: required string Title
    4: required string Summary // 摘要
    5: required ArticleContentSummary ContentSummary // 一句话原文
    6: required list<SummaryOutline> Outline
    7: required list<string> tags
    8: required string CreatedAt
    9: required i32 TradingProposal  // 买卖建议
    10: optional list<Symbol> SymbolList
}

struct ArticleContentSummary {
    1: required string Original
    2: required string Translation
}

struct SummaryOutline {
    1: required string Title
    2: required string Content
}

struct ArticleSummaryListResponse {
    1: required list<ArticleSummary> ArticleSummaryList
}

struct ArticleCountRequest {
    1: required i32 Status
    2: optional list<string> SymbolIdList
    3: optional string ArticleType
    4: optional list<string> SiteIdList
}

struct ArticleCountResponse {
    1: required i32 Count
}

struct ArticleSummaryCountRequest {
    1: optional string ArticleType
}

struct ArticleSummaryCountResponse {
    1: required i32 Count
}

struct ArticleSummaryDetailRequest {
    1: required string SummaryID
}

struct ArticleSummaryDetailResponse {
    1: required ArticleSummary ArticleSummaryDetail
}

struct SymbolListRequest {
}

struct SymbolListResponse {
  1: required list<Symbol> Symbol
}

struct GetSymbolRequest {
    1: optional string ID
    2: optional string SymbolTag
}

struct GetSybmolResponse {
    1: required Symbol Symbol
}

struct Symbol {
    1: required string ID
    2: required string Symbol
    3: required string Company
    4: required string Source
    5: required string CompanyZH
    6: required string CompanyUrl
    7: required string CompanyAddress
    8: required string Description
    9: required string CompanyBusiness
}

struct SearchSymbolRequest {
    1: required string KeyWord
}

struct SearchSymbolResponse {
    1: list<Symbol> SymbolList
}

struct CrawDataRequest {
    1: optional string Source
}

struct UpdateSymbolRequest {
    1: required string ID
    2: optional string	Company
    3: optional string	CompanyZH
    4: optional string	CompanyUrl
    5: optional string	CompanyAddress
    6: optional string	Description
    7: optional string	CompanyBusiness
}

struct UpdateSymbolResponse{
    1: required bool Success
}

service SpeedyRead {
    Response echo(1: Request req)
    // site
    GetSiteResponse GetSiteInfo (1: GetSiteRequest req) // 获取抓取网站的信息
    CreateSiteResponse CreateSiteInfo (1: CreateSiteRequest req) // 新增抓取网站
    // article
    GetArticleListResponse ArticleList(1:GetArticleListRequest req) // 后台文章列表
    CreateArticleResponse CreateArticle(1:CreateArticleRequest req) // 创建文章，用于接口创建文章原内容
    RejectArticleResponse RejectArticle(1:RejectArticleRequest req) // 审批拒绝此篇文章
    ArticleCountResponse ArticleCount(1: ArticleCountRequest req) //
    // article_summary
    SaveArticleSummaryResponse SaveArticleSummary(1:SaveArticleSummaryRequest req) // 文章总结生成后更新
    ArticleSummaryListResponse GetArticleSummaryList(1: ArticleSummaryListRequest req)
    ArticleSummaryCountResponse ArticleSummaryCount(1: ArticleSummaryCountRequest req) //
    ArticleSummaryDetailResponse ArticleSummaryDetail(1: ArticleSummaryDetailRequest req) //
    // import symbol
    Response importSymbol(1: Request req)
    UpdateSymbolResponse UpdateSymbol(1: UpdateSymbolRequest req)
    SymbolListResponse GetSymbolList(1: SymbolListRequest req)
    SearchSymbolResponse SearchSymbol(1: SearchSymbolRequest req)
    GetSybmolResponse GetSymbol(1: GetSymbolRequest req)
    // craw data
    Response CrawData(1: CrawDataRequest req)
}