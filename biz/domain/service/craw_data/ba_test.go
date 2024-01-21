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
