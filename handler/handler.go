package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/DavisLinger/trans/boot"
	"github.com/DavisLinger/trans/model"
	pb "github.com/DavisLinger/trans/proto"
)

func OneHandler(cfg model.Config) {
	timer := time.Now()
	if cfg.File == "" {
		log.Panicln("invalid value of flag file")
	}
	name := getName(cfg.File)
	file, err := os.Open(cfg.File)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panicln(err)
	}
	str := base64.StdEncoding.EncodeToString(data)
	resp, err := boot.Client.Trans(context.Background(), &pb.TransportReq{
		FileName: name,
		FileSize: int64(len(data)),
		Data:     str,
	})
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("上传文件:%v成功,大小:%v,耗时:%v\n", resp.FileName, len(data), time.Since(timer))
}

func BatchHandler(cfg model.Config) {
	if len(cfg.FileList) == 0 {
		log.Panicln("zero length file list")
	}
	stream, err := boot.Client.BatchUpLoad(context.Background())
	if err != nil {
		log.Panicln(err)
	}
	for _, val := range cfg.FileList {
		name := getName(val)
		file, err := os.Open(val)
		if err != nil {
			log.Panicln(err)
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			file.Close()
			log.Panicln(err)
		}
		str := base64.StdEncoding.EncodeToString(data)
		err = stream.Send(&pb.TransportReq{
			FileName: name,
			FileSize: int64(len(data)),
			Data:     str,
		})
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.FileName)
}
func FolderHandler(cfg model.Config) {
	log.Println("folder handler~")
	if cfg.Path == "" {
		log.Fatal("invalid params of path")
	}
	stream, err := boot.Client.BatchUpLoad(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("cfg_path", cfg.Path)
	err = filepath.Walk(cfg.Path, func(path string, info os.FileInfo, err error) error {
		log.Println("walk_path", path)
		log.Printf("is_dir:%v,path:%v,name:%v", info.IsDir(), path, info.Name())
		if info == nil {
			return nil
		}
		if !info.IsDir() {
			log.Println("name:", info.Name())
			str, l, err := readFile2Str(path)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println("stream send:", info.Name())
			err = stream.Send(&pb.TransportReq{
				FileName: info.Name(),
				FileSize: int64(l),
				Data:     str,
			})
			if err != nil {
				log.Println(err)
				return err
			}
		}
		//return nil
		//} else {
		//
		//}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("上传列表:", res)
}

func getName(s string) string {
	_, file := path.Split(s)
	return file
}

func readFile2Str(s string) (string, int, error) {
	log.Println("path:", s)
	f, err := os.Open(s)
	if err != nil {
		return "", 0, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", 0, err
	}
	return base64.StdEncoding.EncodeToString(data), len(data), nil
}
