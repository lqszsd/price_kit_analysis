package HotService

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kit_get_price/hot_symbol/models"
	"net/http"
)

type Service interface {
	HotList(ctx context.Context) (out models.MoneyList, err error)
}

type baseServer struct {
}
func NewService()Service{
	var server Service
	server = &baseServer{}
	return server
}

func (s baseServer)HotList(ctx context.Context)(out models.MoneyList, err error){
	return GetMoneyList(),nil
}

func GetMoneyList() models.MoneyList {
	resp, _ := http.Get("http://127.0.0.1:8080/api/public/stock_hot_rank_em")
	by, _ := ioutil.ReadAll(resp.Body)
	var money models.MoneyList
	json.Unmarshal(by, &money)
	fmt.Println(money)
	return money
}
