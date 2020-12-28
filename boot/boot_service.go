package boot

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/DavisLinger/trans/model"
	pb "github.com/DavisLinger/trans/proto"
)

var Client pb.TransportClient

func InitClient(cfg model.Config) error {
	cre, err := credentials.NewClientTLSFromFile(cfg.PemPath, "davis")
	if err != nil {
		return err
	}
	conn, err := grpc.Dial(cfg.ServerAddr, grpc.WithTransportCredentials(cre))
	if err != nil {
		return err
	}
	Client = pb.NewTransportClient(conn)
	return nil
}
