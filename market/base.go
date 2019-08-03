package market

import (
	"github.com/bitly/go-simplejson"
	"strings"
	"fmt"
	"AllMarket/model"
	"github.com/wonderivan/logger"
)

type JSON = simplejson.Json

var (//
	//wss://www.hbdm.com/ws
	HuobiEndpoint = "wss://api.huobi.pro/ws"

	OkexEndpoint = "wss://real.okex.com:10442/ws/v3"

	HuobiTickerList []string
	OkexTickerList  []string
)

func getAllTicker() {
	for value, _ := range model.StandardBiMap {
		for v, _ := range model.DealBiMap {
			if v != value {
				lowerValue := strings.ToLower(value)
				lowerV := strings.ToLower(v)
				ticker := fmt.Sprintf("%s%s", lowerV, lowerValue)
				HuobiTickerList = append(HuobiTickerList, ticker)
				ticker = fmt.Sprintf("%s-%s",v,value)
				OkexTickerList = append(OkexTickerList,ticker)
			}
		}
	}
	logger.Debug("getAllTicker")
	return
}

func Init() {
	getAllTicker()
}
