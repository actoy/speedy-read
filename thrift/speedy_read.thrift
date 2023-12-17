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

service SpeedyRead {
    Response echo(1: Request req)
    // site
    GetSiteResponse GetSiteInfo (1: GetSiteRequest req)
    CreateSiteResponse CreateSiteInfo (1: CreateSiteRequest req)
    // article
    GetArticleListResponse ArticleList(1:GetArticleListRequest req)
    CreateArticleResponse CreateArticle(1:CreateArticleRequest req)
}