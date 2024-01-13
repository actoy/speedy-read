namespace go speedy_read

const string TypeArticle = "article"
const string TypeNew     = "new"
const i32 TradingBearish = 1 // 看空
const i32 TradingBullish = 2 // 看多

struct Request {
  1: string message
}

struct Response {
  1: string message
}

struct GetSiteRequest {
}

struct GetSiteResponse {
  1: list<SiteInfo> SiteList
}

struct SiteInfo {
  1: string ID
  2: SiteMeta SiteMeta
  3: string Url
  4: string Description
  5: string Tag
  6: string CreatedAt
  7: string UpdatedAt
}

struct SiteMeta {
  1: string ID
  2: string SiteID
  3: string MetaType
  4: string MetaKey
  5: string MetaValue
  6: string CreatedAt
  7: string UpdatedAt
}

struct CreateSiteRequest {
  1: string MetaType
  2: string MetaKey
  3: string MetaValue
  4: string Url
  5: string Description
  6: string Tag
}


struct CreateSiteResponse {
    1: string ID
}

struct GetArticleListRequest {
    1: list<string> SiteIdList
    2: string ArticleType
    3: i32 Limit
    4: i32 Offset
}

struct GetArticleListResponse {
  1: list<Article> ArticleList
}

struct Article {
    1: string ID
    2: Author Author
    3: SiteInfo Site
    4: list<ArticleMeta> ArticleMetaList
    5: string Language
    6: string PublishAt
    7: string Url      // 原文连接
    8: string Type     // 原文类型
    9: string Title    // 原文标题
    10: string Content // 原文内容
    11: i32 Status
    12: i32 Score
    13: string CreatedAt
    14: string UpdatedAt
}

struct Author {
	1: string ID
	2: string Url
	3: string AuthorName
	4: string Image
    5: string CreatedAt
    6: string UpdatedAt
}

struct ArticleMeta {
  1: string ID
  2: string ArticleID
  3: string MetaType
  4: string MetaKey
  5: string MetaValue
  6: string CreatedAt
  7: string UpdatedAt
}

struct CreateArticleRequest {
    1: Article Article
}


struct CreateArticleResponse {
    1: string ID
}

struct RejectArticleRequest {
    1: string ArticleID
}


struct RejectArticleResponse {
    1: bool Success
}

struct SaveArticleSummaryRequest {
    1: string ArticleID
    2: string Title
    3: string Content // 原文内容
    4: string Summary // 摘要
    5: ArticleContentSummary ContentSummary // 一句话原文
    6: list<SummaryOutline> Outline
    7: list<string> tags
    8: i32 TradingProposal  // 买卖建议
}

struct SaveArticleSummaryResponse {
    1: string ID
}

struct ArticleSummaryListRequest {
    1: i32 Limit
    2: i32 Offset
}

struct ArticleSummary {
    1: string ID
    2: Article Article
    3: string Title
    4: string Content // 原文内容
    5: string Summary // 摘要
    6: ArticleContentSummary ContentSummary // 一句话原文
    7: list<SummaryOutline> Outline
    8: list<string> tags
    9: string CreatedAt
    10: i32 TradingProposal  // 买卖建议
}

struct ArticleContentSummary {
    1: string Original
    2: string Translation
}

struct SummaryOutline {
    1: string Title
    2: string Content
}

struct ArticleSummaryListResponse {
    1: list<ArticleSummary> ArticleSummaryList
}

struct ArticleCountRequest {
    1: i32 Status
}

struct ArticleCountResponse {
    1: i32 Count
}

struct ArticleSummaryCountRequest {

}

struct ArticleSummaryCountResponse {
    1: i32 Count
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
}