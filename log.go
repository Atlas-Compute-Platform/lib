package lib

import (
	"fmt"
	"io"
)

const FMT_LOG = "%s(%s)!\n"

func LogError(w io.Writer, sym string, err error) {
	fmt.Fprintf(w, FMT_LOG, sym, err)
}
