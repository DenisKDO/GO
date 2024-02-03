package _func

import (
	"fmt"
	"net/http"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All national teams that my API have:\n-Russia\n-Japan")
}
