package core

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func PrintJSON(obj interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(obj)
}

func ReadFromStdin(obj interface{}) (error, bool) {
	dec := json.NewDecoder(os.Stdin)
	err := dec.Decode(obj)
	return err, errors.Is(err, io.EOF)
}
