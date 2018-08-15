package http_server

import (
	"github.com/gorilla/mux"
	"net/http"
	"simple_cab/action"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", action.Index).Methods("GET")
	router.HandleFunc("/v1/trip_info/{date}/count", action.GetCabsPickupCountInfo).Methods("POST")
	router.HandleFunc("/v1/trip_info/update_cache", action.UpdateCache).Methods("PUT")
	return router
}
