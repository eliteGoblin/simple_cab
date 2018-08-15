package http_server

import (
	"fmt"
	"net/http"
	"simple_cab/config"
	log "simple_cab/logging"
)

func ListenAndServe() {
	router := NewRouter()
	simpleCabConfig := config.GetInstance().SimpleCab
	addr := fmt.Sprintf("%s:%d", simpleCabConfig.Host, simpleCabConfig.Port)
	log.Infof("cache built complete, simple cab service will listen on addr %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
