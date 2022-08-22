package data

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mu      sync.Mutex
	x       = make(map[string]*sql.DB)
	ErrNoDb = errors.New("no db")
)

//mysql链接池配置
func GetConn(name string) *sql.DB {
	mu.Lock()
	defer mu.Unlock()

	conn, err := getDB(name)
	if err != nil && conn != nil {
		return conn
	}
	_ = addConn(name, ConfigMap[name])
	conn, _ = getDB(name)
	return conn
}

func getDB(name string) (*sql.DB, error) {
	db, ok := x[name]
	if ok {
		return db, nil
	}
	return nil, ErrNoDb
}

func addConn(name string, c Config) error {
	db, err := sql.Open("mysql", c.Dsn())
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)

	if c.MaxConnLifeTime <= 0 {
		c.MaxConnLifeTime = 600 // Default 10 minutes
	}
	db.SetConnMaxLifetime(time.Duration(c.MaxConnLifeTime) * time.Second)

	_, ok := x[name]
	if ok {
		return errors.New("db exists")
	}
	x[name] = db
	return nil
}

//格式化配置
func (c Config) Dsn() string {
	if c.Options == "" {
		format := "%s:%s@tcp(%s)/%s?charset=utf8mb4&timeout=5s"
		return fmt.Sprintf(format, c.User, c.Pwd, c.Addr, c.Db)
	} else {
		format := "%s:%s@tcp(%s)/%s?charset=utf8mb4&timeout=5s&%s"
		return fmt.Sprintf(format, c.User, c.Pwd, c.Addr, c.Db, c.Options)
	}
}
