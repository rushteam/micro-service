package db

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Default 默认分组
const Default = "default"

//Master 主库
const Master = "master"

//Salve 从库
const Salve = "salve"

//Db ..
type Db struct {
	Db        *gorm.DB
	SourceDSN string `yaml:",flow"`
}

//PoolItem ..
type PoolItem struct {
	DbType string          `yaml:"db_type"`
	Items  map[string][]Db `yaml:",flow"`
}

//Pool ..
type Pool struct {
	// Db map[string]map[string][]*gorm.DB
	// groups map[string]PoolItem
	Db map[string]PoolItem `yaml:"db_config"`
}

//GetDb ..
func (p Pool) GetDb(groupName string, poolName string) (*gorm.DB, error) {
	if _, ok := p.Db[Default]; !ok {
		return nil, errors.New("db group is null")
	}
	if _, ok := p.Db[Default].Items[Master]; !ok {
		return nil, errors.New("db master config is null")
	}
	if len(p.Db[Default].Items[Master]) < 1 {
		return nil, errors.New("db master config is null")
	}
	if poolName == Master {
		return p.Db[Default].Items[poolName][0].Db, nil
	}
	if _, ok := p.Db[Default].Items[poolName]; !ok {
		return p.Db[Default].Items[poolName][0].Db, nil
	}
	//现在只能有一个从库，后期需要随机一个，这里先不扩展了
	return p.Db[Default].Items[groupName][0].Db, nil
}

/*
db_config.yml
	db_config:
		default:
			db_type:"mysql"
			master:
				- ""
			slave:
				- ""
				- ""
dbInstance map[string]map[string][]*gorm.DB

GetDb(group string,pool string)
*/

//CreateDb 初始化db数据
func CreateDb(dbType, sourceDSN string) *gorm.DB {
	if dbType == "" {
		dbType = "mysql"
	}
	var err error
	db, err := gorm.Open(dbType, sourceDSN)
	// defer db.Close()
	if err != nil {
		log.Printf("[db] connect failed (%s) %s\r\n", sourceDSN, err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Hour)
	return db
}

//InitDefaultDb ..
func (p *Pool) InitDefaultDb(dbType, sourceDSN string, logMode bool) *gorm.DB {
	groupName := Default
	poolName := Master
	if _, ok := p.Db[groupName]; !ok {
		p.Db = make(map[string]PoolItem)
		pooItem := PoolItem{}
		pooItem.Items = make(map[string][]Db)
		p.Db[groupName] = pooItem
	}
	if _, ok := p.Db[groupName].Items[poolName]; !ok {
		p.Db[groupName].Items[poolName] = make([]Db, 0)
	}
	db := Db{SourceDSN: sourceDSN}
	db.Db = CreateDb(dbType, sourceDSN)
	db.Db.LogMode(logMode)
	p.Db[groupName].Items[poolName] = append(p.Db[groupName].Items[poolName], db)
	return db.Db
}

//InitDb ..
func InitDb(dbType, sourceDSN string, logMode bool) *Pool {
	var pool = &Pool{}
	pool.InitDefaultDb(dbType, sourceDSN, logMode)
	return pool
}

//InitDbConfig ..
func InitDbConfig(configFile string) *Pool {
	var pool = &Pool{}
	for _, poolItem := range pool.Db {
		for _, dbs := range poolItem.Items {
			for _, db := range dbs {
				db.Db = CreateDb(poolItem.DbType, db.SourceDSN)
			}
		}
	}
	return pool
}
