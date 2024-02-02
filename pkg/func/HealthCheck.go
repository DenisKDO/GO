package _func

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//response
	fmt.Fprintf(w, "API by Kim Denis Olegovich\n Mini API related to Japan National Voleyball team")
}
