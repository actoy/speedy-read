package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"os"
	"speedy/read/biz/infra"
	speedyRead "speedy/read/kitex_gen/speedy_read/speedyread"
)

func main() {
	// init infra
	infra.Init()
	logInit()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
	svr := speedyRead.NewServer(new(SpeedyReadImpl), server.WithServiceAddr(addr))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}

func logInit() {
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	klog.SetOutput(f)
}
