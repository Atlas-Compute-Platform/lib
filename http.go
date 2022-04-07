package lib

import (
	"fmt"
	"net/http"
)

func ApiPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "pong")
}
