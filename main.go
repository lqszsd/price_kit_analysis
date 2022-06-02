package main

import (
	"go.uber.org/ratelimit"
	"kit_get_price/PriceService"
	"kit_get_price/end_point"
	"kit_get_price/trans_port"
	"kit_get_price/utils"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	server:=PriceService.NewService(utils.GetLogger())
	uberLimit := ratelimit.New(10)

	endpoints := end_point.NewEndPoint(server, utils.GetLogger(),uberLimit)

	httpHandler := trans_port.NewHttpHandler(endpoints, utils.GetLogger())
	http.ListenAndServe("0.0.0.0:8888", httpHandler)

}
