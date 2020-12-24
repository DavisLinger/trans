package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/DavisLinger/transport/client/proto"
)

func main() {
	var server string
	var file string
	flag.StringVar(&server, "server", "127.0.0.1:39329", "服务器端口")
	flag.StringVar(&file, "file", "", "上传的文件")
	flag.Parse()
	cre, err := credentials.NewClientTLSFromFile("../keys/server.pem", "davis")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(cre))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewTransportClient(conn)

	resp, err := client.Trans(context.Background(), &pb.TransportReq{
		FileName: file,
		FileSize: 100,
		Data:     "xxx.abc",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("传输完毕,文件名:", resp.FileName)
}
