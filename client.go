package main

import (
	"flag"
	"log"

	"github.com/DavisLinger/trans/boot"
	"github.com/DavisLinger/trans/handler"
	"github.com/DavisLinger/trans/model"
)

func main() {
	var ServerAddr string
	var Type string
	var File string
	var FileList model.List
	var PemPath string
	var Path string
	flag.StringVar(&ServerAddr, "server", "", "server address")
	flag.StringVar(&Type, "type", "", "type")
	flag.StringVar(&File, "file", "", "file which will be uploaded")
	flag.StringVar(&PemPath, "sign", "", "file which was sign by server")
	flag.StringVar(&Path, "path", "", "path which will be uploaded")
	flag.Var(&FileList, "list", "file list")
	flag.Parse()
	cfg := model.NewConfig(ServerAddr, Type, File, FileList, PemPath, Path)
	log.Println(cfg)
	err := boot.InitClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	switch cfg.Type {
	case "one":
		handler.OneHandler(cfg)
	case "batch":
		handler.BatchHandler(cfg)
	case "folder":
		handler.FolderHandler(cfg)
	default:
		log.Fatal("no such handler")
	}

}
