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

service SpeedyRead {
    Response echo(1: Request req)
    GetSiteResponse GetSiteInfo (1: GetSiteRequest req)
    CreateSiteResponse CreateSiteInfo (1: CreateSiteRequest req)
}