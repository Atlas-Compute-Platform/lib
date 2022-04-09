package lib

/*
	Atlas Shared Library
	Thijs Haker
*/

import (
	"fmt"
	"io"
	"net/http"
)

const (
	VERS    = "0.0"
	PORT    = ":8800"
	FMT_LOG = "%s(%s)!\n"
)

func SetCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func ApiPing(w http.ResponseWriter, r *http.Request) {
	SetCors(&w)
	fmt.Fprint(w, "pong")
}

func LogError(w io.Writer, sym string, err error) {
	fmt.Fprintf(w, FMT_LOG, sym, err)
}
