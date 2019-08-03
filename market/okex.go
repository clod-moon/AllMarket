package market

import (
	"AllMarket/webclient"
	"AllMarket/model"
	//"regexp"
	"github.com/wonderivan/logger"
	"encoding/json"
	"strconv"
	"time"
	"fmt"
)

var (
//regx = regexp.MustCompile(`market.([a-zA-Z]+).detail`)

//OkexEndpoint = "wss://real.okex.com:10442/ws/v3"
)

type Okex struct {
	Open      string    `json:"open_24h"`
	Close     string    `json:"last" `
	Low       string    `json:"low_24h"`
	Timestamp time.Time `json:"timestamp"`
}

func swapOkexMarket(o *model.Okex, tmpO *Okex) {
	o.Open, _ = strconv.ParseFloat(tmpO.Open, 64)
	o.Close, _ = strconv.ParseFloat(tmpO.Close, 64)
	o.Low, _ = strconv.ParseFloat(tmpO.Low, 64)
	o.Rose, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", (o.Open-o.Close)/o.Open*100+0.005), 64)
}

func GetOkexMarket() {

	srcMarket, ok := model.SrcMarketMap["okex"]
	if !ok {
		logger.Error("can not find okex")
		return
	}

	market, err := webclient.NewMarket(srcMarket)
	if err != nil {
		panic(err)
	}
	//"{\"op\": \"subscribe\", \"args\": [\"index/ticker:BTC-USD\",\"index/ticker:ETH-USD\"]}"
	for _, ticker := range OkexTickerList {

		market.Subscribe(ticker, func(topic string, resp *JSON) {

			modelOkex, ok := model.OkexMap[ticker]
			if !ok {
				logger.Error("can not find ticke:", ticker)
				return
			}

			strTick := resp.Get("tick")

			tick, err := strTick.Encode()
			if err != nil {
				logger.Error(err.Error())
				return
			}
			var tmpOkex Okex
			err = json.Unmarshal(tick, &tmpOkex)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			swapOkexMarket(&modelOkex, &tmpOkex)
			//fmt.Println("----------------->modelHuoBi:",modelHuoBi)
			modelOkex.Update()
		})
	}

	market.Loop()
}
