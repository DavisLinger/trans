package main

import (
	"context"
	"encoding/base64"
	"flag"
	"github.com/DavisLinger/trans/cmd"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/DavisLinger/trans/proto"
)

func main2() {
	var server string
	var file string
	flag.StringVar(&server, "server", "127.0.0.1:39329", "服务器端口")
	flag.StringVar(&file, "file", "", "上传的文件")
	flag.Parse()
	log.Println("file:", file)
	log.Println("server:", server)
	cre, err := credentials.NewClientTLSFromFile("../keys/server.pem", "davis")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(cre))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewTransportClient(conn)
	log.Println(f.Name())
	resp, err := client.Trans(context.Background(), &pb.TransportReq{
		FileName: getName(f.Name()),
		FileSize: int64(len(data)),
		Data:     base64.StdEncoding.EncodeToString(data),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("传输完毕,文件名:", resp.FileName)
}
func main() {
	cmd.Execute()
}

func getName(s string) string {
	data := strings.Split(s, `\`)
	p := data[len(data)-1]
	data = strings.Split(p, "/")
	return data[len(data)-1]
}
