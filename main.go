package main

import (
	"AllMarket/model"
	"fmt"
	"AllMarket/market"
)



func init() {

	model.Init()

	market.Init()
}

func main() {


	for key,v := range model.StandardBiMap{
		fmt.Println(key,"  ",v)
	}

	market.GetHuobiMarket()

	return
}
