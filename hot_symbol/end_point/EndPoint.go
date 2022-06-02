package end_point

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"kit_get_price/hot_symbol/HotService"
	"kit_get_price/userMiddware"
)

type EndPoint struct {
	HotListEndPoint      endpoint.Endpoint
}


//创建端点服务
func NewEndPoint(svc HotService.Service, log *zap.Logger, limiter ratelimit.Limiter) EndPoint {
	var HotListEndPoint endpoint.Endpoint
	{
		HotListEndPoint = NewHotListEndPoint(svc)
		HotListEndPoint = userMiddware.AuthMiddleware(log)(HotListEndPoint)

		HotListEndPoint = userMiddware.NewUberRateMiddleware(limiter)(HotListEndPoint)
	}
	return EndPoint{HotListEndPoint: HotListEndPoint}
}


//实例化结构体 并且实现请求转化
func NewHotListEndPoint(s HotService.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.HotList(ctx)
	}
}
