package userMiddware

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"kit_get_price/utils"
)

func AuthMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			token := fmt.Sprint(ctx.Value(utils.JWT_CONTEXT_KEY))
			var data map[string]string
			data=make(map[string]string,0)
			data["code"]="200";
			data["msg"]="请登录"
			if token == "" {
				err = errors.New("请登录")
				return data, err
			}
			jwtInfo, err := utils.ParseToken(token)
			if err != nil {
				return data, err
			}
			if v, ok := jwtInfo["Name"]; ok {
				ctx = context.WithValue(ctx, "name", v)
			}
			return next(ctx, request)
		}
	}
}
func NewUberRateMiddleware(limit ratelimit.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			limit.Take()
			return next(ctx, request)
		}
	}
}
