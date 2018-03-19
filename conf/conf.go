package conf

import (
	"encoding/json"
	"os"
)

type Http struct {
	Addr     string
	TimeOut  int
	CertFile string
	KeyFile  string
	Assets   string
}

type DB struct {
	Host   string
	Port   string
	User   string
	Passwd string
	DbName string
}

type Redis struct {
	Addr        string
	MaxIdle     int
	IdleTimeOut int
	Passwd      string
	Expire      int
}

type App struct {
	AppID            string
	AppSecret        string
	SessionLifeCycle int
}

type Conf struct {
	Http  *Http
	DB    *DB
	Redis *Redis
	App   *App
}

func ReadConfig(confpath string) (*Conf, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	config := Conf{}
	err := decoder.Decode(&config)

	return &config, err
}
