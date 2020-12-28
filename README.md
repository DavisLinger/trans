# Trans

基于 grpc 进行批量传输的服务

## Generate Proto

```shell
 protoc -I  . --go_out=. --plugins=grpc: *.proto
```

## Server

### 编译

```shell
GOOS=linux GOARCH=amd64 go build -o server server.go
```

### 运行

```shell
./server -key xxx.key -sign xxx.pem -port port
```

eg:

```shell
./server -key ../keys/server.key -sign ../keys/server.pem
```

**参数说明**:

key:服务器 key

sign:基于服务器 key 进行签名的文件

port:grpc 服务器端口

## Client

### 编译

```shelll
GOOS=linux GOARCH=amd64 go build -o client client.go
```

### 运行

#### 单个文件传输

```shell
./client -server ip:port -type one -sign xxx.pem -file your_file
```

eg:

```shell
./client -server 127.0.0.1:39329 -type one -sign ./keys/server.pem -file ~/Music/像昨天一样晚安.ape
```

#### 多个文件传输(基于客户端流)

```shell
./client -server ip:port -type batch -sign xxx.pem -list file1 -list file2
```

eg:

```shell
./client -server 127.0.0.1:39329 -type batch -sign ./keys/server.pem -list ~/Music/像昨天一样晚安.ape -list ./client.go
```

#### 基于文件夹传输

注意:这个模式会递归便利所有文件到服务器的统一路径下,慎重使用

```shell
./client -server ip:port -type folder -sign xxx.pem -path your_path
```

eg:

```shell
./client -server 127.0.0.1::39329 -type folder -sign ./keys/server.pem -path ~/Projects/TodoApi
```
