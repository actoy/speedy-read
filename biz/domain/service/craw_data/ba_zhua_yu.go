package craw_data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/patrickmn/go-cache"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	once           sync.Once
	localCacheUtil *cache.Cache
)

func init() {
	once.Do(func() {
		localCacheUtil = cache.New(20*time.Hour, 10*time.Minute)
	})
}

const (
	tokenUrl        = "https://openapi.bazhuayu.com/token"
	notExportedUrl  = "https://openapi.bazhuayu.com/data/notexported"
	markExportedUrl = "https://openapi.bazhuayu.com/data/markexported"

	tokenCacheKey = "access_token_key"
)

type errorMsg struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// 获取token响应
type token struct {
	Data      tokenData `json:"data"`
	RequestID string    `json:"requestId"`
}

type tokenData struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

// NotExport 未导出数据响应
type NotExport struct {
	Data      NotExportData `json:"data"`
	RequestID string        `json:"requestId"`
}

type NotExportData struct {
	Total   int32        `json:"total"`
	Current int32        `json:"current"`
	Data    []ExportData `json:"data"`
}

type ExportData struct {
	PublishAt      string `json:"publish_at"`
	AuthorName     string `json:"author_name"`
	AuthorUrl      string `json:"author_url"`
	ArticleUrl     string `json:"article_url"`
	ArticleTitle   string `json:"article_title"`
	ArticleContent string `json:"article_content"`
}

// markExport 标记为已导出响应
type markExport struct {
	Data      []ExportData `json:"data,omitempty"`
	Error     *errorMsg    `json:"error,omitempty"`
	RequestID string       `json:"requestId"`
}

// GetNotExportedData
// https://www.cnblogs.com/Xinenhui/p/17496684.html
func GetNotExportedData(ctx context.Context, taskID string) ([]ExportData, error) {
	request, err := http.NewRequest("GET", notExportedUrl, nil)
	if err != nil {
		klog.CtxErrorf(ctx, "http get not exported nre request error is %v", err)
		return nil, err
	}
	token := getApiToken(ctx)
	request.Header.Add("Authorization", token.TokenType+" "+token.AccessToken) //token
	query := request.URL.Query()
	query.Add("taskID", taskID)
	query.Add("size", "1000")
	request.URL.RawQuery = query.Encode()
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		klog.CtxErrorf(ctx, "http get not exported error is %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	resBody := NotExport{}
	_ = json.Unmarshal(body, &resBody)
	return resBody.Data.Data, nil
}

func MarkExported(ctx context.Context, taskID string) bool {
	//发送json格式的参数
	data := map[string]interface{}{
		"taskId": taskID,
	}
	// 序列化
	bytesData, _ := json.Marshal(data)

	//新建请求
	request, err := http.NewRequest("POST", markExportedUrl, strings.NewReader(string(bytesData)))
	if err != nil {
		klog.CtxErrorf(ctx, "http mark exported nre request error is %v", err)
		return false
	}
	token := getApiToken(ctx)
	//请求头部信息
	request.Header.Add("Authorization", token.TokenType+" "+token.AccessToken) //token
	//post formData表单请求
	request.Header.Add("Content-Type", "application/json")

	//实例化一个客户端
	client := &http.Client{}
	//发送请求到服务端
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	resBody := markExport{}
	_ = json.Unmarshal(body, &resBody)
	if resBody.Error != nil {
		return false
	}
	return true
}

func getApiToken(ctx context.Context) tokenData {
	if val, found := localCacheUtil.Get(tokenCacheKey); found {
		klog.CtxInfof(ctx, "get token from cache")
		resp := tokenData{}
		_ = json.Unmarshal([]byte(val.(string)), &resp)
		return resp
	}
	// 参数
	data := make(map[string]interface{})
	data["username"] = "x04mlqa0"
	data["password"] = "@Free4me"
	data["grant_type"] = "password"
	// 序列化
	reqBody, _ := json.Marshal(data)
	resp, err := http.Post(tokenUrl, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		klog.CtxErrorf(ctx, "http post token url error is %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	resBody := token{}
	_ = json.Unmarshal(body, &resBody)
	cacheData, _ := json.Marshal(resBody.Data)
	localCacheUtil.SetDefault(tokenCacheKey, string(cacheData))
	return resBody.Data
}
