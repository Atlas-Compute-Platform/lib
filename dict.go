package lib

import (
	"encoding/json"
	"fmt"
	"io"
)

type Dict map[string]string

func ImportDict(buf []byte) (Dict, error) {
	var (
		err error
		dic = make(Dict)
	)

	if json.Unmarshal(buf, &dic) != nil {
		return nil, err
	}

	return dic, nil
}

func ExportDict(dic Dict) ([]byte, error) {
	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(&dic); err != nil {
		return nil, err
	}

	return buf, nil
}

func ReceiveDict(r io.Reader) (Dict, error) {
	var (
		buf []byte
		err error
		dic Dict
	)

	if buf, err = io.ReadAll(r); err != nil {
		return nil, err
	}
	if dic, err = ImportDict(buf); err != nil {
		return nil, err
	}

	return dic, nil
}

func SendDict(dic Dict, w io.Writer) error {
	var (
		buf []byte
		err error
	)

	if buf, err = ExportDict(dic); err != nil {
		return err
	}
	if _, err = fmt.Fprint(w, string(buf)); err != nil {
		return err
	}

	return nil
}
