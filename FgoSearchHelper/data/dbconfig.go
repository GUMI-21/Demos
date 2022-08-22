package data

import "time"

type Config struct {
	Addr            string  `toml:"addr" json:"addr"`
	User            string  `toml:"user" json:"user"`
	Pwd             string  `toml:"pwd" json:"pwd"`
	Db              string  `toml:"db" json:"db"`
	Options         string  `toml:"options" json:"options"`
	MaxOpenConns    int     `toml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns    int     `toml:"max_idle_conns" json:"max_idle_conns"`
	MaxConnLifeTime float64 `toml:"max_conn_life_time" json:"max_conn_life_time"`
}

var ConfigMap = map[string]Config{
	"mysql": DbLocalMysql,
}

var DbLocalMysql = Config{
	Addr:            "127.0.0.1:3306",
	User:            "root",
	Pwd:             "luoying123..",
	Db:              "test",
	MaxOpenConns:    5,
	MaxIdleConns:    3,
	MaxConnLifeTime: time.Hour.Hours(),
}
