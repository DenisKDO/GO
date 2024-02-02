package _func

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	var jsonResponse []byte
	var err error
	for _, item := range PrepareTeam() {
		if item.UniformNumber == params["uniformnumber"] && item.Position == params["position"] {
			jsonResponse, err = json.Marshal(item)
		}
	}
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
