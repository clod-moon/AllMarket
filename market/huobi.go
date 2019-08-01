package market

import (
	"AllMarket/webclient"
	"fmt"
	"AllMarket/model"
	"regexp"
)

var (
	regx = regexp.MustCompile(`market.([a-zA-Z]+).detail`)
)


func GetHuobiMarket(){
	market, err := webclient.NewMarket()
	if err != nil {
		panic(err)
	}
	for _,ticker := range TickerList{
		topic := fmt.Sprintf("market.%s.detail",ticker)
		//fmt.Println("topic:",topic)
		var huobi model.Huobi
		model.HuobiMap[topic] = huobi
		market.Subscribe(topic, func(topic string, json *JSON) {
			close,_ := json.Get("tick").Get("close").Float64()
			open,_ := json.Get("tick").Get("open").Float64()
			tt := regx.FindStringSubmatch(topic)
			//fmt.Println("----->",tt)
			fmt.Println(tt[1],"open：",open," close：",close, "涨幅：",(close-open)/open*100)
		})
	}

	market.Loop()
}
