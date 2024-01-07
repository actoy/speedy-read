namespace go speedy_read

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
    1: i32 Limit
    2: i32 Offset
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
    7: string Url
    8: string Type
    9: string Title
    10: string Content
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
    5: string ContentSummary // 一句话原文
    6: string Outline
    7: list<string> tags
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
    6: string ContentSummary // 一句话原文
    7: string Outline
    8: list<string> tags
    9: string CreatedAt
}

struct ArticleSummaryListResponse {
    1: list<ArticleSummary> ArticleSummaryList
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
    // article_summary
    SaveArticleSummaryResponse SaveArticleSummary(1:SaveArticleSummaryRequest req) // 文章总结生成后更新
    ArticleSummaryListResponse GetArticleSummaryList(1: ArticleSummaryListRequest req)
}