package model

import "time"

//

type Okex struct {
	Id           int       `gorm:"primary_key;type:int(11);AUTO_INCREMENT`
	Ticker       string    `json:"instrument_id"`
	DealBiId     int       `gorm:"type:int;not null"`
	StandardBiId int       `gorm:"type:int;not null"`
	Amount       float64   `json:"amount" gorm:"type:float;not null"`
	Open         float64   `json:"open_24h" gorm:"type:float;not null"`
	Close        float64   `json:"last" gorm:"type:float;not null"`
	Rose         float64   `json:"rose" gorm:"type:float;not null"`
	High         float64   `json:"high" gorm:"type:float;not null"`
	Count        float64   `json:"count" gorm:"type:float;not null"`
	Low          float64   `json:"low_24h" gorm:"type:float;not null"`
	Vol          float64   `json:"vol" gorm:"type:float;not null"`
	CreateTime   time.Time `gorm:"type:datetime;not null;"`
	UpdateTime   time.Time `gorm:"type:datetime;not null;"`
}

//func GetOkex()  {
//	var list []Okex
//	DBHd.Find(&list)
//
//	for _,v := range list{
//		ticker := DealBiMap[v.DealBiId]+StandardBiMap[v.StandardBiId]
//		OkexMap[ticker] = v
//	}
//}

func (h *Okex) Update() {
	DBHd.Update(h)
}
