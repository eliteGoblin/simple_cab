package action

import (
	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"simple_cab/cab_info_manager"
	"simple_cab/config"
	log "simple_cab/logging"
	"simple_cab/toolkit"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	maxBodyLength = 1000
)

type CabCountInfoQuery struct {
	NoCache    bool     `json:"no_cache"`
	Medallions []string `json:"medallions"`
}

type CabCountInfoResponse struct {
	MedallionCount map[string]int `json:"medallions"`
}

func NewCabCountInfoResponse() *CabCountInfoResponse {
	return &CabCountInfoResponse{
		MedallionCount: make(map[string]int),
	}
}

func GetCabsPickupCountInfo(res http.ResponseWriter, req *http.Request) {
	// in case coming request is too large
	req.Body = http.MaxBytesReader(res, req.Body, maxBodyLength)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error(err)
		return
	}
	vars := mux.Vars(req)
	dateStr := vars["date"]
	var cabCountQuery CabCountInfoQuery
	err = json.Unmarshal(body, &cabCountQuery)
	if err != nil || dateStr == "" ||
		len(cabCountQuery.Medallions) > config.GetInstance().SimpleCab.MaxMedallionCountPerRequest {
		toolkit.RenderError(res, "request count over limit")
		return
	}
	response := NewCabCountInfoResponse()
	if cabCountQuery.NoCache == false {
		// use cache
		for _, medallion := range cabCountQuery.Medallions {
			if count, errTripCount := cab_info_manager.GetManagerInstance().GetCabTripCount(
				medallion, dateStr); errTripCount != nil {
				log.Errorf("%s err: %s", medallion, errTripCount)
			} else {
				response.MedallionCount[medallion] += count
			}
		}
	} else {
		medallionCount, err := cab_info_manager.GetManagerInstance().GetCabTripCountFromDB(
			cabCountQuery.Medallions, dateStr)
		if err != nil {
			toolkit.RenderError(res, err.Error())
			return
		}
		response.MedallionCount = medallionCount
	}

	toolkit.RenderResponse(res, response)
}
