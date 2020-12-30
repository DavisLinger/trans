package service

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	pb "github.com/DavisLinger/trans/proto"
)

type TranSrv struct {
}

func (t TranSrv) BatchUpLoad(stream pb.Transport_BatchUpLoadServer) error {
	var fileList = make([]string, 0)
	var index int
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.BatchTranResp{
					FileName: fileList,
				})
			} else {
				return err
			}
		}
		log.Printf("stream receive index:%v,file_name:%v", index, req.FileName)
		index++
		path, err := t.writeFile(req)
		if err != nil {
			return err
		}
		fileList = append(fileList, path)
	}
}

func (t TranSrv) Trans(_ context.Context, req *pb.TransportReq) (*pb.TransportResp, error) {
	path, err := t.writeFile(req)
	if err != nil {
		return nil, err
	}
	return &pb.TransportResp{
		FileName: path,
	}, nil
}

func (t TranSrv) writeFile(req *pb.TransportReq) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	pre := ""
	if strings.ToLower(runtime.GOOS) == "windows" {
		pre = `\`
	} else {
		pre = `/`
	}
	path := pwd + pre + req.FileName
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	data, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		return "", err

	}
	_, err = file.Write(data)
	if err != nil {
		return "", err
	}
	return path, nil
}
