package _func

import (
	"encoding/json"
	"net/http"
)

type shortInfo struct {
	Fn  string `json:"first-name"`
	Sn  string `json:"second-name"`
	Pos string `json:"pos"`
	Un  string `json:"uniform-number"`
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	team := PrepareTeam()
	var shortInfos []shortInfo
	for _, item := range team {
		var shortInf shortInfo
		shortInf.Fn = item.FirstName
		shortInf.Sn = item.SecondName
		shortInf.Pos = item.Position
		shortInf.Un = item.UniformNumber

		shortInfos = append(shortInfos, shortInf)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(shortInfos)
	if err != nil {
		return
	}

	w.Write(jsonResponse)

}
