package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net"


	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/DavisLinger/transport/proto"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 39329, "服务器端口")
	flag.Parse()
	fmt.Println("grpc server port:", port)
	// boot grpc server
	cre, err := credentials.NewServerTLSFromFile("../keys/server.pem", "../keys/server.key")
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