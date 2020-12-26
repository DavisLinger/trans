package trans

import (
	"context"
	"encoding/base64"
	"os"
	"runtime"
	"strings"

	pb "github.com/DavisLinger/transport/server/proto"
)

type TranSrv struct {
}

func (t TranSrv) Trans(ctx context.Context, req *pb.TransportReq) (*pb.TransportResp, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer file.Close()
	data, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		return nil, err
	}
	_, err = file.Write(data)
	if err != nil {
		return nil, err
	}
	return &pb.TransportResp{
		FileName: path,
	}, nil
}
