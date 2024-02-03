package main

import (
	_func "github.com/DenisKDO/TSIS1/pkg/func"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//create a new router
	router := mux.NewRouter()

	//specify endpoints, handler functions and HTTP methods
	router.HandleFunc("/health-check", _func.HealthCheck).Methods("GET")
	router.HandleFunc("/team/{country}", _func.GetTeam).Methods("GET")
	router.HandleFunc("/team/{country}/{position}/{uniformnumber}", _func.GetPlayer).Methods("GET")
	http.Handle("/", router)
	//start and listen to request
	http.ListenAndServe(":8080", router)
}
