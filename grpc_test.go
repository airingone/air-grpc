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

	//每个服务全局注册一次
	etcdConfig := config.GetGrpcConfig("grpc_test")
	r := airetcd.NewGrpcResolver(config.GetEtcdConfig("etcd").Addrs)
	resolver.Register(r)

	//conn初始化一次即可，grpc会维护连接
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(etcdConfig.TimeOutMs)*time.Millisecond)
	conn, err := grpc.DialContext(ctx, etcdConfig.Name, //obejct会传给etcd作为watch对象
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
