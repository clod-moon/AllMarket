package model

import (
	"github.com/bitly/go-simplejson"
	"time"
	"github.com/jinzhu/gorm"
)

type JSON = simplejson.Json

var (
	db            *gorm.DB
	StandardBiMap = make(map[int]string)
	DealBiMap     = make(map[int]string)
	HuobiMap      = make(map[string]Huobi)
	BianMap       = make(map[string]Bian)
	OkexMap       = make(map[string]Okex)
)

type StandardBi struct {
	Id         int       `gorm:"primary_key;type:int(11);AUTO_INCREMENT`
	Name       string    `gorm:"type:varchar(30);not null"`
	CreateTime time.Time `gorm:"type:datetime;not null;"`
	UpdateTime time.Time `gorm:"type:datetime;not null;"`
}

func GetAllStandardBi() {
	var list []StandardBi
	db.Find(&list)
	for _, v := range list {
		StandardBiMap[v.Id] = v.Name
	}
}

func (s *StandardBi) GetName() string {
	db.Find(s, "id= ?", s.Id)
	return s.Name
}

type DealBi struct {
	Id         int       `gorm:"primary_key;type:int(11);AUTO_INCREMENT`
	Name       string    `gorm:"type:varchar(30);not null"`
	CreateTime time.Time `gorm:"type:datetime;not null;"`
	UpdateTime time.Time `gorm:"type:datetime;not null;"`
}

func GetAllDealBi() {
	var list []DealBi
	db.Find(&list)
	for _, v := range list {
		DealBiMap[v.Id] = v.Name
	}
}

func (d *DealBi) GetName() string {
	db.Find(d, " id = ? ", d.Id)
	return d.Name
}
