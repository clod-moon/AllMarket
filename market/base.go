package market

import (
	"github.com/bitly/go-simplejson"
	"strings"
	"fmt"
	"AllMarket/model"
)

type JSON = simplejson.Json

var (//
	TickerList []string
)

func getAllTicker() {
	i:= 1
	for value, _ := range model.StandardBiMap {
		for v, _ := range model.DealBiMap {
			if v != value {
				var huobi model.Huobi
				var bian model.Bian
				var okex model.Okex

				huobi.StandardBiId = model.StandardBiMap[value]
				bian.StandardBiId = model.StandardBiMap[value]
				okex.StandardBiId = model.StandardBiMap[value]

				huobi.DealBiId = model.DealBiMap[v]
				bian.DealBiId = model.DealBiMap[v]
				okex.DealBiId = model.DealBiMap[v]

				value = strings.ToLower(value)
				v = strings.ToLower(v)
				ticker := fmt.Sprintf("%s%s", v, value)
				TickerList = append(TickerList, ticker)
				model.HuobiMap[ticker] = huobi
				model.BianMap[ticker] = bian
				model.OkexMap[ticker] = okex
				huobi.Id = i
				bian.Id = i
				okex.Id = i
				//model.DBHd.Create(huobi)
				//model.DBHd.Create(bian)
				//model.DBHd.Create(okex)
				i++
			}
		}
	}
	return
}

func Init() {
	getAllTicker()
}
