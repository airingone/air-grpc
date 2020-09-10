# air-grpc

etcd不兼容grpc新版本，等支持后再去掉replace：

1.go mod:
replace google.golang.org/grpc v1.31.1 => google.golang.org/grpc v1.26.0
replace google.golang.org/api v0.15.1 => google.golang.org/api v0.14.0

2.protoc-gen-go降版本，终端输入：
go get github.com/golang/protobuf/protoc-gen-go@v1.2.0
将go/bin/的protoc-gen-go拷贝到/usr/local/bin/protoc-gen-go 

3.重新生成协议文件
protoc --go_out=plugins=grpc:. helloword.proro
