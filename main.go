package main

import (
	"AllMarket/model"
	"AllMarket/market"
	"sync"
)

var(
	wg sync.WaitGroup
)

func init() {

	model.Init()

	market.Init()


}

func main() {

	//wg.Add(1)
	//go market.GetHuobiMarket()

	wg.Add(1)
	go market.GetOkexMarket()

	wg.Wait()
	return
}
