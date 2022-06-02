package PriceService

import ("context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"kit_get_price/login/my_model"
	"kit_get_price/utils"
)

type Login struct {
	Account string
	PassWord string
}
const ContextReqUUid = "req_uuid"
type Service interface {
	Login(ctx context.Context,in Login) (out my_model.LoginAck, err error)
}

type baseServer struct {
	logger *zap.Logger
}

func NewService(log *zap.Logger)Service{
	var server Service
	server = &baseServer{log}
	server = NewLogMiddlewareServer(log)(server)
	return server
}


func (s baseServer) Login(ctx context.Context, in Login) (ack my_model.LoginAck, err error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 Price Service", "Login 处理请求"))
	if in.Account != "lq" || in.PassWord != "lq" {
		err = errors.New("用户信息错误")
		return
	}
	ack.Token, err = utils.CreateJwtToken(in.Account,1)
	ack.Code = 200
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)), zap.Any("调用 v3_service Service", "Login 处理请求"), zap.Any("处理返回值", ack))
	return
}

