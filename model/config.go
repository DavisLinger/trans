package model

import "fmt"

type Config struct {
	ServerAddr string
	Type       string
	File       string
	FileList   []string
	PemPath    string
	Path       string
}

func NewConfig(ServerAddr string, Type string, File string, FileList []string, PemPath string, path string) Config {
	return Config{
		ServerAddr: ServerAddr,
		Type:       Type,
		File:       File,
		FileList:   FileList,
		PemPath:    PemPath,
		Path:       path,
	}
}

type List []string

func (l *List) String() string {
	return fmt.Sprintf("%v", []string(*l))
}

func (l *List) Set(str string) error {
	*l = append(*l, str)
	return nil
}
