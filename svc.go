package lib

import (
	"encoding/json"
	"fmt"
	"io"
)

type SvcForm struct {
	Name string
	Port string
}

func LoadSvcForm(r io.Reader) (SvcForm, error) {
	var (
		buf []byte
		err error
		svc = new(SvcForm)
	)

	if buf, err = io.ReadAll(r); err != nil {
		return SvcForm{}, err
	}
	if json.Unmarshal(buf, svc) != nil {
		return SvcForm{}, err
	}

	return *svc, nil
}

func StoreSvcForm(svc SvcForm, w io.Writer) error {
	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(&svc); err != nil {
		return err
	}
	if _, err = fmt.Fprint(w, string(buf)); err != nil {
		return err
	}

	return nil
}
