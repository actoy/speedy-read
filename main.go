package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"speedy/read/biz/infra"
	speedyRead "speedy/read/kitex_gen/speedy_read/speedyread"
)

func main() {
	// init infra
	infra.Init()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
	svr := speedyRead.NewServer(new(SpeedyReadImpl), server.WithServiceAddr(addr))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
