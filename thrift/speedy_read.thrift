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
  1: i64 ID
  2: i64 SourceID
  3: string SourceType
  4: string Url
  5: string Description
  6: string CreatedAt
  7: string UpdatedAt
}

struct CreateSiteRequest {
  1: i64 SourceID
  2: string SourceType
  3: string Url
  4: string Description
}


struct CreateSiteResponse {
    1: i64 ID
}

struct GetArticleListRequest {
    1: i32 Limit
    2: i32 Offset
}

struct GetArticleListResponse {
  1: list<Article> ArticleList
}

struct Article {
    1: i64 ID
    2: Author Author
    3: SiteInfo Site
    4: string Language
    5: string PublishAt
    6: string Url
    7: string Type
    8: string Title
    9: string Content
    10: i32 Status
    11: i32 Score
    12: string CreatedAt
    13: string UpdatedAt
}

struct Author {
	1: i64 ID
	2: string Url
	3: string AuthorName
	4: string Image
    5: string CreatedAt
    6: string UpdatedAt
}

struct CreateArticleRequest {
    1: Article Article
}


struct CreateArticleResponse {
    1: i64 ID
}

struct RejectArticleRequest {
    1: i64 ArticleID
}


struct RejectArticleResponse {
    1: bool Success
}

struct SaveArticleSummaryRequest {
    1: i64 ArticleID
    2: string Title
    3: string Content // 原文内容
    4: string Summary // 摘要
    5: string ContentSummary // 一句话原文
    6: string Outline
    7: list<string> tags
}

struct SaveArticleSummaryResponse {
    1: i64 ID
}

struct ArticleSummaryListRequest {
    1: i32 Limit
    2: i32 Offset
}

struct ArticleSummary {
    1: i64 ID
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