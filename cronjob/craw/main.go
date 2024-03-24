package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"speedy/read/kitex_gen/speedy_read"
	"speedy/read/kitex_gen/speedy_read/speedyread"
)

func main() {
	client, err := speedyread.NewClient("speedy_read", client.WithHostPorts("127.0.0.1:3000"))
	if err != nil {
		klog.Error(err)
	}
	req := &speedy_read.Request{
		Message: "message",
	}
	resp, err := client.Echo(context.Background(), req)
	if err != nil {
		klog.Error(err)
	}
	fmt.Print(resp)
}
