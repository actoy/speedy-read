package craw_data

import (
	"context"
	"fmt"
	"testing"
)

func TestGetApiToken(t *testing.T) {
	token := getApiToken(context.Background())
	fmt.Println(token)
	token = getApiToken(context.Background())
	fmt.Println(token)
}

func TestGetNotExportedData(t *testing.T) {
	resp, _ := GetNotExportedData(context.Background(), "d0470acc-8bc9-c5b0-cdf9-8251f4e922af")
	fmt.Println(resp)
}

func TestUpdateItems(t *testing.T) {
	ctx := context.Background()
	taskID := "ad024688-a5a7-0dac-3fb2-585051d9baed"
	data, _ := GetNotExportedData(ctx, taskID)
	fmt.Println(data)
	item := []string{
		"https://www.marketbeat.com/originals/5-top-healthcare-stocks-for-earnings-growth-in-2024/",
		"https://www.marketbeat.com/originals/intuitive-surgicals-post-earnings-dip-is-a-healthy-time-to-buy/",
	}
	UpdateItems(ctx, taskID, item)
	Start(ctx, taskID)
	data, _ = GetNotExportedData(ctx, taskID)
	fmt.Println(data)
}
