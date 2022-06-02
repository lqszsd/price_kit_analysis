package trans_port

import (
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"kit_get_price/login/PriceService"
	"kit_get_price/login/end_point"
	"kit_get_price/utils"
	"net/http"
)

func NewHttpHandler(endpoint end_point.EndPoint, log *zap.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder), //程序中的全部报错都会走这里面
		httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			UUID := uuid.NewV5(uuid.NewV4(), "req_uuid").String()
			log.Debug("给请求添加uuid", zap.Any("UUID", UUID))
			ctx = context.WithValue(ctx, PriceService.ContextReqUUid, UUID)
			ctx = context.WithValue(ctx, utils.JWT_CONTEXT_KEY, request.Header.Get("Authorization"))
			log.Debug("给请求添加uuid", zap.Any("UUID", UUID))
			return ctx
		}),
	}
	m := http.NewServeMux()
	m.Handle("/login", httptransport.NewServer(
		endpoint.LoginEndPoint,
		decodeHTTPLoginRequest,    //解析请求值
		encodeHTTPGenericResponse, //返回值
		options...,
	))
	return m
}


func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	fmt.Println("errorEncoder", err.Error())
	w.Header().Set("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error(),Code:"400"})
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	utils.GetLogger().Debug(fmt.Sprint(ctx.Value(PriceService.ContextReqUUid)), zap.Any("请求结束封装返回值", response))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	return json.NewEncoder(w).Encode(response)
}


func decodeHTTPLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var login PriceService.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		return nil, err
	}
	utils.GetLogger().Debug(fmt.Sprint(ctx.Value(PriceService.ContextReqUUid)), zap.Any(" 开始解析请求数据", login))
	fmt.Println(4444,login)
	return login, nil
}
type errorWrapper struct {
	Code string `json:"code"`
	Error string `json:"errors"`
}
