package main

import (
	"AllMarket/model"
	"AllMarket/market"
)



func init() {

	model.Init()

	market.Init()


}

func main() {

	market.GetHuobiMarket()

	return
}
