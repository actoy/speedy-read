package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"log"
	"net"
	"os"
	"speedy/read/biz/infra"
	speedyRead "speedy/read/kitex_gen/speedy_read/speedyread"

	"github.com/cloudwego/kitex/server"
)

func main() {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
	f, err := os.OpenFile("./output/log/rpc/output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	klog.SetOutput(f)

	// init infra
	infra.Init()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
	svr := speedyRead.NewServer(new(SpeedyReadImpl), server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
