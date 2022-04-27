package lib

/*
	Atlas Shared Library
	Thijs Haker
*/

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	PORT    = ":8800"
	FMT_LOG = "%s => %s\n"
	CONFIG  = "config.json"
)

var (
	SvcName = "Atlas Microservice"
	SvcVers = "0.0"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "%s %s\n", SvcName, SvcVers)
	flag.PrintDefaults()
	os.Exit(1)
}

func SetCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func ApiPing(w http.ResponseWriter, r *http.Request) {
	SetCors(&w)
	fmt.Fprint(w, "pong")
}

func LogFatal(w io.Writer, sym string, err error, ec int) {
	fmt.Fprintf(w, FMT_LOG, sym, err)
	os.Exit(ec)
}

func LogError(w io.Writer, sym string, err error) {
	fmt.Fprintf(w, FMT_LOG, sym, err)
}

func LogMessage(w io.Writer, msg string) {
	fmt.Fprintf(w, "%s", msg)
}
