package air_grpc

import (
	"context"
	"github.com/airingone/log"
)

//请求的context，主要将日志统一打requestid等，可以在逻辑函数调用传递
type GrpcContext struct {
	Ctx        context.Context
	LogHandler *log.LogHandler
}

//创建grpc context
func NewGrpcContext(ctx context.Context, requestId string) *GrpcContext {
	grpcCtx := &GrpcContext{
		Ctx:        ctx,
		LogHandler: log.NewLogHandler(),
	}
	grpcCtx.LogHandler.SetRequestId(requestId)

	return grpcCtx
}
