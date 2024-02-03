package _func

import (
	"fmt"
	"net/http"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "All avialable teams: Russia, Japan")
}
