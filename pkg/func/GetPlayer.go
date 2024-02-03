package _func

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// function that get information about members of the team
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	//content type, file json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//request variables
	params := mux.Vars(r)

	//creating jsonResponse and error to write
	var jsonResponse []byte
	var err error
	//finding player
	for _, item := range PrepareTeam() {
		if item.UniformNumber == params["uniformnumber"] && item.Position == params["position"] && item.Country == params["country"] {
			jsonResponse, err = json.Marshal(item)
		}
	}
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
