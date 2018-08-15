package action

import (
	"io/ioutil"
	"net/http"
	"simple_cab/cab_info_manager"
	log "simple_cab/logging"
)

func UpdateCache(res http.ResponseWriter, req *http.Request) {
	// in case coming request is too large
	_, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Warn(err)
		return
	}
	cab_info_manager.GetManagerInstance().UpdateCache()
}
