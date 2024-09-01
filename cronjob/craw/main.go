package main

import (
	"context"
	"fmt"
	"speedy/read/kitex_gen/speedy_read"
	"speedy/read/kitex_gen/speedy_read/speedyread"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	client, err := speedyread.NewClient("speedy_read", client.WithHostPorts("127.0.0.1:3000"))
	if err != nil {
		klog.Error(err)
	}
	source := ""
	req := &speedy_read.CrawDataRequest{
		Source: &source,
	}
	resp, err := client.CrawData(context.Background(), req)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(resp)
}
