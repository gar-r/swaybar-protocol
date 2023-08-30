package swaybarprotocol

import (
	"encoding/json"
	"io"
)

func Init(w io.Writer, h *Header) error {
	err := write(w, h)
	if err != nil {
		return err
	}
	err = writeStr(w, "\n")
	if err != nil {
		return err
	}
	return writeStr(w, "[")
}

func Output(w io.Writer, b *Body) error {
	err := write(w, b)
	if err != nil {
		return err
	}
	return writeStr(w, ",")
}

func write(w io.Writer, obj interface{}) error {
	enc := json.NewEncoder(w)
	return enc.Encode(obj)
}

func writeStr(w io.Writer, s string) error {
	_, err := w.Write([]byte(s))
	return err
}
