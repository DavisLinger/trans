package trans

import (
	pb "github.com/DavisLinger/transport/server/proto"
)

type TranSrv struct {
	srv *TranServer
}

func (t TranSrv) Trans(in *pb.TransportServer) (*pb.TransportResp, error) {
	return nil, nil
}

func NewStreamServer() TranSrv {
	return TranSrv{
		srv: new(TranServer),
	}
}

type TranServer struct {
}

func (t *TranServer) SendAndClose(pb *pb.TransportResp) error {
	return nil
}
func (t *TranServer) Recv() (*pb.TransportReq, error) {
	return nil, nil
}
