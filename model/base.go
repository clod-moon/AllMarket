package model

import (
	"github.com/bitly/go-simplejson"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"log"
	"github.com/wonderivan/logger"
)

type JSON = simplejson.Json

var (
	username string = "root"
	password string = "root"
	dbName   string = "market"
	host     string = "192.168.0.103"
	port     int    = 3306

	DBHd          *gorm.DB
	StandardBiMap = make(map[string]int)
	DealBiMap     = make(map[string]int)
	HuobiMap      = make(map[string]Huobi)
	BianMap       = make(map[string]Bian)
	OkexMap       = make(map[string]Okex)
	SMap          = make(map[int]string)
	DMap          = make(map[int]string)
	SrcMarketMap  = make(map[string]SrcMarket)
)

func Init() {

	mysqlstr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	DB, err := gorm.Open("mysql", mysqlstr)
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}
	DBHd = DB

	DBHd.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "coin_" + defaultTableName
	}

	CheckTable()

	GetAllStandardBi()

	GetAllDealBi()

	GetAllSrcMarket()

	GetHuoBi()

	logger.Debug("model Init")
	//FillSrcMarket()
	//FillHuobi()
}

func CheckTable(){

	if !DBHd.HasTable(&StandardBi{}) {
		err := DBHd.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&StandardBi{}).Error
		if err != nil {
			panic(err)
		}
	}

	if !DBHd.HasTable(&DealBi{}) {
		err := DBHd.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&DealBi{}).Error
		if err != nil {
			panic(err)
		}
	}

	if !DBHd.HasTable(&SrcMarket{}) {
		err := DBHd.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&SrcMarket{}).Error
		if err != nil {
			panic(err)
		}
	}

	if !DBHd.HasTable(&Bian{}) {
		err := DBHd.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Bian{}).Error
		if err != nil {
			panic(err)
		}
	}

	if !DBHd.HasTable(&Huobi{}) {
		err := DBHd.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Huobi{}).Error
		if err != nil {
			panic(err)
		}
	}

	if !DBHd.HasTable(&Okex{}) {
		err := DBHd.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Okex{}).Error
		if err != nil {
			panic(err)
		}
	}
}

type StandardBi struct {
	Id         int       `gorm:"primary_key;type:int(11);AUTO_INCREMENT`
	Name       string    `gorm:"type:varchar(30);not null"`
	CreateTime time.Time `gorm:"type:datetime;not null;"`
	UpdateTime time.Time `gorm:"type:datetime;not null;"`
}

func GetAllStandardBi() {
	var list []StandardBi
	DBHd.Find(&list)
	for _, v := range list {
		StandardBiMap[v.Name] = v.Id
		SMap[v.Id] = v.Name
	}
}

func (s *StandardBi) GetName() string {
	DBHd.Find(s, "id= ?", s.Id)
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
	DBHd.Find(&list)
	for _, v := range list {
		DealBiMap[v.Name] = v.Id
		DMap[v.Id] = v.Name
	}
}

func (d *DealBi) GetName() string {
	DBHd.Find(d, " id = ? ", d.Id)
	return d.Name
}

type SrcMarket struct {
	Id         int       `gorm:"primary_key;type:int(11);AUTO_INCREMENT`
	Name       string    `gorm:"type:varchar(30);not null"`
	WsUrl      string    `gorm:"type:varchar(255);not null"`
	Template   string    `gorm:"type:varchar(255);not null"`
	IsUnzip    bool      `gorm:"type:bool;not null"`
	CreateTime time.Time `gorm:"type:datetime;not null;"`
	UpdateTime time.Time `gorm:"type:datetime;not null;"`
}

func GetAllSrcMarket() {
	var list []SrcMarket
	//SrcMarketMap
	DBHd.Find(&list)
	for _,v:=range list{
		SrcMarketMap[v.Name] = v
	}
}

func FillHuobi() {
	var list []DealBi
	DBHd.Order("id").Find(&list)

	i := 1
	for k1, _ := range SMap {
		for _, k2 := range list {
			if k1 == k2.Id {
				continue
			}
			fmt.Println("k2:", k2)
			var huobi Huobi
			var b Bian
			var o Okex
			huobi.Id = i
			b.Id = i
			o.Id = i

			huobi.StandardBiId = k1
			huobi.DealBiId = k2.Id
			b.StandardBiId = k1
			b.DealBiId = k2.Id
			o.StandardBiId = k1
			o.DealBiId = k2.Id

			o.CreateTime = time.Now()
			o.UpdateTime = time.Now()

			b.CreateTime = time.Now()
			b.UpdateTime = time.Now()

			huobi.CreateTime = time.Now()
			huobi.UpdateTime = time.Now()

			DBHd.Create(huobi)
			DBHd.Create(b)
			DBHd.Create(o)
			i++
		}
	}
}

func FillSrcMarket() {
	var list [2]SrcMarket

	//var src SrcMarket
	list[0].Id = 1
	list[0].Name = "huobi"
	list[0].WsUrl = `wss://api.huobi.pro/ws`
	list[0].Template = `{"Sub":"market.%s.detail"}`
	list[0].IsUnzip = true
	list[0].UpdateTime = time.Now()
	list[0].CreateTime = time.Now()

	list[1].Id = 2
	list[1].Name = "okex"
	list[1].WsUrl = `wss://real.okex.com:10442/ws/v3`
	list[1].Template = `{"op": "subscribe", "args": ["index/ticker:%s"]}`
	list[1].IsUnzip = true
	list[1].UpdateTime = time.Now()
	list[1].CreateTime = time.Now()

	fmt.Println(list)

	for _, v := range list {
		DBHd.Create(v)
	}
}
