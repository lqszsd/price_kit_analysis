package PriceService

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"kit_get_price/login/my_model"
)
type NewMiddlewareServer func(service Service) Service

type logMiddlewareServer struct {
	logger *zap.Logger
	next   Service
}

func (l logMiddlewareServer) Login(ctx context.Context,in Login)(out my_model.LoginAck, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint("test"), zap.Any("调用 Login logMiddlewareServer", "Login"), zap.Any("req", in), zap.Any("res", out), zap.Any("err", err))
	}()
	out, err = l.next.Login(ctx, in)
	return
}

func NewLogMiddlewareServer(log *zap.Logger) NewMiddlewareServer {
	return func(service Service) Service {
		return logMiddlewareServer{
			logger: log,
			next:   service,
		}
	}
}
