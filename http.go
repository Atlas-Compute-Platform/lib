package lib

import (
	"fmt"
	"net/http"
)

func SetCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}

func ApiPing(w http.ResponseWriter, r *http.Request) {
	w = SetCors(w)
	fmt.Fprint(w, "pong")
}
