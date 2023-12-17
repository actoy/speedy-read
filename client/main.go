package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"log"
	"speedy/read/kitex_gen/speedy_read"
	"speedy/read/kitex_gen/speedy_read/speedyread"
)

func main() {
	client, err := speedyread.NewClient("speedy_read", client.WithHostPorts("127.0.0.1:3000"))
	if err != nil {
		log.Fatal(err)
	}
	//testCreateSite(client)
	testGetSiteList(client)
}

func testCreateSite(client speedyread.Client) {
	createParams := &speedy_read.CreateSiteRequest{
		SourceID:    int64(1),
		SourceType:  "sourceType",
		Url:         "url",
		Description: "desc",
	}
	id, err := client.CreateSiteInfo(context.Background(), createParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(id)
}

func testGetSiteList(client speedyread.Client) {
	siteList, err := client.GetSiteInfo(context.Background(), &speedy_read.GetSiteRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(siteList)
}
