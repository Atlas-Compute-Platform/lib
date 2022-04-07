package lib

import (
	"fmt"
	"io"
)

func LogError(w io.Writer, sym string, err error) {
	fmt.Fprintf(w, "%s(%s)!\n", sym, err)
}
