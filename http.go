package lib

import (
	"fmt"
	"net/http"
)

const PORT = ":8800"

func SetCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}

func ApiPing(w http.ResponseWriter, r *http.Request) {
	w = SetCors(w)
	fmt.Fprint(w, "pong")
}
