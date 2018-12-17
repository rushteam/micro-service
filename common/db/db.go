package db

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DefaultGroup 默认分组
const DefaultGroup = "default"

//MasterType 主库
const MasterType = "master"

//SalveType 从库
const SalveType = "salve"

var defaultPoolGroup PoolGroup

//Configs ..
type Configs map[string]map[string]struct {
	DbType string `yaml:"db_type"`
	DSN    string `yaml:"DSN"`
}

//Engine ..
type Engine struct {
	DB *gorm.DB
}

//Pool 数据库池
type Pool map[string][]*Engine

//GetEngine 获取一个数据池
func (p Pool) GetEngine(name string) (*Engine, error) {
	if name == "" {
		name = MasterType
	}
	var engine *Engine
	engineList, ok := p[name]
	if !ok || len(engineList) < 1 {
		return engine, errors.New("not found engine")
	}
	if name == MasterType {
		engine = engineList[0]
	} else {
		//todo 随机算法取一个,目前仅取第一个
		engine = engineList[0]
	}
	return engine, nil
}

//PoolGroup 数据池组
type PoolGroup map[string]Pool

//GetPool 获取一个数据池
func (g PoolGroup) GetPool(groupName string) (*Pool, error) {
	if groupName == "" {
		groupName = DefaultGroup
	}
	p, ok := g[groupName]
	if !ok {
		return &p, errors.New("not found pool in group")
	}
	return &p, nil
}

//Load ..
func (g PoolGroup) Load(conf Configs) error {
	for groupName, pg := range conf {
		if _, ok := g[groupName]; !ok {
			pool := make(Pool, 0)
			g[groupName] = pool
		}
		for poolName, p := range pg {
			if _, ok := g[groupName][poolName]; !ok {
				e := Connect(p.DbType, p.DSN)
				g[groupName][poolName] = append(g[groupName][poolName], &Engine{e})
			}
		}
	}
	return nil
}

/*
db_config.yml
	db_configs:
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

//Connect 初始化db数据
func Connect(dbType, DSN string) *gorm.DB {
	if dbType == "" {
		dbType = "mysql"
	}
	var err error
	db, err := gorm.Open(dbType, DSN)
	// defer db.Close()
	if err != nil {
		log.Printf("[db] connect failed (%s) %s\r\n", DSN, err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Hour)
	return db
}

//Init ..
func Init(conf Configs) {
	defaultPoolGroup = PoolGroup{}
	defaultPoolGroup.Load(conf)
}

//Get ..
func Get(groupName, poolName string) (*Engine, error) {
	var engine *Engine
	pool, err := defaultPoolGroup.GetPool(groupName)
	if err != nil {
		return nil, err
	}
	engine, err = pool.GetEngine(poolName)
	return engine, err
}

//Default ..
func Default() (*Pool, error) {
	pool, err := defaultPoolGroup.GetPool(DefaultGroup)
	return pool, err
}

type Model struct {
}
