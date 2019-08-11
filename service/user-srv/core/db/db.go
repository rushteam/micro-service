package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/micro/go-micro/util/log"
	"upper.io/db.v3/mysql"
)

var (
	db     *sql.DB
	m      sync.RWMutex
	inited bool
)

// Init 初始化数据库
func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Logf(err.Error())
		return
	}

	settings, _ := mysql.ParseURL("root:dream@tcp(127.0.0.1:3306)/rushteam?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s")
	sess, err := mysql.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer sess.Close()

	inited = true
}

// GetSession 获取db
func GetSession() *sql.DB {
	return db
}
