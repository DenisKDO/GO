package _func

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

//adding struct with short information about players
type shortInfo struct {
	Fn  string `json:"first-name"`
	Sn  string `json:"second-name"`
	Pos string `json:"pos"`
	Un  string `json:"uniform-number"`
}

//function that get national team information by their country
func GetTeam(w http.ResponseWriter, r *http.Request) {
	//creating var with info about team
	team := PrepareTeam()

	//creating slice with short information to convert it into json
	var shortInfos []

	//var with parameters
	params := mux.Vars(r)

	//adding short information to our slice
	for _, item := range team {
		var shortInf shortInfo
		shortInf.Fn = item.FirstName
		shortInf.Sn = item.SecondName
		shortInf.Pos = item.Position
		shortInf.Un = item.UniformNumber

		//finding neccecary team member by their country
		if item.Country == params["country"] {
			shortInfos = append(shortInfos, shortInf)
		}
	}

	//convert our slice of struct into json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(shortInfos)
	if err != nil {
		return
	}

	w.Write(jsonResponse)

}
