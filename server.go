package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/DavisLinger/trans/proto"
	"github.com/DavisLinger/trans/service"
)

func main() {
	var port int
	var key string
	var pem string
	flag.IntVar(&port, "port", 39329, "服务器端口")
	flag.StringVar(&key, "key", ".", "server key")
	flag.StringVar(&pem, "sign", ".", "server pem")
	flag.Parse()
	fmt.Println("grpc server port:", port)
	fmt.Println("grpc server key:", key)
	fmt.Println("grpc server pem:", pem)
	if key == "" || pem == "" {
		log.Fatalf("加载证书失败,key:%v,pem:%v", key, pem)
	}
	// boot grpc server
	cre, err := credentials.NewServerTLSFromFile(pem, key)
	if err != nil {
		log.Fatal("加载证书失败,err:", err)
	}
	var engine *grpc.Server
	// 8<<40 means 8GB size
	// 8<<40 == 8*1024*1024*1024*1024
	engine = grpc.NewServer(grpc.Creds(cre), grpc.MaxSendMsgSize(8<<40), grpc.MaxRecvMsgSize(8<<40))
	pb.RegisterTransportServer(engine, new(service.TranSrv))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
	err = engine.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
