package main

import (
	"flag"
	"fmt"
	"runtime"
	"simple_cab/cab_info_manager"
	"simple_cab/config"
	"simple_cab/data_model/db"
	"simple_cab/http_server"
	log "simple_cab/logging"
)

var (
	Version   string
	Build     string
	BuildTime string
)

func main() {
	log.Info("Hello Data Republic")
	config.GetInstance()
	ver := flag.Bool("version", false, "Display version")
	if *ver {
		fmt.Printf("Version: %s, BuildTime: %s\n", Version, BuildTime)
		return
	}
	log.Infof("%+v", *config.GetInstance())
	db.InitConn(config.GetInstance().MySQL.DBName)
	runtime.GOMAXPROCS(runtime.NumCPU())
	cab_info_manager.GetManagerInstance()
	http_server.ListenAndServe()
}
