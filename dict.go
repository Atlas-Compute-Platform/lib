package lib

import (
	"encoding/json"
	"fmt"
	"io"
)

type Dict map[string]string

func LoadDict(r io.Reader) (Dict, error) {
	var (
		buf []byte
		err error
		dic = make(Dict)
	)

	if buf, err = io.ReadAll(r); err != nil {
		return nil, err
	}
	if json.Unmarshal(buf, &dic) != nil {
		return nil, err
	}

	return dic, nil
}

func StoreDict(dic Dict, w io.Writer) error {
	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(&dic); err != nil {
		return err
	}
	if _, err = fmt.Fprint(w, string(buf)); err != nil {
		return err
	}

	return nil
}
