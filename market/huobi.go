package market

import (
	"AllMarket/webclient"
	"fmt"
	"AllMarket/model"
	"regexp"
	"github.com/wonderivan/logger"
	"encoding/json"
	"strconv"
)

var (
	regx = regexp.MustCompile(`market.([a-zA-Z]+).detail`)
)

func swapHuobiMarket(h, tmph *model.Huobi) {
	h.Open = tmph.Open
	h.Close = tmph.Close
	h.Amount = tmph.Amount
	h.High = tmph.High
	h.Low = tmph.Low
	h.Count = tmph.Count
	h.Vol = tmph.Vol
	h.Rose, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", (h.Open-h.Close)/h.Open*100+0.005), 64)
}

func GetHuobiMarket() {

	srcMarket,ok:=model.SrcMarketMap["huobi"]
	if !ok{
		logger.Error("can not find huobi")
		return
	}

	market, err := webclient.NewMarket(srcMarket)
	if err != nil {
		panic(err)
	}

	for _, ticker := range TickerList {
		topic := fmt.Sprintf("market.%s.detail", ticker)
		//fmt.Println("topic:",topic)
		//var huobi model.Huobi
		//model.HuobiMap[topic] = huobi
		market.Subscribe(ticker,topic, func(topic string, resp *JSON) {

			modelHuoBi ,ok:= model.HuobiMap[ticker]
			if !ok{
				logger.Error("can not find ticke:",ticker)
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
			swapHuobiMarket(&modelHuoBi,&tmpHuobi)
			//fmt.Println("----------------->modelHuoBi:",modelHuoBi)
			modelHuoBi.Update()
		})
	}

	market.Loop()
}
