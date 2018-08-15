package config

import (
	"github.com/koding/multiconfig"
	"sync"
)

type (
	MySQLCnf struct {
		Host         string
		Port         int `default:"3306"`
		User         string
		Password     string
		DBName       string `toml:"db_name"`
		MaxIdleConns int    `toml:"max_idle_conns"`
		MaxOpenConns int    `toml:"max_open_conns"`
		EnableLog    bool   `toml:"enable_log" default:"false"`
	}
	SimpleCabCnf struct {
		Host                        string `toml:"host"`
		Port                        int    `toml:"port"`
		MaxMedallionCountPerRequest int    `toml:"max_count_per_request"`
	}
)

type AllConfig struct {
	MySQL     MySQLCnf     `toml:"mysql"`
	SimpleCab SimpleCabCnf `toml:"simple_cab"`
}

var instanceConfig *AllConfig
var once sync.Once

func GetInstance() *AllConfig {
	once.Do(func() {
		m := multiconfig.NewWithPath(getConfigFilePath())
		cnf := new(AllConfig)
		m.MustLoad(cnf)
		instanceConfig = cnf
	})
	return instanceConfig
}
