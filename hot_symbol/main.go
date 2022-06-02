package main

import (
	"go.uber.org/ratelimit"
	"kit_get_price/hot_symbol/HotService"
	"kit_get_price/hot_symbol/end_point"
	"kit_get_price/hot_symbol/transport"
	"kit_get_price/utils"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	server:=HotService.NewService()
	uberLimit := ratelimit.New(10)

	endpoints := end_point.NewEndPoint(server, utils.GetLogger(),uberLimit)

	httpHandler := transport.NewHttpHandler(endpoints, utils.GetLogger())
	http.ListenAndServe("0.0.0.0:9999", httpHandler)

}
