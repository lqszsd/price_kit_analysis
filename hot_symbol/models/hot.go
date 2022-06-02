package models

type MoneyList []struct {
	Rank   int     `json:"当前排名"`
	Code   string  `json:"代码"`
	Name   string  `json:"股票名称"`
	Price  float64 `json:"最新价"`
	Change float64 `json:"涨跌幅"`
}
