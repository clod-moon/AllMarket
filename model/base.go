package model

import (
	"github.com/bitly/go-simplejson"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"log"
)

type JSON = simplejson.Json



var (

	username string = "root"
	password string = "root"
	dbName   string = "market"
	host     string = "192.168.1.216"
	port     int    = 3306

	DBHd            *gorm.DB
	StandardBiMap = make(map[string]int)
	DealBiMap     = make(map[string]int)
	HuobiMap      = make(map[string]Huobi)
	BianMap       = make(map[string]Bian)
	OkexMap       = make(map[string]Okex)
)


func Init(){

	mysqlstr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	DB, err := gorm.Open("mysql", mysqlstr)
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}
	DBHd = DB

	DBHd.SingularTable(true)
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "coin_" + defaultTableName
	}

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

	GetAllStandardBi()

	GetAllDealBi()
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
	}
}

func (d *DealBi) GetName() string {
	DBHd.Find(d, " id = ? ", d.Id)
	return d.Name
}
