package end_point

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"kit_get_price/login/PriceService"
	"kit_get_price/userMiddware"
)

type EndPoint struct {
	LoginEndPoint      endpoint.Endpoint
}

//创建端点服务
func NewEndPoint(svc PriceService.Service, log *zap.Logger, limiter ratelimit.Limiter) EndPoint {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = NewLoginEndPoint(svc)

		loginEndPoint = userMiddware.NewUberRateMiddleware(limiter)(loginEndPoint)
	}
	return EndPoint{LoginEndPoint: loginEndPoint}
}

//实例化结构体 并且实现请求转化
func NewLoginEndPoint(s PriceService.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PriceService.Login)
		return s.Login(ctx, req)
	}
}

