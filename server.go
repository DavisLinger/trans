package main

import (
	"flag"
	"fmt"
	"github.com/DavisLinger/trans/server/trans"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/DavisLinger/trans/proto"
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
	engine := grpc.NewServer(grpc.Creds(cre), grpc.MaxSendMsgSize(math.MaxInt32), grpc.MaxRecvMsgSize(math.MaxInt32))
	pb.RegisterTransportServer(engine, new(trans.TranSrv))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
	err = engine.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
