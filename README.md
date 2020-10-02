# grpc组件
## 1.组件描述
这里没有做更多封装，在gin_test.go文件给出一个基于etcd发现服务的列子。

## 2.由于etcd支持的grpc版本有限制，需指定版本
1.go mod:
replace google.golang.org/grpc v1.31.1 => google.golang.org/grpc v1.26.0
replace google.golang.org/api v0.15.1 => google.golang.org/api v0.14.0

2.protoc-gen-go降版本，终端输入：
go get github.com/golang/protobuf/protoc-gen-go@v1.2.0
将go/bin/的protoc-gen-go拷贝到/usr/local/bin/protoc-gen-go 

3.重新生成协议文件
protoc --go_out=plugins=grpc:. helloword.proro
