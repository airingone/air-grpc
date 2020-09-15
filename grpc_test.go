package air_grpc

import (
	"context"
	airetcd "github.com/airingone/air-etcd"
	"github.com/airingone/config"
	"github.com/airingone/log"
	"github.com/airingone/pro_proto/helloword"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"net"
	"testing"
	"time"
)

//protoc --go_out=plugins=grpc:. helloword.proro
//sever
func TestGrpcServer(t *testing.T) {
	config.InitConfig()                     //配置文件初始化
	log.InitLog(config.GetLogConfig("log")) //日志初始化
	airetcd.RegisterLocalServerToEtcd(config.GetString("server.name"),
		config.GetUInt32("server.port"), config.GetStringSlice("etcd.addrs")) //将服务注册到etcd集群

	listen, err := net.Listen("tcp", ":"+config.GetString("server.port"))
	if err != nil {
		log.Fatal("[GRPC]: TestGrpcServer listen failed, %+v", err)
	}
	server := grpc.NewServer()
	helloword.RegisterGreeterServer(server, &Server{})
	server.Serve(listen)
}

type Server struct {
}

//服务实现
func (s *Server) SayHello(ctx context.Context, in *helloword.HelloRequest) (*helloword.HelloReply, error) {
	c := NewGrpcContext(ctx, "1234567") //需要在请求包带requestid
	result := hello(c, in.Name)
	return &helloword.HelloReply{Message: result}, nil
}

func hello(ctx *GrpcContext, name string) string {
	ctx.LogHandler.Logger.Info("hello process succ")

	return "Hello" + name
}

//请求
func TestGrpcClient(t *testing.T) {
	config.InitConfig()                     //配置文件初始化
	log.InitLog(config.GetLogConfig("log")) //日志初始化

	var endpoints []string //etcd集群的地址
	endpoints = append(endpoints, "127.0.0.1:2380")
	r := airetcd.NewGrpcResolver(endpoints)
	resolver.Register(r)

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.DialContext(ctx, config.GetString("server.name"),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(roundrobin.Name),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithBlock())
	defer conn.Close()

	req := &helloword.HelloRequest{
		Name: "test",
	}
	c := helloword.NewGreeterClient(conn)
	rsp, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Error("[GRPC]: SayHello Err, %v", err)
		return
	}
	log.Error("[GRPC]: rsp:%+v", rsp)
}
