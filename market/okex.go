package market

import (
	"AllMarket/webclient"
	"fmt"
	"AllMarket/model"
	//"regexp"
	"github.com/wonderivan/logger"
	"encoding/json"
	"strconv"
)

var (
	//regx = regexp.MustCompile(`market.([a-zA-Z]+).detail`)

	//OkexEndpoint = "wss://real.okex.com:10442/ws/v3"
)

func swapOkexMarket(h, tmph *model.Huobi) {
	h.Open = tmph.Open
	h.Close = tmph.Close
	h.Amount = tmph.Amount
	h.High = tmph.High
	h.Low = tmph.Low
	h.Count = tmph.Count
	h.Vol = tmph.Vol
	h.Rose, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", (h.Open-h.Close)/h.Open*100+0.005), 64)
}

func GetOkexMarket() {

	srcMarket,ok:=model.SrcMarketMap["okex"]
	if !ok{
		logger.Error("can not find okex")
		return
	}

	market, err := webclient.NewMarket(srcMarket)
	if err != nil {
		panic(err)
	}
	//"{\"op\": \"subscribe\", \"args\": [\"index/ticker:BTC-USD\",\"index/ticker:ETH-USD\"]}"
	for _, ticker := range TickerList {
		topic := fmt.Sprintf("market.%s.detail", ticker)

		market.Subscribe(ticker,topic, func(topic string, resp *JSON) {

			modelHuoBi ,ok:= model.HuobiMap[topic]
			if !ok{
				logger.Error("can not find ticke:",topic)
				return
			}

			strTick := resp.Get("tick")

			tick, err := strTick.Encode()
			if err != nil {
				logger.Error(err.Error())
				return
			}
			var tmpHuobi model.Huobi
			err = json.Unmarshal(tick, &tmpHuobi)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			swapOkexMarket(&modelHuoBi,&tmpHuobi)
			//fmt.Println("----------------->modelHuoBi:",modelHuoBi)
			modelHuoBi.Update()
		})
	}

	market.Loop()
}
