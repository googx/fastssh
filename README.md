# fastssh


#### 项目模块 

###### 主机信息管理模块

1. 创建项目 
```
    创建服务
    micro new --gopath=false --namespace=top.goox --plugin=registry=etcd:broker=nats --type=srv hosts
    
    Output:
Creating service top.goox.srv.hosts in hosts

.
├── main.go
├── plugin.go
├── handler
│   └── hosts.go
├── subscriber
│   └── hosts.go
├── proto/hosts
│   └── hosts.proto
├── Dockerfile
├── Makefile
└── README.md


download protobuf for micro:

brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro

compile the proto file hosts.proto:

cd hosts
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. proto/hosts/hosts.proto
```
